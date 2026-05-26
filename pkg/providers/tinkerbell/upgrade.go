package tinkerbell

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	rufiov1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/rufio"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/rufiounreleased"
	"github.com/aws/eks-anywhere/pkg/types"
)

func (p *Provider) SetupAndValidateUpgradeCluster(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec, currentClusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// If we've been given a CSV with additional hardware for the cluster, validate it and
// write it to the catalogue so it can be used for further processing.

// Retrieve all unprovisioned hardware from the existing cluster and populate the catalogue so
// it can be considered for the upgrade.

// Retrieve all provisioned hardware from the existing cluster and populate diskExtractors's
// disksProvisionedHardware map for use during upgrade

// Remove all the provisioned hardware from the existing cluster if repeated from the hardware csv input.

// skip extra hardware validation for InPlace upgrades

// Update stack helm enviorment variable NO_PROXY value and append management cluster's Control plane Endpoint IP in case of workload cluster upgrade

// Apply hardware for workload clusters

// Check if the hardware in the catalogue have a BMCRef. Since we only allow either all hardware with bmc
// or no hardware with bmc, its sufficient to check the first hardware.

// Waiting to ensure all the new and exisiting baseboardmanagement connections are valid.

// SetupAndValidateUpgradeManagementComponents performs necessary setup for upgrade management components operation.
func (p *Provider) SetupAndValidateUpgradeManagementComponents(_ context.Context, _ *cluster.Spec) error {
	_ = "STUB: not implemented"

	// validateAvailableHardwareForUpgrade adds the necessary hardware assertions for an upgrade
	// these includes both rollingUpgrades across control-plane and worker nodes
	// and modular upgrades pertaining to only control-plane or worker components.
	return nil
}

func (p *Provider) validateAvailableHardwareForUpgrade(ctx context.Context, currentSpec, newClusterSpec *cluster.Spec) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Hardware selectors for controlPlane and worker nodes are mutually exclusive, so its safe to copy
// as no keys are going to be overwritten

// ScaleUpDown should not be supported in case of either rolling upgrade or eksa version upgrade.

// validateHardwareReqForControlPlaneRollOut checks if new ControlPlane nodes need to be rolled out
// and returns the HardwareRequirements for ControlPlane HardwareSelector
// as we open up more feature gates in our spec this function can be used to
// validate against the features that affect control plane only ex. API server extra args, etc.
func (p *Provider) validateHardwareReqForControlPlaneRollOut(currentSpec, newClusterSpec *cluster.Spec) (MinimumHardwareRequirements, error) {
	_ = "STUB: not implemented"
	return *new(MinimumHardwareRequirements), nil
}

// validateHardwareReqForWorkerNodeGroupsRollOut checks if new Worker nodes need to be rolled out for a given worker node group configuration
// and returns the HardwareRequirements for Worker node group HardwareSelector
// as we open up more feature gates in our spec this function can be used to
// validate against the features that affect worker groups only.
func (p *Provider) validateHardwareReqForWorkerNodeGroupsRollOut(currentSpec, newClusterSpec *cluster.Spec) (MinimumHardwareRequirements, error) {
	_ = "STUB: not implemented"
	return *new(MinimumHardwareRequirements), nil
}

// ApplyHardwareToCluster adds all the hardwares to the cluster.
func (p *Provider) applyHardwareUpgrade(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) PostMoveManagementToBootstrap(ctx context.Context, bootstrapCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// Check if the hardware in the catalogue have a BMCRef. Since we only allow either all hardware with bmc
	// or no hardware with bmc, its sufficient to check the first hardware.
	return nil
}

// Convert bmcTimeout duration to string format for WaitForRufioMachines

// Waiting to ensure all the new and exisiting baseboardmanagement connections are valid.

func (p *Provider) RunPostControlPlaneUpgrade(ctx context.Context, oldClusterSpec *cluster.Spec, clusterSpec *cluster.Spec, workloadCluster *types.Cluster, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// @TODO: do we need this for bare metal upgrade?
	return nil
}

// Use retrier so that cluster upgrade does not fail due to any intermittent failure while connecting to kube-api server

// This is unfortunate, but ClusterResourceSet's don't support any type of reapply of the resources they manage
// Even if we create a new ClusterResourceSet, if such resources already exist in the cluster, they won't be reapplied
// The long term solution is to add this capability to the cluster-api controller,
// with a new mode like "ReApplyOnChanges" or "ReApplyOnCreate" vs the current "ReApplyOnce"
/* err := p.retrier.Retry(
	func() error {
		return p.resourceSetManager.ForceUpdate(ctx, resourceSetName(clusterSpec), constants.EksaSystemNamespace, managementCluster, workloadCluster)
	},
)
if err != nil {
	return fmt.Errorf("failed updating the tinkerbell provider resource set post upgrade: %v", err)
} */

// ValidateNewSpec satisfies the Provider interface.
func (p *Provider) ValidateNewSpec(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// for any operation other than k8s version change, hookImageURL is immutable

func (p *Provider) validateMachineCfgsImmutability(ctx context.Context, clstr *types.Cluster, currentClstr *v1alpha1.Cluster, desiredClstrSpec *cluster.Spec, currentWNGs, desiredWNGs []v1alpha1.WorkerNodeGroupConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

// newWNGs contains the set of worker node group names specified in the desired spec that are new.

// If the machine config reference is for a new worker node group don't bother with
// immutability checks as we want users to be able to add worker node groups.

func validateRefsUnchanged(current, desired []v1alpha1.WorkerNodeGroupConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

// For every current worker node group that still exists in the desired config, ensure the
// machine config is still the same.

func (p *Provider) hardwareCSVIsProvided() bool { _ = "STUB: not implemented"; return false }

func (p *Provider) isScaleUpDown(oldCluster *v1alpha1.Cluster, newCluster *v1alpha1.Cluster) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Provider) isRollingUpgrade(currentSpec, newClusterSpec *cluster.Spec) bool {
	_ = "STUB: not implemented"
	return false
}

// WorkerNodeGroupWithK8sVersion maps each worker node group configurations in s to its K8s version.
func WorkerNodeGroupWithK8sVersion(spec *cluster.Spec) map[string]v1alpha1.KubernetesVersion {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) validateMachineCfg(ctx context.Context, cluster *types.Cluster, newConfig *v1alpha1.TinkerbellMachineConfig, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCoreComponentsUpgrade satisfies the Provider interface.
func (p *Provider) PreCoreComponentsUpgrade(
	ctx context.Context,
	cluster *types.Cluster,
	managementComponents *cluster.ManagementComponents,
	clusterSpec *cluster.Spec,
) error {
	_ = "STUB: not implemented"
	// When a workload cluster the cluster object could be nil. Noop if it is.
	return nil
}

// Tink stack and all the crds live on management cluster
// skip stack upgrade for workload clusters.

// For management clusters, upgrade the stack and apply hardware
// Attempt the upgrade. This should upgrade the stack in the management cluster by updating
// images, installing new CRDs and possibly removing old ones.

// Check if cluster has legacy chart installed

// Upgrade legacy chart to add resource policy keep to the CRDs

// load balancer is handled by kube-vip in control plane nodes
// configure load balancer based on datacenterConfig.Spec.SkipLoadBalancerDeployment

// Uninstall legacy chart

// annotate existing CRDs to point to new CRDs chart

// Check if cluster has stack chart installed (current production, pre-mono-repo)

// Uninstall stack chart before installing new mono-repo chart

// upgrade install CRDs chart

// upgrade install tink stack chart

// load balancer is handled by kube-vip in control plane nodes
// configure load balancer based on datacenterConfig.Spec.SkipLoadBalancerDeployment

// We introduced the Rufio dependency prior to its initial release. Between its introduction
// and its official release breaking changes occured to the CRDs. We're using the presence
// of the obsolete BaseboardManagement CRD to determine if there's an old Rufio installed.
// If there is, we need to convert all obsolete BaseboardManagement CRs to Machine CRs (the
// CRD that superseeds BaseboardManagement).

// Remove the unreleased Rufio CRDs from the cluster; this will also remove any residual
// resources.

func (p *Provider) handleRufioUnreleasedCRDs(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// Firstly, retrieve all BaseboardManagement CRs and convert them to Machine CRs.
	return nil
}

// Secondly, iterate over all Hardwarfe CRs and update the BMCRef to point to the new Machine
// CR.

func toRufioMachines(items []rufiounreleased.BaseboardManagement) []rufiov1.Machine {
	_ = "STUB: not implemented"
	return nil
}

// We need to populate type meta because we apply with kubectl (leakage).

func (p *Provider) annotateCRDs(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// machine

// task

// job

// hardware

// template

// workflow
