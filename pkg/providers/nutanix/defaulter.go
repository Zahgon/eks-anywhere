package nutanix

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// Defaulter implements the defaulting logic for NutanixDatacenterConfig and NutanixMachineConfig.
type Defaulter struct{}

// NewDefaulter returns a new Defaulter.
func NewDefaulter() *Defaulter { _ = "STUB: not implemented"; return nil }

// SetDefaultsForDatacenterConfig sets defaults for a NutanixDatacenterConfig.
func (d *Defaulter) SetDefaultsForDatacenterConfig(dcConf anywherev1.NutanixDatacenterConfig) {
	_ = "STUB: not implemented"
	return

	// SetDefaultsForMachineConfig sets defaults for a NutanixMachineConfig.
}

func (d *Defaulter) SetDefaultsForMachineConfig(machineConf anywherev1.NutanixMachineConfig) {
	_ = "STUB: not implemented"
	return
}
