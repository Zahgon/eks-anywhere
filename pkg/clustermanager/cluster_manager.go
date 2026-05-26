package clustermanager

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	tinkerbellv1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/capt/v1beta1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/diagnostics"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	maxRetries             = 30
	defaultBackOffPeriod   = 5 * time.Second
	machineBackoff         = 1 * time.Second
	defaultMachinesMinWait = 30 * time.Minute

	// DefaultMaxWaitPerMachine is the default max time the cluster manager will wait per a machine.
	DefaultMaxWaitPerMachine = 10 * time.Minute
	// DefaultClusterWait is the default max time the cluster manager will wait for the capi cluster to be in ready state.
	DefaultClusterWait = 60 * time.Minute
	// DefaultControlPlaneWait is the default time the cluster manager will wait for the control plane to be ready.
	DefaultControlPlaneWait = 60 * time.Minute
	// DefaultControlPlaneWaitAfterMove is the default max time the cluster manager will wait for the control plane to be in ready state after the capi move operation.
	DefaultControlPlaneWaitAfterMove = 15 * time.Minute
	// DefaultDeploymentWait is the default max time the cluster manager will wait for the deployment to be available.
	DefaultDeploymentWait = 30 * time.Minute

	// DefaultEtcdWait is the default time the cluster manager will wait for ectd to be ready.
	DefaultEtcdWait = 60 * time.Minute
	// DefaultUnhealthyMachineTimeout is the default timeout for an unhealthy machine health check.
	DefaultUnhealthyMachineTimeout = 5 * time.Minute
	// DefaultNodeStartupTimeout is the default timeout for a machine without a node to be considered to have failed machine health check.
	DefaultNodeStartupTimeout = 10 * time.Minute
	// DefaultClusterctlMoveTimeout is arbitrarily established.  Equal to kubectl wait default timeouts.
	DefaultClusterctlMoveTimeout = 30 * time.Minute
)

var (
	clusterctlNetworkErrorRegex              = regexp.MustCompile(`.*failed to connect to the management cluster:.*`)
	clusterctlMoveProvisionedInfraErrorRegex = regexp.MustCompile(`.*failed to check for provisioned infrastructure*`)
	kubectlResourceNotFoundRegex             = regexp.MustCompile(`.*the server doesn't have a resource type "(.*)".*`)
	eksaClusterResourceType                  = fmt.Sprintf("clusters.%s", v1alpha1.GroupVersion.Group)
)

type ClusterManager struct {
	eksaComponents     EKSAComponents
	ClientFactory      ClientFactory
	clusterClient      ClusterClient
	retrier            *retrier.Retrier
	writer             filewriter.FileWriter
	diagnosticsFactory diagnostics.DiagnosticBundleFactory

	machineMaxWait                   time.Duration
	machineBackoff                   time.Duration
	machinesMinWait                  time.Duration
	controlPlaneWaitTimeout          time.Duration
	controlPlaneWaitAfterMoveTimeout time.Duration
	externalEtcdWaitTimeout          time.Duration
	unhealthyMachineTimeout          time.Duration
	nodeStartupTimeout               time.Duration
	clusterWaitTimeout               time.Duration
	deploymentWaitTimeout            time.Duration
	clusterctlMoveTimeout            time.Duration
}

// ClientFactory builds Kubernetes clients.
type ClientFactory interface {
	// BuildClientFromKubeconfig builds a Kubernetes client from a kubeconfig file.
	BuildClientFromKubeconfig(kubeconfigPath string) (kubernetes.Client, error)
}

// CAPIClient performs operations on a cluster-api management cluster.
type CAPIClient interface {
	BackupManagement(ctx context.Context, cluster *types.Cluster, managementStatePath, clusterName string) error
	MoveManagement(ctx context.Context, from, target *types.Cluster, clusterName string) error
	InitInfrastructure(ctx context.Context, managementComponents *cluster.ManagementComponents, clusterSpec *cluster.Spec, cluster *types.Cluster, provider providers.Provider) error
	GetWorkloadKubeconfig(ctx context.Context, clusterName string, cluster *types.Cluster) ([]byte, error)
}

