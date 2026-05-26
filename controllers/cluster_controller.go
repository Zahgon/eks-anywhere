package controllers

import (
	"context"
	"time"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	"github.com/go-logr/logr"
	v1beta1patch "sigs.k8s.io/cluster-api/util/deprecated/v1beta1/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	c "github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
	"github.com/aws/eks-anywhere/pkg/controller/clientutil"
	"github.com/aws/eks-anywhere/pkg/controller/clusters"
	"github.com/aws/eks-anywhere/pkg/curatedpackages"
	"github.com/aws/eks-anywhere/pkg/providers/vsphere"
	"github.com/aws/eks-anywhere/pkg/registrymirror"
	"github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	defaultRequeueTime = time.Minute
	// ClusterFinalizerName is the finalizer added to clusters to handle deletion.
	ClusterFinalizerName = "clusters.anywhere.eks.amazonaws.com/finalizer"
	releaseV022          = "v0.22.0"
)

// ClusterReconciler reconciles a Cluster object.
type ClusterReconciler struct {
	client                     client.Client
	providerReconcilerRegistry ProviderClusterReconcilerRegistry
	awsIamAuth                 AWSIamConfigReconciler
	clusterValidator           ClusterValidator
	packagesClient             PackagesClient
	machineHealthCheck         MachineHealthCheckReconciler
	vSpherefailureDomainMover  FailureDomainApplier
}

// PackagesClient handles curated packages operations from within the cluster
// controller.
type PackagesClient interface {
	EnableFullLifecycle(ctx context.Context, log logr.Logger, clusterName, kubeConfig string, chart *v1alpha1.Image, registry *registrymirror.RegistryMirror, options ...curatedpackages.PackageControllerClientOpt) error
	ReconcileDelete(context.Context, logr.Logger, curatedpackages.KubeDeleter, *anywherev1.Cluster) error
	Reconcile(context.Context, logr.Logger, client.Client, *anywherev1.Cluster) error
}

type ProviderClusterReconcilerRegistry interface {
	Get(datacenterKind string) clusters.ProviderClusterReconciler
}

// AWSIamConfigReconciler manages aws-iam-authenticator installation and configuration for an eks-a cluster.
type AWSIamConfigReconciler interface {
	EnsureCASecret(ctx context.Context, logger logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error)
	Reconcile(ctx context.Context, logger logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error)
	ReconcileDelete(ctx context.Context, logger logr.Logger, cluster *anywherev1.Cluster) error
	ReconcileWorkloadClusterDelete(ctx context.Context, logger logr.Logger, cluster *anywherev1.Cluster, awsIamConfig *anywherev1.AWSIamConfig) error
}

// MachineHealthCheckReconciler manages machine health checks for an eks-a cluster.
type MachineHealthCheckReconciler interface {
	Reconcile(ctx context.Context, logger logr.Logger, cluster *anywherev1.Cluster) error
}

// ClusterValidator runs cluster level preflight validations before it goes to provider reconciler.
type ClusterValidator interface {
	ValidateManagementClusterName(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) error
}

// ClusterReconcilerOption allows to configure the ClusterReconciler.
type ClusterReconcilerOption func(*ClusterReconciler)

// SpecBuilder builds a cluster specification from an EKS Anywhere Cluster object.
type SpecBuilder interface {
	BuildSpec(ctx context.Context, eksaCluster *anywherev1.Cluster) (*c.Spec, error)
}

// FailureDomainSpecBuilder transforms a cluster specification into VSphere-specific failure domains.
type FailureDomainSpecBuilder interface {
	BuildFailureDomainSpec(log logr.Logger, clusterSpec *c.Spec) (*vsphere.FailureDomains, error)
}

// ObjectReconciler applies failure domain objects to a Kubernetes cluster.
type ObjectReconciler interface {
	ReconcileObjects(ctx context.Context, fd *vsphere.FailureDomains) error
}

// FailureDomainApplier orchestrates the end-to-end process of applying failure domains to a cluster.
type FailureDomainApplier interface {
	ApplyFailureDomains(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) error
}

// DefaultSpecBuilder is the standard implementation of SpecBuilder that uses a Kubernetes client.
type DefaultSpecBuilder struct {
	client client.Client
}

