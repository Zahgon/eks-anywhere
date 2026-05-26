package snow

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	defaultAwsSshKeyName       = "eksa-default"
	snowballMinSoftwareVersion = 102
	minimumVCPU                = 2
)

// Validator includes a client registry that maintains a snow device aws client map,
// and a local imds service that is used to fetch metadata of the host instance.
type Validator struct {
	// clientRegistry maintains a device aws client mapping.
	clientRegistry ClientRegistry

	// imds is a local imds client built with the default aws config. This imds client can only
	// interact with the local instance metata service.
	imds LocalIMDSClient
}

// ValidatorOpt updates an Validator.
type ValidatorOpt func(*Validator)

// WithIMDS returns a ValidatorOpt that sets the imds client.
func WithIMDS(imds LocalIMDSClient) ValidatorOpt {
	_ = "STUB: not implemented"
	return *new(ValidatorOpt)
}

// NewValidator creates a snow validator.
func NewValidator(clientRegistry ClientRegistry, opts ...ValidatorOpt) *Validator {
	_ = "STUB: not implemented"
	return nil
}

// ValidateEC2SshKeyNameExists validates the ssh key existence in each device in the device list.
func (v *Validator) ValidateEC2SshKeyNameExists(ctx context.Context, m *v1alpha1.SnowMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateEC2ImageExistsOnDevice validates the ami id (if specified) existence in each device in the device list.
func (v *Validator) ValidateEC2ImageExistsOnDevice(ctx context.Context, m *v1alpha1.SnowMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateDeviceIsUnlocked verifies if all snow devices in the device list are unlocked.
func (v *Validator) ValidateDeviceIsUnlocked(ctx context.Context, m *v1alpha1.SnowMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateInstanceTypeInDevice(ctx context.Context, client AwsClient, instanceType, deviceIP string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateInstanceType validates whether the instance type is compatible to run in each device.
func (v *Validator) ValidateInstanceType(ctx context.Context, m *v1alpha1.SnowMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateDeviceSoftware validates whether the snow software is compatible to run eks-a in each device.
func (v *Validator) ValidateDeviceSoftware(ctx context.Context, m *v1alpha1.SnowMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateControlPlaneIP checks whether the control plane ip is valid for creating a snow cluster.
func (v *Validator) ValidateControlPlaneIP(ctx context.Context, controlPlaneIP string) error {
	_ = "STUB: not implemented"
	return nil
}

// the admin instance is not running inside snow devices or doesn't have a public IP
