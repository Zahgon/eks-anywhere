package tinkerbell

import (
	tinkv1alpha1 "github.com/tinkerbell/tink/api/v1alpha1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/networkutils"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
)

func validateOsFamily(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

func validateK8sVersionForBottleRocketOS(kubernetesVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateUpgradeRolloutStrategy(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

func validateAutoScalerDisabledForInPlace(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// We do not support different strategy types for Inplace between CP and worker nodes so it is okay to check only CP

func validateOSImageURL(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

func validateK8sVersionInOSImageURLs(spec *ClusterSpec) error {
	_ = "STUB: not implemented"
	// If the user specifies the OSImageURL via the datacenter config then ensure all kube versions specified
	// on the cluster config are specified in the OSImageURL as the user could technically use a single image.
	//
	// When the user specifies OSImageURLs on each individual machine config (typical for modular upgrades) ensure
	// each machine config OSImageURL specifies the Kubernetes version. We don't explicitly take into consideration
	// the fact control plane, etcd and worker node groups can all reference the same machine config. If 2 components
	// specify different kube versions this will ensure both are present in the image URL (as above).
	return nil
}

// For Bottlerocket we vend images but we still allow the user to specify them if they wish. We only want
// to default machine config OSImageURLs if the datacenter config doesn't specify one and we default
// to whatever is in the bundle.
//
// TODO: Investigate how we could refactor our logic to make this unnecessary.
//
// We validate elsewhere that all machine configs specify the same OSFamily so we can rely on the
// control plane machine config only for the need to default OSImageURLs.

func defaultBottlerocketOSImageURLs(spec *ClusterSpec) { _ = "STUB: not implemented"; return }

func containsK8sVersion(imageURL, k8sVersion string) bool { _ = "STUB: not implemented"; return false }

// we set the containsK8sVersion to false if the OS image URL does not contain the specified kubernetes version.
// For ex if the kubernetes version is 1.23,
// the image url should include 1.23 or 1-23, 1_23 or 123 i.e. ubuntu-1-23.gz or similar in the string.

func validateISOURL(spec *ClusterSpec) error { _ = "STUB: not implemented"; return nil }

func validateMachineRefExists(
	ref *v1alpha1.Ref,
	machineConfigs map[string]*v1alpha1.TinkerbellMachineConfig,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMachineConfigNamespacesMatchDatacenterConfig(
	datacenterConfig *v1alpha1.TinkerbellDatacenterConfig,
	machineConfigs map[string]*v1alpha1.TinkerbellMachineConfig,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIPUnused(client networkutils.NetClient, ip string) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePortsAvailable(client networkutils.NetClient, host string) error {
	_ = "STUB: not implemented"
	return nil
}

func getPortsUnavailable(client networkutils.NetClient, host string) []string {
	_ = "STUB: not implemented"
	return nil
}

// GetSelectorsFromMachineConfig extracts hardware selectors from a machine config.
// If HardwareAffinity is set, it extracts matchLabels from Required terms.
// Otherwise, it returns the HardwareSelector.
func GetSelectorsFromMachineConfig(config *v1alpha1.TinkerbellMachineConfig) []v1alpha1.HardwareSelector {
	_ = "STUB: not implemented"
	return nil
}

// minimumHardwareRequirement defines the minimum requirement for a hardware selector.
type minimumHardwareRequirement struct {
	// MinCount is the minimum number of hardware required to satisfy the requirement
	MinCount int
	// Selector defines what labels should be present on Hardware to consider it eligable for
	// this requirement.
	Selector v1alpha1.HardwareSelector
	// count is used internally by validation to sum the actual available hardware.
	count int
}

// MinimumHardwareRequirements is a collection of minimumHardwareRequirement instances.
// it stores requirements in a map where the key is derived from selectors. This ensures selectors
// specifying the same key-value pairs are combined.
type MinimumHardwareRequirements map[string]*minimumHardwareRequirement

// Add a minimumHardwareRequirement to r.
func (r *MinimumHardwareRequirements) Add(selector v1alpha1.HardwareSelector, min int) error {
	_ = "STUB: not implemented"
	return nil
}

// validateMinimumHardwareRequirements validates all requirements can be satisfied using hardware
// registered with catalogue.
func validateMinimumHardwareRequirements(requirements MinimumHardwareRequirements, catalogue *hardware.Catalogue) error {
	_ = "STUB: not implemented"
	// Count all hardware that meets the selector requirements for each requirement.
	// This does not consider whether or not a piece of hardware is selectable by multiple
	// selectors. That requires a different validation ideally run before this one.
	return nil
}

// Validate counts of hardware meet the minimum required count.

// validateHardwareSatisfiesOnlyOneSelector ensures hardware in allHardware meets one and only one
// selector in selectors. selectors uses the selectorSet construct to ensure we don't
// operate on duplicate selectors given a selector can be re-used among groups as they may reference
// the same TinkerbellMachineConfig.
func validateHardwareSatisfiesOnlyOneSelector(allHardware []*tinkv1alpha1.Hardware, selectors selectorSet) error {
	_ = "STUB: not implemented"
	return nil
}

// selectorSet defines a set of selectors. Selectors should be added using the Add method to ensure
// deterministic key generation. The construct is useful to avoid treating selectors that are the
// same as different.
type selectorSet map[string]v1alpha1.HardwareSelector

// Add adds selector to ss.
func (ss *selectorSet) Add(selector v1alpha1.HardwareSelector) error {
	_ = "STUB: not implemented"
	return nil
}

func getMatchingHardwareSelectors(
	hw *tinkv1alpha1.Hardware,
	selectors selectorSet,
) []v1alpha1.HardwareSelector {
	_ = "STUB: not implemented"
	return nil
}

func getHardwareSelectorsAsStrings(selectors []v1alpha1.HardwareSelector) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
