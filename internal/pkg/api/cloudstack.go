package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

type CloudStackConfig struct {
	datacenterConfig *anywherev1.CloudStackDatacenterConfig
	machineConfigs   map[string]*anywherev1.CloudStackMachineConfig
}

type CloudStackFiller func(config CloudStackConfig)

// CloudStackToConfigFiller transforms a set of CloudStackFiller's in a single ClusterConfigFiller.
func CloudStackToConfigFiller(fillers ...CloudStackFiller) ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(ClusterConfigFiller)
}

func updateCloudStack(config *cluster.Config, fillers ...CloudStackFiller) {
	_ = "STUB: not implemented"
	return
}

func WithCloudStackComputeOfferingForAllMachines(value string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackAz(az anywherev1.CloudStackAvailabilityZone) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func RemoveCloudStackAzs() CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackAffinityGroupIds(value []string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithUserCustomDetails(value map[string]string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithSymlinks(value map[string]string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackTemplateForAllMachines(value string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackConfigNamespace(ns string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackSSHAuthorizedKey(value string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackDomain(value string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackAccount(value string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

// WithCloudStackCredentialsRef returns a CloudStackFiller that updates the edentialsRef of all availability zones.
func WithCloudStackCredentialsRef(value string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackStringFromEnvVar(envVar string, opt func(string) CloudStackFiller) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackAzFromEnvVars(cloudstackAccountVar, cloudstackDomainVar, cloudstackZoneVar, cloudstackCredentialsVar, cloudstackNetworkVar, cloudstackManagementServerVar string, opt func(zone anywherev1.CloudStackAvailabilityZone) CloudStackFiller) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

func WithCloudStackMachineConfig(name string, fillers ...CloudStackMachineConfigFiller) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}

// WithCloudStackConfigNamespaceForAllMachinesAndDatacenter sets the namespace for all Machines and Datacenter objects.
func WithCloudStackConfigNamespaceForAllMachinesAndDatacenter(ns string) CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackFiller)
}
