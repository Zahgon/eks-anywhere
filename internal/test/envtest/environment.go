package envtest

import (
	"context"
	"testing"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	tinkv1alpha1 "github.com/tinkerbell/tink/api/v1alpha1"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	cloudstackv1 "sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta3"
	vspherev1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"
	addonsv1 "sigs.k8s.io/cluster-api/api/addons/v1beta2"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	clusterctlv1 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
	dockerv1beta2 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	tinkerbellv1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/capt/v1beta1"
	rufiov1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/rufio"
	snowv1 "github.com/aws/eks-anywhere/pkg/providers/snow/api/v1beta1"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	capiPackage         = "sigs.k8s.io/cluster-api"
	capdPackage         = "sigs.k8s.io/cluster-api/test"
	capvPackage         = "sigs.k8s.io/cluster-api-provider-vsphere"
	tinkerbellPackage   = "github.com/tinkerbell/tink"
	etcdProviderPackage = "github.com/aws/etcdadm-controller"
	capcPackage         = "sigs.k8s.io/cluster-api-provider-cloudstack"

	kubebuilderAssetsEnvVar = "KUBEBUILDER_ASSETS"
)

func init() {
	// Register CRDs in Scheme in init so fake clients benefit from it
	utilruntime.Must(corev1.AddToScheme(scheme.Scheme))
	utilruntime.Must(releasev1.AddToScheme(scheme.Scheme))
	utilruntime.Must(clusterv1beta2.AddToScheme(scheme.Scheme))
	utilruntime.Must(clusterctlv1.AddToScheme(scheme.Scheme))
	utilruntime.Must(controlplanev1beta2.AddToScheme(scheme.Scheme))
	utilruntime.Must(bootstrapv1beta2.AddToScheme(scheme.Scheme))
	utilruntime.Must(vspherev1.AddToScheme(scheme.Scheme))
	utilruntime.Must(dockerv1beta2.AddToScheme(scheme.Scheme))
	utilruntime.Must(cloudstackv1.AddToScheme(scheme.Scheme))
	utilruntime.Must(etcdv1.AddToScheme(scheme.Scheme))
	utilruntime.Must(admissionv1beta1.AddToScheme(scheme.Scheme))
	utilruntime.Must(anywherev1.AddToScheme(scheme.Scheme))
	utilruntime.Must(eksdv1alpha1.AddToScheme(scheme.Scheme))
	utilruntime.Must(snowv1.AddToScheme(scheme.Scheme))
	utilruntime.Must(addonsv1.AddToScheme(scheme.Scheme))
	utilruntime.Must(tinkerbellv1.AddToScheme(scheme.Scheme))
	utilruntime.Must(tinkv1alpha1.AddToScheme(scheme.Scheme))
	utilruntime.Must(rufiov1alpha1.AddToScheme(scheme.Scheme))
}

var packages = []moduleWithCRD{
	mustBuildModuleWithCRDs(capiPackage,
		withAdditionalCustomCRDPath("bootstrap/kubeadm/config/crd/bases"),
		withAdditionalCustomCRDPath("controlplane/kubeadm/config/crd/bases"),
	),
	mustBuildModuleWithCRDs(tinkerbellPackage),
	mustBuildModuleWithCRDs(capvPackage,
		withMainCustomCRDPath("config/default/crd/bases"),
	),
	mustBuildModuleWithCRDs(capdPackage,
		withMainCustomCRDPath("infrastructure/docker/config/crd/bases"),
	),
	mustBuildModuleWithCRDs(etcdProviderPackage),
}

type Environment struct {
	scheme  *runtime.Scheme
	client  client.Client
	env     *envtest.Environment
	manager manager.Manager
	// apiReader is a non cached client (only for reads), helpful when testing the actual state of objects
	apiReader client.Reader
	cancelF   context.CancelFunc
}

type EnvironmentOpt func(ctx context.Context, e *Environment)

func WithAssignment(envRef **Environment) EnvironmentOpt {
	_ = "STUB: not implemented"
	return *new(EnvironmentOpt)
}

// RunWithEnvironment runs a suite of tests with an envtest that is shared across all tests
// We use testing.M as the input to avoid having this called directly from a test
// This ensures the envtest setup is always run from a TestMain.
func RunWithEnvironment(m *testing.M, opts ...EnvironmentOpt) int {
	_ = "STUB: not implemented"
	return 0
}

func newEnvironment(ctx context.Context) (*Environment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// start webhook server using Manager

func ensureEnvtest(ctx context.Context, rootDir string) error {
	_ = "STUB: not implemented"
	// Only if the envtest config envvar is not set, try to setup assets
	return nil
}

// Last line of output is expected to be KUBEBUILDER_ASSETS=[path to envtest setup]

func (e *Environment) stop() error { _ = "STUB: not implemented"; return nil }

// Cancels context that will stop the manager

func (e *Environment) Client() client.Client {
	_ = "STUB: not implemented"

	// APIReader returns a non cached reader client.
	return *new(client.Client)
}

func (e *Environment) APIReader() client.Reader {
	_ = "STUB: not implemented"
	return *

	// Manager returns a Manager for the test environment.
	new(client.Reader)
}

func (e *Environment) Manager() manager.Manager {
	_ = "STUB: not implemented"
	return *new(manager.Manager)
}

func (e *Environment) CreateNamespaceForTest(ctx context.Context, t *testing.T) string {
	_ = "STUB: not implemented"
	return ""
}

func getRootPath() string { _ = "STUB: not implemented"; return "" }

func currentDir() string { _ = "STUB: not implemented"; return "" }