// EKSAComponents allows to manage the eks-a components installation in a cluster.
type EKSAComponents interface {
	Install(ctx context.Context, log logr.Logger, cluster *types.Cluster, managementComponents *cluster.ManagementComponents, spec *cluster.Spec) error
	Upgrade(ctx context.Context, log logr.Logger, cluster *types.Cluster, currentManagementComponents, newManagementComponents *cluster.ManagementComponents, newSpec *cluster.Spec) (*types.ChangeDiff, error)
}

type ClusterManagerOpt func(*ClusterManager)

// DefaultRetrier builds a retrier with the default configuration.
func DefaultRetrier() *retrier.Retrier { _ = "STUB: not implemented"; return nil }

// New constructs a new ClusterManager.
func New(client ClientFactory, clusterClient ClusterClient, writer filewriter.FileWriter, diagnosticBundleFactory diagnostics.DiagnosticBundleFactory, eksaComponents EKSAComponents, opts ...ClusterManagerOpt) *ClusterManager {
	_ = "STUB: not implemented"
	return nil
}

func WithControlPlaneWaitTimeout(timeout time.Duration) ClusterManagerOpt {
	_ = "STUB: not implemented"
	return *new(ClusterManagerOpt)
}

func WithExternalEtcdWaitTimeout(timeout time.Duration) ClusterManagerOpt {
	_ = "STUB: not implemented"
	return *new(ClusterManagerOpt)
}

func WithMachineBackoff(machineBackoff time.Duration) ClusterManagerOpt {
	_ = "STUB: not implemented"
	return *new(ClusterManagerOpt)
}

func WithMachineMaxWait(machineMaxWait time.Duration) ClusterManagerOpt {
	_ = "STUB: not implemented"
	return *new(ClusterManagerOpt)
}

func WithMachineMinWait(machineMinWait time.Duration) ClusterManagerOpt {
	_ = "STUB: not implemented"
	return *new(ClusterManagerOpt)
}

// WithUnhealthyMachineTimeout sets the timeout of an unhealthy machine health check.
func WithUnhealthyMachineTimeout(timeout time.Duration) ClusterManagerOpt {
	_ = "STUB: not implemented"
	return *new(ClusterManagerOpt)
}

// WithNodeStartupTimeout sets the timeout of a machine without a node to be considered to have failed machine health check.
func WithNodeStartupTimeout(timeout time.Duration) ClusterManagerOpt {
	_ = "STUB: not implemented"
	return *new(ClusterManagerOpt)
}

func WithRetrier(retrier *retrier.Retrier) ClusterManagerOpt {
	_ = "STUB: not implemented"
	return *new(ClusterManagerOpt)
}

// WithNoTimeouts disables the timeout for all the waits and retries in cluster manager.
func WithNoTimeouts() ClusterManagerOpt { _ = "STUB: not implemented"; return *new(ClusterManagerOpt) }

func clusterctlMoveWaitForInfrastructureRetryPolicy(totalRetries int, err error) (retry bool, wait time.Duration) {
	_ = "STUB: not implemented"
	// Retry both network and cluster move errors.
	return false, *new(time.Duration)
}

func clusterctlMoveRetryPolicy(totalRetries int, err error) (retry bool, wait time.Duration) {
	_ = "STUB: not implemented"
	// Retry only network errors.
	return false, *new(time.Duration)
}

func kubectlWaitRetryPolicy(totalRetries int, err error) (retry bool, wait time.Duration) {
	_ = "STUB: not implemented"
	// Sometimes it is possible that the clusterctl move is successful,
	// but the clusters.cluster.x-k8s.io resource is not available on the cluster yet.
	//
	// Retry on transient 'server doesn't have a resource type' errors.
	// Use existing exponential backoff implementation for retry on these errors.
	return false, *new(time.Duration)
}

func exponentialRetryWaitTime(totalRetries int) time.Duration {
	_ = "STUB: not implemented"
	// Exponential backoff on errors.  Retrier built-in backoff is linear, so implementing here.
	return *new(time.Duration)
}

// Retrier first calls the policy before retry #1.  We want it zero-based for exponentiation.

