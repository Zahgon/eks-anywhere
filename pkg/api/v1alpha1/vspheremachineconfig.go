package v1alpha1

const (
	VSphereMachineConfigKind = "VSphereMachineConfig"
	DefaultVSphereDiskGiB    = 25
	DefaultVSphereNumCPUs    = 2
	DefaultVSphereMemoryMiB  = 8192
	DefaultVSphereOSFamily   = Bottlerocket
)

// Used for generating yaml for generate clusterconfig command.
func NewVSphereMachineConfigGenerate(name string) *VSphereMachineConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *VSphereMachineConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *VSphereMachineConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *VSphereMachineConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func setVSphereMachineConfigDefaults(machineConfig *VSphereMachineConfig) {
	_ = "STUB: not implemented"
	return
}

func validateVSphereMachineConfig(config *VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateVSphereMachineConfigHasTemplate(config *VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}
