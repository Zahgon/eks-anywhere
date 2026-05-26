package v1alpha1

import (
	"regexp"

	snowv1 "github.com/aws/eks-anywhere/pkg/providers/snow/api/v1beta1"
)

const (
	SnowMachineConfigKind                   = "SnowMachineConfig"
	DefaultSnowSSHKeyName                   = ""
	DefaultSnowInstanceType                 = "sbe-c.large"
	DefaultSnowPhysicalNetworkConnectorType = SFPPlus
	DefaultOSFamily                         = Ubuntu
	MinimumContainerVolumeSizeUbuntu        = 8
	MinimumContainerVolumeSizeBottlerocket  = 25
	MinimumNonRootVolumeSize                = 8
)

var snowInstanceTypesRegex = regexp.MustCompile(`^sbe-[cg]\.\d*x?large$`)

// NewSnowMachineConfigGenerate generates snowMachineConfig example for generate clusterconfig command.
func NewSnowMachineConfigGenerate(name string) *SnowMachineConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (s *SnowMachineConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (s *SnowMachineConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (s *SnowMachineConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func validateSnowMachineConfig(config *SnowMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSnowMachineConfigInstanceType(instanceType string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSnowMachineConfigContainerVolume(config *SnowMachineConfig) error {
	_ = "STUB: not implemented"
	// The Bottlerocket AWS Variant AMI only has 2 Gi of data volume, which is insufficient to store EKS-A and user container volumes.
	// Thus the ContainersVolume is required and its size must be no smaller than 25 Gi.
	return nil
}

func validateSnowMachineConfigNonRootVolumes(volumes []*snowv1.Volume) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSnowMachineConfigNetwork(network SnowNetwork) error {
	_ = "STUB: not implemented"
	return nil
}

func setSnowMachineConfigDefaults(config *SnowMachineConfig) { _ = "STUB: not implemented"; return }