// BuildSpec is a wrapper method for building and obtaining all neccessary objects from a cluster.
func (b *DefaultSpecBuilder) BuildSpec(ctx context.Context, cluster *anywherev1.Cluster) (*c.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DefaultFailureDomainSpecBuilder is the standard implementation of FailureDomainSpecBuilder.
type DefaultFailureDomainSpecBuilder struct{}

// BuildFailureDomainSpec wrapper to the vsphere package's FailureDomainsSpec function.
func (b *DefaultFailureDomainSpecBuilder) BuildFailureDomainSpec(log logr.Logger, clusterSpec *c.Spec) (*vsphere.FailureDomains, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DefaultObjectReconciler is the standard implementation of ObjectReconciler that applies objects using a client.
type DefaultObjectReconciler struct {
	client client.Client
}

// ReconcileObjects applies failure domain objects to the cluster using serverside reconciliation.
func (r *DefaultObjectReconciler) ReconcileObjects(ctx context.Context, fd *vsphere.FailureDomains) error {
	_ = "STUB: not implemented"
	return nil
}

// FailureDomainMover defines config for applying failure domain objects.
type FailureDomainMover struct {
	specBuilder      SpecBuilder
	fdSpecBuilder    FailureDomainSpecBuilder
	objectReconciler ObjectReconciler
}

// NewFailureDomainMover builds FailureDomainMover with default dependencies.
func NewFailureDomainMover(client client.Client) *FailureDomainMover {
	_ = "STUB: not implemented"
	return nil
}

// NewFailureDomainMoverWithDependencies builds FailureDomainMover with specified dependencies.
func NewFailureDomainMoverWithDependencies(
	specBuilder SpecBuilder,
	fdSpecBuilder FailureDomainSpecBuilder,
	objectReconciler ObjectReconciler,
) *FailureDomainMover {
	_ = "STUB: not implemented"
	return nil
}

// NewClusterReconciler constructs a new ClusterReconciler.
func NewClusterReconciler(client client.Client, registry ProviderClusterReconcilerRegistry, awsIamAuth AWSIamConfigReconciler, clusterValidator ClusterValidator, pkgs PackagesClient, machineHealthCheck MachineHealthCheckReconciler, failuredomainmover FailureDomainApplier, opts ...ClusterReconcilerOption) *ClusterReconciler {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterReconciler) SetupWithManager(mgr ctrl.Manager, log logr.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch;update
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;create;delete;update;patch
// +kubebuilder:rbac:groups="",namespace=eksa-system,resources=secrets,verbs=patch;update
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;create;delete;patch;update
// +kubebuilder:rbac:groups="",resources=nodes,verbs=list
// +kubebuilder:rbac:groups=addons.cluster.x-k8s.io,resources=clusterresourcesets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=clusters;gitopsconfigs;snowmachineconfigs;snowdatacenterconfigs;snowippools;vspheredatacenterconfigs;vspheremachineconfigs;dockerdatacenterconfigs;tinkerbellmachineconfigs;tinkerbelltemplateconfigs;tinkerbelldatacenterconfigs;cloudstackdatacenterconfigs;cloudstackmachineconfigs;nutanixdatacenterconfigs;nutanixmachineconfigs;oidcconfigs;fluxconfigs,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=awsiamconfigs,verbs=get;list;watch;update;patch;delete
// +kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=clusters/status;snowmachineconfigs/status;snowippools/status;vspheredatacenterconfigs/status;vspheremachineconfigs/status;dockerdatacenterconfigs/status;tinkerbelldatacenterconfigs/status;tinkerbellmachineconfigs/status;tinkerbelltemplateconfigs/status;cloudstackdatacenterconfigs/status;cloudstackmachineconfigs/status;awsiamconfigs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=bundles,verbs=get;list;watch
// +kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=clusters/finalizers;snowmachineconfigs/finalizers;snowippools/finalizers;vspheredatacenterconfigs/finalizers;vspheremachineconfigs/finalizers;cloudstackdatacenterconfigs/finalizers;cloudstackmachineconfigs/finalizers;dockerdatacenterconfigs/finalizers;bundles/finalizers;awsiamconfigs/finalizers;tinkerbelldatacenterconfigs/finalizers;tinkerbellmachineconfigs/finalizers;tinkerbelltemplateconfigs/finalizers,verbs=update
// +kubebuilder:rbac:groups=bootstrap.cluster.x-k8s.io,resources=kubeadmconfigtemplates,verbs=create;get;list;patch;update;watch
// +kubebuilder:rbac:groups="cluster.x-k8s.io",resources=machinedeployments,verbs=list;watch;get;patch;update;create;delete
// +kubebuilder:rbac:groups="cluster.x-k8s.io",resources=clusters,verbs=list;watch;get;patch;update;create;delete
// +kubebuilder:rbac:groups="cluster.x-k8s.io",resources=machinehealthchecks,verbs=list;watch;get;patch;create
// +kubebuilder:rbac:groups=clusterctl.cluster.x-k8s.io,resources=providers,verbs=get;list;watch
// +kubebuilder:rbac:groups=controlplane.cluster.x-k8s.io,resources=kubeadmcontrolplanes,verbs=list;get;watch;patch;update;create;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;get;list;update;watch;delete
// +kubebuilder:rbac:groups=distro.eks.amazonaws.com,resources=releases,verbs=get;list;watch
// +kubebuilder:rbac:groups=etcdcluster.cluster.x-k8s.io,resources=*,verbs=create;get;list;patch;update;watch
// +kubebuilder:rbac:groups=tinkerbell.org,resources=hardware,verbs=list;watch
// +kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=machines,verbs=list;watch
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=awssnowclusters;awssnowmachinetemplates;awssnowippools;vsphereclusters;vspheremachinetemplates;dockerclusters;dockermachinetemplates;tinkerbellclusters;tinkerbellmachinetemplates;cloudstackclusters;cloudstackmachinetemplates;nutanixclusters;nutanixmachinetemplates;vspherefailuredomains;vspheredeploymentzones,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=packages.eks.amazonaws.com,resources=packages,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=packages.eks.amazonaws.com,namespace=eksa-system,resources=packagebundlecontrollers,verbs=delete
// +kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=eksareleases,verbs=get;list;watch
// The eksareleases permissions are being moved to the ClusterRole due to client trying to list this resource from cache.
// When trying to list resources not already in cache, it starts an informer for that type using the scope of the cache.
// So if the manager is cluster-scoped, the new informers created by the cache will be cluster-scoped

// Reconcile reconciles a cluster object.
// nolint:gocyclo
// TODO: Reduce high cycomatic complexity. https://github.com/aws/eks-anywhere-internal/issues/1449
func (r *ClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Fetch the Cluster objects

// Initialize the patch helper

// Always attempt to patch the object and status after each reconciliation.

// We want the observedGeneration to indicate, that the status shown is up-to-date given the desired spec of the same generation.
// However, if there is an error while updating the status, we may get a partial status update, In this case,
// a partially updated status is not considered up to date, so we should not update the observedGeneration

// Patch ObservedGeneration only if the reconciliation completed without error

// Only requeue if we are not already re-queueing and the Cluster ready condition is false.
// We do this to be able to update the status continuously until the cluster becomes ready,
// since there might be changes in state of the world that don't trigger reconciliation requests

// If the cluster is paused, return without any further processing.

// AddFinalizer	is idempotent

// If there is no difference between the aggregated generation and childrenReconciledGeneration,
// and there is no difference in the reconciled generation and .metadata.generation of the cluster,
// then return without any further processing.

// Failure messages are cleared in the reconciler loop after running validations. But sometimes,
// it seems that Cluster failure messages on the status are is not cleared for some reason
//  after successfully passing the validation. The theory is that if the inital patch operation
// is not successful, and the reconciliation is skipped going forward, it may never be cleared.
//
// When the controller reaches here, it denotes a completed reconcile. So, we can safely
// clear any failure messages or reasons that may be left over as there is no further processing
// for the controller to do.

func (r *ClusterReconciler) reconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster, aggregatedGeneration int64) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// At the end of the reconciliation, if there have been no requeues or errors, we update the cluster's status.
// NOTE: This update must be the last step in the reconciliation process to denote the complete reconciliation.
// No other mutating changes or reconciliations must happen in this loop after this step, so all such changes must
// be placed above this line.

// TODO(eksa-controller-SME): properly handle packages reconcile error and not triggering machine upgrade when
// packages reconcile is still in progress.
// Moving the packages reconcile after the above two generation fields are set, so that packages reconcile error
// does not cause side effect of rolling out of workload cluster machines during management cluster upgrade.

func (r *ClusterReconciler) preClusterProviderReconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	// Run some preflight validations that can't be checked in webhook
	return *new(controller.Result), nil
}

func (r *ClusterReconciler) postClusterProviderReconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// Check for orphaned AWS IAM config objects and clean up if needed

func (r *ClusterReconciler) packagesReconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	// Self-managed clusters can support curated packages, but that support
	// comes from the CLI at this time.
	return *new(controller.Result), nil
}

func (r *ClusterReconciler) updateStatus(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	// When EKS-A cluster is fully deleted, we do not need to update the status. Without this check
	// the subsequent patch operations would fail if the status is updated after it is fully deleted.
	return nil
}

// Always update the readyCondition by summarizing the state of other conditions.

// cleanupOrphanedAWSIamConfig checks for orphaned AWSIamConfig objects and cleans up resources.
func (r *ClusterReconciler) cleanupOrphanedAWSIamConfig(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	// Look for orphaned AWSIamConfig objects owned by this cluster
	return nil
}

// Clean up workload cluster resources

// Clean up management cluster resources (CA secrets, kubeconfig secrets)

// Delete the orphaned AWSIamConfig object itself

// isOwnedByCluster checks if an object is owned by the given cluster.
func (r *ClusterReconciler) isOwnedByCluster(obj client.Object, cluster *anywherev1.Cluster) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *ClusterReconciler) reconcileDelete(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Creates vspheredeploymentzone and vspherefailuredomain CR on bootstrap cluster prior to delete
// These CRs are not migrated over during pivot and must be present to cleanly delete vspheremachines
// This solution isn't ideal, but would require redesign

// TODO delete GitOps,Datacenter and MachineConfig objects

func (r *ClusterReconciler) buildClusterConfig(ctx context.Context, clus *anywherev1.Cluster) (*c.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ClusterReconciler) applyFailureDomains(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyFailureDomains orchestrates the end-to-end process of applying failure domain objects to a cluster.
func (m *FailureDomainMover) ApplyFailureDomains(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if VSphereDatacenter has no failure domains

func (r *ClusterReconciler) ensureClusterOwnerReferences(ctx context.Context, clus *anywherev1.Cluster, config *c.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// obj already had the owner reference

func patchCluster(ctx context.Context, patchHelper *v1beta1patch.Helper, cluster *anywherev1.Cluster, patchOpts ...v1beta1patch.Option) error {
	_ = "STUB: not implemented"
	// Patch the object, ignoring conflicts on the conditions owned by this controller.
	return nil
}

// Add each condition her that the controller should ignored conflicts for.

// Always attempt to patch the object and status after each reconciliation.

// aggregatedGeneration computes the combined generation of the resources linked
// by the cluster by summing up the .metadata.generation value for all the child
// objects of this cluster.
func aggregatedGeneration(config *c.Config) int64 { _ = "STUB: not implemented"; return 0 }

func getManagementCluster(ctx context.Context, clus *anywherev1.Cluster, client client.Client) (*anywherev1.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ClusterReconciler) setDefaultBundlesRefOrEksaVersion(ctx context.Context, clus *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEksaRelease(ctx context.Context, client client.Client, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateExtendedKubernetesVersionSupport(ctx context.Context, client client.Client, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip the signature validation for those versions prior to 'v0.22.0'

// Get the release manifest using Kubernetes client

// getReleaseManifestFromCluster retrieves the EKS Distro release manifest using the Kubernetes client.
// This is used in the controller context where we always use the Kubernetes client to fetch the manifest.
func getReleaseManifestFromCluster(ctx context.Context, clusterSpec anywherev1.Cluster, bundle *v1alpha1.Bundles, k *clientutil.KubeClient) (*eksdv1alpha1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