// BackupCAPI takes backup of management cluster's resources during the upgrade process.
func (c *ClusterManager) BackupCAPI(ctx context.Context, cluster *types.Cluster, managementStatePath, clusterName string) error {
	_ = "STUB: not implemented"
	// Network errors, most commonly connection refused or timeout, can occur if either source
	// cluster becomes inaccessible during the move operation.  If this occurs without retries, clusterctl
	// abandons the move operation, and fails cluster upgrade.
	// Retrying once connectivity is re-established completes the partial move.
	// Here we use a retrier, with the above defined clusterctlMoveRetryPolicy policy, to attempt to
	// wait out the network disruption and complete the move.
	// Keeping clusterctlMoveTimeout to the same as MoveManagement since both uses the same command with the differrent params.
	return nil
}

// BackupCAPIWaitForInfrastructure takes backup of bootstrap cluster's resources during the upgrade process
// like BackupCAPI but with a retry policy to wait for infrastructure provisioning in addition to network errors.
func (c *ClusterManager) BackupCAPIWaitForInfrastructure(ctx context.Context, cluster *types.Cluster, managementStatePath, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) backupCAPI(ctx context.Context, cluster *types.Cluster, managementStatePath, clusterName string, retrier *retrier.Retrier) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) MoveCAPI(ctx context.Context, from, to *types.Cluster, clusterName string, clusterSpec *cluster.Spec, checkers ...types.NodeReadyChecker) error {
	_ = "STUB: not implemented"
	return nil
}

// Network errors, most commonly connection refused or timeout, can occur if either source or target
// cluster becomes inaccessible during the move operation.  If this occurs without retries, clusterctl
// abandons the move operation, leaving an unpredictable subset of the CAPI components copied to target
// or deleted from source.  Retrying once connectivity is re-established completes the partial move.
// Here we use a retrier, with the above defined clusterctlMoveRetryPolicy policy, to attempt to
// wait out the network disruption and complete the move.

