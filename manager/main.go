package main

import (
	"context"
	"flag"
	"os"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	"github.com/go-logr/logr"
	nutanixv1 "github.com/nutanix-cloud-native/cluster-api-provider-nutanix/api/v1beta1"
	"github.com/spf13/pflag"
	tinkv1alpha1 "github.com/tinkerbell/tink/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	logsv1 "k8s.io/component-base/logs/api/v1"
	_ "k8s.io/component-base/logs/json/register"
	"k8s.io/klog/v2"
	cloudstackv1 "sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta3"
	vspherev1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"
	addonsv1 "sigs.k8s.io/cluster-api/api/addons/v1beta2"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	clusterctlv1 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
	dockerv1beta2 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1beta2"
	capiflags "sigs.k8s.io/cluster-api/util/flags"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	tinkerbellv1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/capt/v1beta1"
	rufiov1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/rufio"
	"github.com/aws/eks-anywhere/pkg/controller/clientutil"
	"github.com/aws/eks-anywhere/pkg/features"
	snowv1 "github.com/aws/eks-anywhere/pkg/providers/snow/api/v1beta1"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

var scheme = runtime.NewScheme()

const WEBHOOK = "webhook"

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(anywherev1.AddToScheme(scheme))
	utilruntime.Must(releasev1.AddToScheme(scheme))
	utilruntime.Must(clusterv2.AddToScheme(scheme))
	utilruntime.Must(clusterctlv1.AddToScheme(scheme))
	utilruntime.Must(controlplanev1beta2.AddToScheme(scheme))
	utilruntime.Must(vspherev1.AddToScheme(scheme))
	utilruntime.Must(cloudstackv1.AddToScheme(scheme))
	utilruntime.Must(dockerv1beta2.AddToScheme(scheme))
	utilruntime.Must(etcdv1.AddToScheme(scheme))
	utilruntime.Must(bootstrapv1beta2.AddToScheme(scheme))
	utilruntime.Must(eksdv1alpha1.AddToScheme(scheme))
	utilruntime.Must(snowv1.AddToScheme(scheme))
	utilruntime.Must(addonsv1.AddToScheme(scheme))
	utilruntime.Must(tinkerbellv1.AddToScheme(scheme))
	utilruntime.Must(tinkv1alpha1.AddToScheme(scheme))
	utilruntime.Must(rufiov1alpha1.AddToScheme(scheme))
	utilruntime.Must(nutanixv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

type config struct {
	enableLeaderElection bool
	probeAddr            string
	gates                []string
	logging              *logsv1.LoggingConfiguration
	managerOptions       capiflags.ManagerOptions
}

func newConfig() *config { _ = "STUB: not implemented"; return nil }

func initFlags(fs *pflag.FlagSet, config *config) { _ = "STUB: not implemented"; return }

// +kubebuilder:rbac:groups=authentication.k8s.io,resources=tokenreviews,verbs=create
// +kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create
func main() {
	config := newConfig()
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	initFlags(pflag.CommandLine, config)
	pflag.Parse()

	// Temporary logger for initialization
	setupLog := ctrl.Log.WithName("setup")

	if err := logsv1.ValidateAndApply(config.logging, nil); err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// klog.Background will automatically use the right logger.
	ctrl.SetLogger(klog.Background())
	// Once controller-runtime logger has been setup correctly, retrieve again
	setupLog = ctrl.Log.WithName("setup")

	features.FeedGates(config.gates)

	_, metricsServerOpts, err := capiflags.GetManagerOptions(config.managerOptions)
	if err != nil {
		setupLog.Error(err, "Unable to start manager: invalid metrics server flags")
		os.Exit(1)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:  scheme,
		Metrics: *metricsServerOpts,
		WebhookServer: webhook.NewServer(webhook.Options{
			Port: 9443,
		}),
		HealthProbeBindAddress: config.probeAddr,
		LeaderElection:         config.enableLeaderElection,
		LeaderElectionID:       "f64ae69e.eks.amazonaws.com",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// Setup the context that's going to be used in controllers and for the manager.
	ctx := ctrl.SetupSignalHandler()

	closer := setupReconcilers(ctx, setupLog, mgr)
	defer func() {
		setupLog.Info("Closing reconciler dependencies")
		if err := closer.Close(ctx); err != nil {
			setupLog.Error(err, "Failed closing reconciler dependencies")
		}
	}()
	setupWebhooks(setupLog, mgr)
	setupChecks(setupLog, mgr)
	//+kubebuilder:scaffold:builder

	// Adding this indexer to allow for listing Cluster objects by name if they are in different namespaces.
	if err := mgr.GetFieldIndexer().
		IndexField(ctx, &anywherev1.Cluster{}, "metadata.name", clientutil.ClusterNameIndexer); err != nil {
		setupLog.Error(err, "unable to create index for Cluster name")
		os.Exit(1)
	}
	setupLog.Info("Starting manager")
	if err := mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

type closable interface {
	Close(ctx context.Context) error
}

func setupReconcilers(ctx context.Context, setupLog logr.Logger, mgr ctrl.Manager) closable {
	_ = "STUB: not implemented"
	return *new(closable)
}

func setupWebhooks(setupLog logr.Logger, mgr ctrl.Manager) { _ = "STUB: not implemented"; return }

func setupCoreWebhooks(setupLog logr.Logger, mgr ctrl.Manager) { _ = "STUB: not implemented"; return }

func setupVSphereWebhooks(setupLog logr.Logger, mgr ctrl.Manager) {
	_ = "STUB: not implemented"
	return
}

func setupCloudstackWebhooks(setupLog logr.Logger, mgr ctrl.Manager) {
	_ = "STUB: not implemented"
	return
}

func setupSnowWebhooks(setupLog logr.Logger, mgr ctrl.Manager) { _ = "STUB: not implemented"; return }

func setupTinkerbellWebhooks(setupLog logr.Logger, mgr ctrl.Manager) {
	_ = "STUB: not implemented"
	return
}

func setupNutanixWebhooks(setupLog logr.Logger, mgr ctrl.Manager) {
	_ = "STUB: not implemented"
	return
}

func setupChecks(setupLog logr.Logger, mgr ctrl.Manager) { _ = "STUB: not implemented"; return }
