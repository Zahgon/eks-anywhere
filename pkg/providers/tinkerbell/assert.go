package tinkerbell

import (
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	tinkerbellv1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/capt/v1beta1"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	"github.com/aws/eks-anywhere/pkg/networkutils"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
)

// TODO(chrisdoherty) Add worker node group assertions

// AssertMachineConfigsValid iterates over all machine configs in calling validateMachineConfig.
func AssertMachineConfigsValid(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

// AssertDatacenterConfigValid asserts the DatacenterConfig in spec is valid.
func AssertDatacenterConfigValid(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

// AssertMachineConfigNamespaceMatchesDatacenterConfig ensures all machine configuration instances
// are configured with the same namespace as the provider specific data center configuration
// namespace.
func AssertMachineConfigNamespaceMatchesDatacenterConfig(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// AssertControlPlaneMachineRefExists ensures the control plane machine ref is referencing a
// known machine config.
func AssertControlPlaneMachineRefExists(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// AssertEtcdMachineRefExists ensures that, if the etcd configuration is specified, it references
// a known machine config.
func AssertEtcdMachineRefExists(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	// Unstacked etcd is optional.
	return nil
}

// AssertWorkerNodeGroupMachineRefsExists ensures all worker node group machine refs are
// referencing a known machine config.
func AssertWorkerNodeGroupMachineRefsExists(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// AssertK8SVersionNot120 ensures Kubernetes version is not set to v1.20.
func AssertK8SVersionNot120(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

func AssertOsFamilyValid(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

// AssertUpgradeRolloutStrategyValid ensures that the upgrade rollout strategy is valid for both CP and worker node configurations.
func AssertUpgradeRolloutStrategyValid(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// AssertAutoScalerDisabledForInPlace ensures that the autoscaler configuration is not enabled when upgrade rollout strategy is InPlace.
func AssertAutoScalerDisabledForInPlace(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// AssertOSImageURL ensures that the OSImageURL value is either set at the datacenter config level or set for each machine config and not at both levels.
func AssertOSImageURL(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

// AssertISOURL ensures that the ISOURL value set is in the expected file format that the smee deployment expects.
func AssertISOURL(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

// AssertcontrolPlaneIPNotInUse ensures the endpoint host for the control plane isn't in use.
// The check may be unreliable due to its implementation.
func NewIPNotInUseAssertion(client networkutils.NetClient) ClusterSpecAssertion {
	_ = "STUB: not implemented"
	return *new(ClusterSpecAssertion)
}

// AssertTinkerbellIPNotInUse ensures tinkerbell ip isn't in use.
func AssertTinkerbellIPNotInUse(client networkutils.NetClient) ClusterSpecAssertion {
	_ = "STUB: not implemented"
	return *new(ClusterSpecAssertion)
}

// AssertTinkerbellIPAndControlPlaneIPNotSame ensures tinkerbell ip and controlplane ip are not the same.
func AssertTinkerbellIPAndControlPlaneIPNotSame(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// AssertHookRetrievableWithoutProxy ensures the executing machine can retrieve Hook
// from the host URL without a proxy configured. It does not guarantee the target node
// will be able to download Hook.
func AssertHookRetrievableWithoutProxy(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// return an error if hookImagesURLPath field is not specified for during Proxy configuration.

// verify hookImagesURLPath is accessible locally too

// AssertPortsNotInUse ensures that ports 80, 42113, and 7172 are available.
func AssertPortsNotInUse(client networkutils.NetClient) ClusterSpecAssertion {
	_ = "STUB: not implemented"
	return *new(ClusterSpecAssertion)
}

// HardwareSatisfiesOnlyOneSelectorAssertion ensures hardware in catalogue only satisfies 1
// of the MachineConfig's HardwareSelector's from the spec.
func HardwareSatisfiesOnlyOneSelectorAssertion(catalogue *hardware.Catalogue) ClusterSpecAssertion {
	_ = "STUB: not implemented"
	return *new(ClusterSpecAssertion)
}

// selectorsFromClusterSpec extracts all selectors specified on MachineConfig's from spec.
// When HardwareAffinity is used, it extracts matchLabels from Required terms.
func selectorsFromClusterSpec(spec *ClusterSpec) (selectorSet, error) {
	_ = "STUB: not implemented"
	return *new(selectorSet), nil
}

// addSelectorsFromMachineConfig extracts selectors from a machine config.
// If HardwareAffinity is set, it extracts matchLabels from Required terms.
// Otherwise, it uses the HardwareSelector.
func addSelectorsFromMachineConfig(config *v1alpha1.TinkerbellMachineConfig, selectors *selectorSet) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract matchLabels from each Required term

// MinimumHardwareAvailableAssertionForCreate asserts that catalogue has sufficient hardware to
// support the ClusterSpec during a create workflow.
//
// It does not protect against intersections or subsets so consumers should ensure a 1-2-1
// mapping between catalogue hardware and selectors.
func MinimumHardwareAvailableAssertionForCreate(catalogue *hardware.Catalogue) ClusterSpecAssertion {
	_ = "STUB: not implemented"
	return *new(ClusterSpecAssertion)
}

// Without Hardware selectors we get undesirable behavior so ensure we have them for
// all MachineConfigs.

// Build a set of required hardware counts per machine group. minimumHardwareRequirements
// will account for the same selector being specified on different groups.

// WorkerNodeHardware holds machine deployment name, replica count and hardware selector for a Tinkerbell worker node.
type WorkerNodeHardware struct {
	MachineDeploymentName string
	Replicas              int
}

// ValidatableCluster allows assertions to pull worker node and control plane information.
type ValidatableCluster interface {
	// WorkerNodeHardwareGroups retrieves a list of WorkerNodeHardwares containing MachineDeployment name,
	// replica count and hardware selector for each worker node of a ValidatableCluster.
	WorkerNodeHardwareGroups() []WorkerNodeHardware

	// ControlPlaneReplicaCount retrieves the control plane replica count of the ValidatableCluster.
	ControlPlaneReplicaCount() int

	// ClusterK8sVersion retreives the Cluster level Kubernetes version
	ClusterK8sVersion() v1alpha1.KubernetesVersion

	// WorkerGroupK8sVersion maps each worker group with its Kubernetes version.
	WorkerNodeGroupK8sVersion() map[string]v1alpha1.KubernetesVersion
}

// ValidatableTinkerbellClusterSpec wraps around the Tinkerbell ClusterSpec as a ValidatableCluster.
type ValidatableTinkerbellClusterSpec struct {
	*ClusterSpec
}

// ControlPlaneReplicaCount retrieves the ValidatableTinkerbellClusterSpec control plane replica count.
func (v *ValidatableTinkerbellClusterSpec) ControlPlaneReplicaCount() int {
	_ = "STUB: not implemented"
	return 0
}

// WorkerNodeHardwareGroups retrieves a list of WorkerNodeHardwares for a ValidatableTinkerbellClusterSpec.
func (v *ValidatableTinkerbellClusterSpec) WorkerNodeHardwareGroups() []WorkerNodeHardware {
	_ = "STUB: not implemented"
	return nil
}

// ClusterK8sVersion retrieves the Kubernetes version set at the cluster level.
func (v *ValidatableTinkerbellClusterSpec) ClusterK8sVersion() v1alpha1.KubernetesVersion {
	_ = "STUB: not implemented"
	return *new(v1alpha1.KubernetesVersion)
}

// WorkerNodeGroupK8sVersion returns each worker node group with its associated Kubernetes version.
func (v *ValidatableTinkerbellClusterSpec) WorkerNodeGroupK8sVersion() map[string]v1alpha1.KubernetesVersion {
	_ = "STUB: not implemented"
	return nil
}

// ValidatableTinkerbellCAPI wraps around the Tinkerbell control plane and worker CAPI obects as a ValidatableCluster.
type ValidatableTinkerbellCAPI struct {
	KubeadmControlPlane *controlplanev1beta2.KubeadmControlPlane
	WorkerGroups        []*clusterapi.WorkerGroup[*tinkerbellv1.TinkerbellMachineTemplate]
}

// ControlPlaneReplicaCount retrieves the ValidatableTinkerbellCAPI control plane replica count.
func (v *ValidatableTinkerbellCAPI) ControlPlaneReplicaCount() int {
	_ = "STUB: not implemented"
	return 0
}

// WorkerNodeHardwareGroups retrieves a list of WorkerNodeHardwares for a ValidatableTinkerbellCAPI.
func (v *ValidatableTinkerbellCAPI) WorkerNodeHardwareGroups() []WorkerNodeHardware {
	_ = "STUB: not implemented"
	return nil
}

// ClusterK8sVersion returns the Kubernetes version in major.minor format for a ValidatableTinkerbellCAPI.
func (v *ValidatableTinkerbellCAPI) ClusterK8sVersion() v1alpha1.KubernetesVersion {
	_ = "STUB: not implemented"
	return *new(v1alpha1.KubernetesVersion)
}

// WorkerNodeGroupK8sVersion returns each worker node group mapped to Kubernetes version in major.minor format for a ValidatableTinkerbellCAPI.
func (v *ValidatableTinkerbellCAPI) WorkerNodeGroupK8sVersion() map[string]v1alpha1.KubernetesVersion {
	_ = "STUB: not implemented"
	return nil
}

func (v *ValidatableTinkerbellCAPI) toK8sVersion(k8sversion string) v1alpha1.KubernetesVersion {
	_ = "STUB: not implemented"
	return *new(v1alpha1.KubernetesVersion)
}

// AssertionsForScaleUpDown asserts that catalogue has sufficient hardware to
// support the scaling up/down from current ClusterSpec to desired ValidatableCluster.
// nolint:gocyclo // TODO: Reduce cyclomatic complexity https://github.com/aws/eks-anywhere-internal/issues/1186
func AssertionsForScaleUpDown(catalogue *hardware.Catalogue, current ValidatableCluster, rollingUpgrade bool) ClusterSpecAssertion {
	_ = "STUB: not implemented"
	return *new(ClusterSpecAssertion)
}

// Without Hardware selectors we get undesirable behavior so ensure we have them for
// all MachineConfigs.

// Build a set of required hardware counts per machine group. minimumHardwareRequirements
// will account for the same selector being specified on different groups.

// worker node group was newly added

// ExtraHardwareAvailableAssertionForRollingUpgrade asserts that catalogue has sufficient hardware to
// support the ClusterSpec during an rolling upgrade workflow.
func ExtraHardwareAvailableAssertionForRollingUpgrade(catalogue *hardware.Catalogue, current ValidatableCluster, eksaVersionUpgrade bool) ClusterSpecAssertion {
	_ = "STUB: not implemented"
	return *new(ClusterSpecAssertion)
}

// Without Hardware selectors we get undesirable behavior so ensure we have them for
// all MachineConfigs.

// Build a set of required hardware counts per machine group. minimumHardwareRequirements
// will account for the same selector being specified on different groups.

func ensureCPHardwareAvailability(spec *ClusterSpec, hwReq MinimumHardwareRequirements) error {
	_ = "STUB: not implemented"
	return nil
}

func ensureWorkerHardwareAvailability(spec *ClusterSpec, current ValidatableCluster, hwReq MinimumHardwareRequirements, eksaVersionUpgrade bool) error {
	_ = "STUB: not implemented"
	return nil
}

// As rolling upgrades and scale up/down is not permitted in a single operation, its safe to access directly using the md name.

// ensureHardwareSelectorsSpecified ensures each machine config present in spec has a hardware
// selector or hardware affinity.
func ensureHardwareSelectorsSpecified(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// hasHardwareSelection returns true if the machine config has either HardwareSelector or HardwareAffinity set.
func hasHardwareSelection(config *v1alpha1.TinkerbellMachineConfig) bool {
	_ = "STUB: not implemented"
	return false
}

// ExtraHardwareAvailableAssertionForNodeRollOut asserts catalogue has sufficient hardware to meet minimum requirement
// and is component agnostic between Control Plane and worker nodes.
func ExtraHardwareAvailableAssertionForNodeRollOut(catalogue *hardware.Catalogue, hwReq MinimumHardwareRequirements) ClusterSpecAssertion {
	_ = "STUB: not implemented"
	return *new(ClusterSpecAssertion)
}

type missingHardwareSelectorErr struct {
	Name string
}

func (e missingHardwareSelectorErr) Error() string { _ = "STUB: not implemented"; return "" }
