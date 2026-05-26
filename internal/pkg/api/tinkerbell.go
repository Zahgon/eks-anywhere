package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

type TinkerbellConfig struct {
	clusterName      string
	datacenterConfig *anywherev1.TinkerbellDatacenterConfig
	machineConfigs   map[string]*anywherev1.TinkerbellMachineConfig
	templateConfigs  map[string]*anywherev1.TinkerbellTemplateConfig
}

// TinkerbellFiller updates a TinkerbellConfig.
type TinkerbellFiller func(config TinkerbellConfig)

// TinkerbellToConfigFiller transforms a set of TinkerbellFiller's in a single ClusterConfigFiller.
func TinkerbellToConfigFiller(fillers ...TinkerbellFiller) ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(ClusterConfigFiller)
}

// updateTinkerbell updates the Tinkerbell datacenter, machine configs and
// template configs in the cluster.Config by applying all the fillers.
func updateTinkerbell(config *cluster.Config, fillers ...TinkerbellFiller) {
	_ = "STUB: not implemented"
	return
}

func WithTinkerbellServer(value string) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

func WithTinkerbellOSImageURL(value string) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

// WithTinkerbellCPMachineConfigOSImageURL sets the OSImageURL & OSFamily for control-plane machine config.
func WithTinkerbellCPMachineConfigOSImageURL(imageURL string, OSFamily anywherev1.OSFamily) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

// WithTinkerbellWorkerMachineConfigOSImageURL sets the OSImageURL & OSFamily for worker machine config.
func WithTinkerbellWorkerMachineConfigOSImageURL(imageURL string, OSFamily anywherev1.OSFamily) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

// WithHookImagesURLPath modify HookImagesURL, it's useful for airgapped testing.
func WithHookImagesURLPath(value string) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

// WithHookIsoBoot sets IsoBoot to true.
func WithHookIsoBoot() TinkerbellFiller { _ = "STUB: not implemented"; return *new(TinkerbellFiller) }

// WithHookIsoURLPath helps in setting the HookOS ISO URL value.
func WithHookIsoURLPath(url string) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

// WithTinkerbellTemplateConfig adds or updates a TinkerbellTemplateConfig.
func WithTinkerbellTemplateConfig(templateConfig *anywherev1.TinkerbellTemplateConfig) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

func WithStringFromEnvVarTinkerbell(envVar string, opt func(string) TinkerbellFiller) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

func WithOsFamilyForAllTinkerbellMachines(value anywherev1.OSFamily) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

func WithSSHAuthorizedKeyForAllTinkerbellMachines(key string) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

func WithHardwareSelectorLabels() TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

func WithTinkerbellEtcdMachineConfig() TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

// RemoveTinkerbellWorkerMachineConfig removes the worker node TinkerbellMachineConfig for single node clusters.
func RemoveTinkerbellWorkerMachineConfig() TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}

// WithStringFromEnvVarTinkerbellMachineFiller runs a TinkerbellMachineFiller function with an envVar value.
func WithStringFromEnvVarTinkerbellMachineFiller(envVar string, opt func(string) TinkerbellMachineFiller) TinkerbellMachineFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellMachineFiller)
}

// TinkerbellMachineFiller updates a TinkerbellMachineConfig.
type TinkerbellMachineFiller func(machineConfig *anywherev1.TinkerbellMachineConfig)

// WithSSHAuthorizedKeyForTinkerbellMachineConfig updates the SSHAuthorizedKey for a TinkerbellMachineConfig.
func WithSSHAuthorizedKeyForTinkerbellMachineConfig(key string) TinkerbellMachineFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellMachineFiller)
}

// WithOsFamilyForTinkerbellMachineConfig updates the OSFamily of a TinkerbellMachineConfig.
func WithOsFamilyForTinkerbellMachineConfig(value anywherev1.OSFamily) TinkerbellMachineFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellMachineFiller)
}

// WithCustomTinkerbellMachineConfig generates a TinkerbellMachineConfig from a hardware selector.
func WithCustomTinkerbellMachineConfig(selector string, machineConfigFillers ...TinkerbellMachineFiller) TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(TinkerbellFiller)
}
