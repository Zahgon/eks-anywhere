package v1alpha1

// DefaultCloudStackUser is the default CloudStackMachingConfig username.
const DefaultCloudStackUser = "capc"

// CloudStackMachineConfigKind is the kind value for a CloudStackMachineConfig.
const CloudStackMachineConfigKind = "CloudStackMachineConfig"

// Taken from https://github.com/shapeblue/cloudstack/blob/08bb4ad9fea7e422c3d3ac6d52f4670b1e89eed7/api/src/main/java/com/cloud/vm/VmDetailConstants.java
// These fields should be modeled separately in eks-a and not used by the additionalDetails cloudstack VM field.
var restrictedUserCustomDetails = [...]string{
	"keyboard", "cpu.corespersocket", "rootdisksize", "boot.mode", "nameonhypervisor",
	"nicAdapter", "rootDiskController", "dataDiskController", "svga.vramSize", "nestedVirtualizationFlag", "ramReservation",
	"hypervisortoolsversion", "platform", "timeoffset", "kvm.vnc.port", "kvm.vnc.address", "video.hardware", "video.ram",
	"smc.present", "firmware", "cpuNumber", "cpuSpeed", "memory", "cpuOvercommitRatio", "memoryOvercommitRatio",
	"Message.ReservedCapacityFreed.Flag", "deployvm", "SSH.PublicKey", "SSH.KeyPairNames", "password", "Encrypted.Password",
	"configDriveLocation", "nic", "network", "ip4Address", "ip6Address", "disk", "diskOffering", "configurationId",
	"keypairnames", "controlNodeLoginUser",
}

// Used for generating yaml for generate clusterconfig command.
func NewCloudStackMachineConfigGenerate(name string) *CloudStackMachineConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *CloudStackMachineConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *CloudStackMachineConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *CloudStackMachineConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func validateCloudStackMachineConfig(machineConfig *CloudStackMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAffinityConfig(machineConfig *CloudStackMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}
