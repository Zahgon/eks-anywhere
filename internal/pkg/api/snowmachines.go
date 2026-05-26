package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type SnowMachineConfigFiller func(m *anywherev1.SnowMachineConfig)

func FillSnowMachineConfig(m *anywherev1.SnowMachineConfig, fillers ...SnowMachineConfigFiller) {
	_ = "STUB: not implemented"
	return
}

func WithSnowMachineDefaultValues() SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}

func WithSnowAMIID(id string) SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}

// WithSnowInstanceType specifies an instance type for the snow machine config.
func WithSnowInstanceType(instanceType string) SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}

func WithSnowPhysicalNetworkConnector(connectorType anywherev1.PhysicalNetworkConnectorType) SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}

func WithSnowSshKeyName(keyName string) SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}

func WithSnowDevices(devices string) SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}

// WithDHCP configures one single primary DNI using DHCP for IP allocation.
func WithDHCP() SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}

// WithStaticIP configures one single primary DNI using static ip for IP allocation.
func WithStaticIP(poolName string) SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}

// WithSnowContainersVolumeSize sets the container volume size for a SnowMachineConfig.
func WithSnowContainersVolumeSize(size int64) SnowMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(SnowMachineConfigFiller)
}
