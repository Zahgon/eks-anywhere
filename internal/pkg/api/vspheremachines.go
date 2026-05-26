package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type VSphereMachineConfigFiller func(m *anywherev1.VSphereMachineConfig)

func FillVSphereMachineConfig(m *anywherev1.VSphereMachineConfig, fillers ...VSphereMachineConfigFiller) {
	_ = "STUB: not implemented"
	return
}

func WithVSphereMachineDefaultValues() VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

func WithDatastore(value string) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

func WithFolder(value string) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

// WithTags add provided tags to all machines.
func WithTags(value []string) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

func WithResourcePool(value string) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

func WithStoragePolicyName(value string) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

func WithTemplate(value string) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

func WithSSHKey(value string) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

// WithStringFromEnvVar returns a VSphereMachineConfigFiller function with the value from an envVar passed to it.
func WithStringFromEnvVar(envVar string, opt func(string) VSphereMachineConfigFiller) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

// WithNetworks sets the Networks field for a VSphereMachineConfig.
func WithNetworks(networks []string) VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(VSphereMachineConfigFiller)
}

func setSSHKeyForFirstUser(m *anywherev1.VSphereMachineConfig, key string) {
	_ = "STUB: not implemented"
	return
}
