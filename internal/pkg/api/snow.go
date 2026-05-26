package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

type SnowConfig struct {
	datacenterConfig *anywherev1.SnowDatacenterConfig
	machineConfigs   map[string]*anywherev1.SnowMachineConfig
	ipPools          map[string]*anywherev1.SnowIPPool
}

type SnowFiller func(config SnowConfig)

// SnowToConfigFiller transforms a set of SnowFiller's in a single ClusterConfigFiller.
func SnowToConfigFiller(fillers ...SnowFiller) ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(ClusterConfigFiller)
}

func updateSnow(config *cluster.Config, fillers ...SnowFiller) { _ = "STUB: not implemented"; return }

func WithSnowStringFromEnvVar(envVar string, opt func(string) SnowFiller) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

func WithSnowAMIIDForAllMachines(id string) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

// WithSnowInstanceTypeForAllMachines specifies an instance type for all the snow machine configs.
func WithSnowInstanceTypeForAllMachines(instanceType string) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

func WithSnowPhysicalNetworkConnectorForAllMachines(connectorType anywherev1.PhysicalNetworkConnectorType) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

func WithSnowSshKeyNameForAllMachines(keyName string) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

func WithSnowDevicesForAllMachines(devices string) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

func WithSnowMachineConfig(name string, fillers ...SnowMachineConfigFiller) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

// WithOsFamilyForAllSnowMachines sets the OSFamily in the SnowMachineConfig.
func WithOsFamilyForAllSnowMachines(value anywherev1.OSFamily) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

// WithChangeForAllSnowMachines applies the same change to all SnowMachineConfigs.
func WithChangeForAllSnowMachines(change SnowMachineConfigFiller) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}

// WithSnowIPPool sets a SnowIPPool.
func WithSnowIPPool(name, ipStart, ipEnd, gateway, subnet string) SnowFiller {
	_ = "STUB: not implemented"
	return *new(SnowFiller)
}
