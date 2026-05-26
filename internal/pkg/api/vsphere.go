package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

type VSphereConfig struct {
	datacenterConfig *anywherev1.VSphereDatacenterConfig
	machineConfigs   map[string]*anywherev1.VSphereMachineConfig
}

type VSphereFiller func(config VSphereConfig)

// VSphereToConfigFiller transforms a set of VSphereFiller's in a single ClusterConfigFiller.
func VSphereToConfigFiller(fillers ...VSphereFiller) ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(ClusterConfigFiller)
}

// updateVSphere updates the vSphere datacenter and machine configs in the
// cluster.Config by applying all the fillers.
func updateVSphere(config *cluster.Config, fillers ...VSphereFiller) {
	_ = "STUB: not implemented"
	return
}

func WithOsFamilyForAllMachines(value anywherev1.OSFamily) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

// WithTagsForAllMachines add provided tags to all machines.
func WithTagsForAllMachines(value []string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithNumCPUsForAllMachines(value int) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithDiskGiBForAllMachines(value int) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithMemoryMiBForAllMachines(value int) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithTLSInsecure(value bool) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithTLSThumbprint(value string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithTemplateForAllMachines(value string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

// WithMachineTemplate configs template in machine config.
func WithMachineTemplate(machineConfigName string, template string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithStoragePolicyNameForAllMachines(value string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithVSphereConfigNamespaceForAllMachinesAndDatacenter(ns string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithSSHAuthorizedKeyForAllMachines(key string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithServer(value string) VSphereFiller { _ = "STUB: not implemented"; return *new(VSphereFiller) }

func WithResourcePoolForAllMachines(value string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

// WithResourcePoolforCPMachines sets the resource pool for control plane machines to the specified value.
func WithResourcePoolforCPMachines(value string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithNetwork(value string) VSphereFiller { _ = "STUB: not implemented"; return *new(VSphereFiller) }

func WithFolderForAllMachines(value string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithDatastoreForAllMachines(value string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithDatacenter(value string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

// WithCloneModeForAllMachines sets the CloneMode for all VSphereMachineConfigs.
func WithCloneModeForAllMachines(value anywherev1.CloneMode) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

// WithNTPServersForAllMachines sets NTP servers for all VSphereMachineConfigs.
func WithNTPServersForAllMachines(servers []string) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

// WithBottlerocketConfigurationForAllMachines sets Bottlerocket configuration for all VSphereMachineConfigs.
func WithBottlerocketConfigurationForAllMachines(value *anywherev1.BottlerocketConfiguration) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithVSphereStringFromEnvVar(envVar string, opt func(string) VSphereFiller) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithVSphereBoolFromEnvVar(envVar string, opt func(bool) VSphereFiller) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

func WithVSphereMachineConfig(name string, fillers ...VSphereMachineConfigFiller) VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}

// RemoveEtcdVsphereMachineConfig removes the etcd VSphereMachineConfig from the cluster spec.
func RemoveEtcdVsphereMachineConfig() VSphereFiller {
	_ = "STUB: not implemented"
	return *new(VSphereFiller)
}