// CreateRegistryCredSecret creates the registry-credentials secret on a managment cluster.
func (c *ClusterManager) CreateRegistryCredSecret(ctx context.Context, mgmt *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// InstallCAPI installs the cluster-api components in a cluster.
func (c *ClusterManager) InstallCAPI(ctx context.Context, managementComponents *cluster.ManagementComponents, clusterSpec *cluster.Spec, cluster *types.Cluster, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) waitForCAPI(ctx context.Context, cluster *types.Cluster, provider providers.Provider, externalEtcdTopology bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) waitForDeployments(ctx context.Context, deploymentsByNamespace map[string][]string, cluster *types.Cluster, timeout string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) SaveLogsManagementCluster(ctx context.Context, spec *cluster.Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) SaveLogsWorkloadCluster(ctx context.Context, provider providers.Provider, spec *cluster.Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func collectDiagnosticBundle(ctx context.Context, bundle diagnostics.DiagnosticBundle) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) waitForControlPlaneReplicasReady(ctx context.Context, managementCluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) waitForMachineDeploymentReplicasReady(ctx context.Context, managementCluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// totalTimeoutForMachinesReadyWait calculates the total timeout when waiting for machines to be ready.
// The timeout increases linearly with the number of machines but can never be less than the configured
// minimun.
func (c *ClusterManager) totalTimeoutForMachinesReadyWait(replicaCount int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *ClusterManager) waitForNodesReady(ctx context.Context, managementCluster *types.Cluster, clusterName string, labels []string, checkers ...types.NodeReadyChecker) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) getNodesCount(ctx context.Context, managementCluster *types.Cluster, clusterName string, labels []string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *ClusterManager) countNodesReady(ctx context.Context, managementCluster *types.Cluster, clusterName string, labels []string, checkers ...types.NodeReadyChecker) (ready int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Extracted from cluster-api: NodeRef is considered a better signal than InfrastructureReady,
// because it ensures the node in the workload cluster is up and running.

// Upgrade updates the eksa components in a cluster according to a Spec.
func (c *ClusterManager) Upgrade(ctx context.Context, cluster *types.Cluster, currentManagementComponents, newManagementComponents *cluster.ManagementComponents, newSpec *cluster.Spec) (*types.ChangeDiff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterManager) CreateEKSANamespace(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) ApplyBundles(ctx context.Context, clusterSpec *cluster.Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// We need to update this config map with the new upgrader images whenever we
// apply a new Bundles object to the cluster in order to support in-place upgrades.

// ApplyReleases applies the EKSARelease manifest.
func (c *ClusterManager) ApplyReleases(ctx context.Context, clusterSpec *cluster.Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// PauseCAPIWorkloadClusters pauses all workload CAPI clusters except the management cluster.
func (c *ClusterManager) PauseCAPIWorkloadClusters(ctx context.Context, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// skip pausing management cluster

func (c *ClusterManager) resumeEksaReconcileForManagementAndWorkloadClusters(ctx context.Context, managementCluster *types.Cluster, clusterSpec *cluster.Spec, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

// ResumeEKSAControllerReconcile resumes a paused EKS-Anywhere cluster.
func (c *ClusterManager) ResumeEKSAControllerReconcile(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec, provider providers.Provider) error {
	_ = "STUB: not implemented"
	// clear pause annotation
	return nil
}

func (c *ClusterManager) resumeReconcileForCluster(ctx context.Context, clusterCreds *types.Cluster, cluster *v1alpha1.Cluster, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

// ResumeCAPIWorkloadClusters resumes all workload CAPI clusters except the management cluster.
func (c *ClusterManager) ResumeCAPIWorkloadClusters(ctx context.Context, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// skip resuming management cluster

// AllowDeleteWhilePaused allows the deletion of paused clusters.
func (c *ClusterManager) AllowDeleteWhilePaused(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) allowDeleteWhilePaused(ctx context.Context, clusterCreds *types.Cluster, cluster *v1alpha1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) PauseEKSAControllerReconcile(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) pauseEksaReconcileForManagementAndWorkloadClusters(ctx context.Context, managementCluster *types.Cluster, clusterSpec *cluster.Spec, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) pauseReconcileForCluster(ctx context.Context, clusterCreds *types.Cluster, cluster *v1alpha1.Cluster, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterManager) GetCurrentClusterSpec(ctx context.Context, clus *types.Cluster, clusterName string) (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterManager) buildSpecForCluster(ctx context.Context, clus *types.Cluster, eksaCluster *v1alpha1.Cluster) (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterManager) getUpgraderImagesFromBundle(ctx context.Context, cluster *types.Cluster, cl *cluster.Spec) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newUpgraderConfigMap(m map[string]string) *corev1.ConfigMap {
	_ = "STUB: not implemented"
	return nil
}

// As the Tink stack gets redeployed in the management cluster tinkerbell IP changes from bootstrap IP
// to the actual Tinkerbell IP specified in datacenter spec. We will need to update this IP in the
// TinkerbellMachineTemplate as the previous bootStrap IP is no longer serving the Tink stack.
// Also there is a new rollout once the eks-a controller comes up on the management cluster as it sees
// the IP change in the template as a diff in spec. To prevent this from happening update the objects
// in-place before the move. Since TinkerbellMachineTemplate is immutable we get the object, update
// the IP, delete and recreate the object.
// For long term, we want to revisit how we handle the bootstrap vs management cluster case in eks-a
// controller specific to baremetal provider as the source of truth gets changed due to the nature of
// tink stack being moved.
// nolint:gocyclo
func updateTinkerbellIPInBootstrapTinkerbellMachineTemplate(ctx context.Context, spec *cluster.Spec, client kubernetes.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// isoURL path is only served in the top level /iso path.

// When an templateOverride is specified in the spec, we do not want to modify it.

// When an templateOverride is specified in the spec, we do not want to modify it.
// We update the tinkebelltemplate config only for the corresponding worker node group.

func updateTemplateOverride(clusterSpec *v1alpha1.Cluster, template tinkerbellv1.TinkerbellMachineTemplate, osImageOverride, tinkIP string, osFamily v1alpha1.OSFamily) (tinkerbellv1.TinkerbellMachineTemplate, error) {
	_ = "STUB: not implemented"
	return *new(tinkerbellv1.TinkerbellMachineTemplate), nil
}
