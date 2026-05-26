package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

// NutanixConfig is a wrapper for the Nutanix provider spec.
type NutanixConfig struct {
	datacenterConfig *anywherev1.NutanixDatacenterConfig
	machineConfigs   map[string]*anywherev1.NutanixMachineConfig
}

type NutanixFiller func(config *NutanixConfig)

// NutanixToConfigFiller transforms a set of NutanixFiller's in a single ClusterConfigFiller.
func NutanixToConfigFiller(fillers ...NutanixFiller) ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(ClusterConfigFiller)
}

func updateNutanix(config *cluster.Config, fillers ...NutanixFiller) {
	_ = "STUB: not implemented"
	return
}

// WithNutanixStringFromEnvVar returns a NutanixFiller that sets the given string value to the given environment variable.
func WithNutanixStringFromEnvVar(envVar string, opt func(string) NutanixFiller) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixIntFromEnvVar returns a NutanixFiller that sets the given integer value to the given environment variable.
func WithNutanixIntFromEnvVar(envVar string, opt func(int) NutanixFiller) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixInt32FromEnvVar returns a NutanixFiller that sets the given int32 value to the given environment variable.
func WithNutanixInt32FromEnvVar(envVar string, opt func(int32) NutanixFiller) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixBoolFromEnvVar returns a NutanixFiller that sets the given int32 value to the given environment variable.
func WithNutanixBoolFromEnvVar(envVar string, opt func(bool) NutanixFiller) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixEndpoint returns a NutanixFiller that sets the endpoint for the Nutanix provider.
func WithNutanixEndpoint(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixPort returns a NutanixFiller that sets the port for the Nutanix provider.
func WithNutanixPort(value int) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixAdditionalTrustBundle returns a NutanixFiller that sets the additional trust bundle for the Nutanix provider.
func WithNutanixAdditionalTrustBundle(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixInsecure returns a NutanixFiller that sets the insecure for the Nutanix provider.
func WithNutanixInsecure(value bool) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixMachineMemorySize returns a NutanixFiller that sets the memory size for the Nutanix machine.
func WithNutanixMachineMemorySize(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixMachineSystemDiskSize returns a NutanixFiller that sets the system disk size for the Nutanix machine.
func WithNutanixMachineSystemDiskSize(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixMachineVCPUsPerSocket returns a NutanixFiller that sets the vCPUs per socket for the Nutanix machine.
func WithNutanixMachineVCPUsPerSocket(value int32) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixMachineBootType returns a NutanixFiller that sets the boot type for the Nutanix machine.
func WithNutanixMachineBootType(value anywherev1.NutanixBootType) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixMachineVCPUSocket returns a NutanixFiller that sets the vCPU sockets for the Nutanix machine.
func WithNutanixMachineVCPUSocket(value int32) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixMachineTemplateImageName returns a NutanixFiller that sets the image name for the Nutanix machine template.
func WithNutanixMachineTemplateImageName(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithOsFamilyForAllNutanixMachines sets the osFamily for all Nutanix machines to value.
func WithOsFamilyForAllNutanixMachines(value anywherev1.OSFamily) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixSubnetName returns a NutanixFiller that sets the subnet name for the Nutanix machine.
func WithNutanixSubnetName(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixPrismElementClusterName returns a NutanixFiller that sets the cluster name for the Nutanix machine.
func WithNutanixPrismElementClusterName(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixMachineTemplateImageUUID returns a NutanixFiller that sets the image UUID for the Nutanix machine.
func WithNutanixMachineTemplateImageUUID(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixSubnetUUID returns a NutanixFiller that sets the subnet UUID for the Nutanix machine.
func WithNutanixSubnetUUID(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixPrismElementClusterUUID returns a NutanixFiller that sets the cluster UUID for the Nutanix machine.
func WithNutanixPrismElementClusterUUID(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}

// WithNutanixSSHAuthorizedKey returns a NutanixFiller that sets the SSH authorized key for the Nutanix machine.
func WithNutanixSSHAuthorizedKey(value string) NutanixFiller {
	_ = "STUB: not implemented"
	return *new(NutanixFiller)
}
