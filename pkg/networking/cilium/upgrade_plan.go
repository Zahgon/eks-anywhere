package cilium

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	// PolicyEnforcementConfigMapKey is the key used in the "cilium-config" ConfigMap to
	// store the value for the PolicyEnforcementMode.
	PolicyEnforcementConfigMapKey = "enable-policy"

	// PolicyEnforcementComponentName is the ConfigComponentUpdatePlan name for the
	// PolicyEnforcement configuration component.
	PolicyEnforcementComponentName = "PolicyEnforcementMode"

	// EgressMasqueradeInterfacesMapKey is the key used in the "cilium-config" ConfigMap to
	// store the value for the EgressMasqueradeInterfaces.
	EgressMasqueradeInterfacesMapKey = "egress-masquerade-interfaces"

	// EgressMasqueradeInterfacesComponentName is the ConfigComponentUpdatePlan name for the
	// egressMasqueradeInterfaces configuration component.
	EgressMasqueradeInterfacesComponentName = "EgressMasqueradeInterfaces"

	// CniExclusiveConfigMapKey is the key used in the "cilium-config" ConfigMap to
	// store the value for the CniExclusive.
	CniExclusiveConfigMapKey = "cni-exclusive"

	// CniExclusiveComponentName is the ConfigComponentUpdatePlan name for the
	// CniExclusive configuration component.
	CniExclusiveComponentName = "CniExclusive"
)

// UpgradePlan contains information about a Cilium installation upgrade.
type UpgradePlan struct {
	DaemonSet VersionedComponentUpgradePlan
	Operator  VersionedComponentUpgradePlan
	ConfigMap ConfigUpdatePlan
}

// Needed determines if an upgrade is needed or not
// Returns true if any of the installation components needs an upgrade.
func (c UpgradePlan) Needed() bool { _ = "STUB: not implemented"; return false }

// VersionUpgradeNeeded determines if a version upgrade is needed or not
// Returns true if any of the installation components needs an upgrade.
func (c UpgradePlan) VersionUpgradeNeeded() bool { _ = "STUB: not implemented"; return false }

// ConfigUpdateNeeded determines if an upgrade is needed on the cilium config or not.
func (c UpgradePlan) ConfigUpdateNeeded() bool { _ = "STUB: not implemented"; return false }

// Reason returns the reason why an upgrade might be needed
// If no upgrade needed, returns empty string
// For multiple components with needed upgrades, it composes their reasons into one.
func (c UpgradePlan) Reason() string { _ = "STUB: not implemented"; return "" }

// VersionedComponentUpgradePlan contains upgrade information for a Cilium versioned component.
type VersionedComponentUpgradePlan struct {
	UpgradeReason string
	OldImage      string
	NewImage      string
}

// Needed determines if an upgrade is needed or not.
func (c VersionedComponentUpgradePlan) Needed() bool { _ = "STUB: not implemented"; return false }

// reason returns the reason for the upgrade if needed.
// If upgrade is not needed, it returns an empty string.
func (c VersionedComponentUpgradePlan) reason() string { _ = "STUB: not implemented"; return "" }

// ConfigUpdatePlan contains update information for the Cilium config.
type ConfigUpdatePlan struct {
	UpdateReason string
	Components   []ConfigComponentUpdatePlan
}

// Needed determines if an upgrade is needed or not.
func (c ConfigUpdatePlan) Needed() bool { _ = "STUB: not implemented"; return false }

// reason returns the reason for the upgrade if needed.
// If upgrade is not needed, it returns an empty string.
func (c ConfigUpdatePlan) reason() string { _ = "STUB: not implemented"; return "" }

// generateUpdateReasonFromComponents reads the update reasons for the components
// and generates a compounded update reason. This is not thread safe.
func (c *ConfigUpdatePlan) generateUpdateReasonFromComponents() { _ = "STUB: not implemented"; return }

// ConfigComponentUpdatePlan contains update information for a Cilium config component.
type ConfigComponentUpdatePlan struct {
	Name               string
	UpdateReason       string
	OldValue, NewValue string
}

// BuildUpgradePlan generates the upgrade plan information for a cilium installation by comparing it
// with a desired cluster Spec.
func BuildUpgradePlan(installation *Installation, clusterSpec *cluster.Spec) UpgradePlan {
	_ = "STUB: not implemented"
	return *new(UpgradePlan)
}

func daemonSetUpgradePlan(ds *appsv1.DaemonSet, clusterSpec *cluster.Spec) VersionedComponentUpgradePlan {
	_ = "STUB: not implemented"
	return *new(VersionedComponentUpgradePlan)
}

func operatorUpgradePlan(operator *appsv1.Deployment, clusterSpec *cluster.Spec) VersionedComponentUpgradePlan {
	_ = "STUB: not implemented"
	return *new(VersionedComponentUpgradePlan)
}

func configMapUpgradePlan(configMap *corev1.ConfigMap, clusterSpec *cluster.Spec) ConfigUpdatePlan {
	_ = "STUB: not implemented"
	return *new(ConfigUpdatePlan)
}

// Check for cniExclusive parameter changes

// Default value when not specified

// If the field is not present in the config, assume the default value of "true"

// ChangeDiff returns the change diff between the current and new cluster specs.
func ChangeDiff(currentSpec, newSpec *cluster.Spec) *types.ChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

func ciliumChangeDiff(currentSpec, newSpec *cluster.Spec) *types.ChangeDiff {
	_ = "STUB: not implemented"
	return nil
}
