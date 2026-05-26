package cluster

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

func cloudstackEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

// We need this conditional check as CloudStackDatacenter will be nil for other providers

// We need this conditional check as CloudStackMachineConfigs will be nil for other providers

func processCloudStackDatacenter(c *Config, objects ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

func processCloudStackMachineConfig(c *Config, objects ObjectLookup, machineRef *anywherev1.Ref) {
	_ = "STUB: not implemented"
	return
}

func getCloudStackDatacenter(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func getCloudStackMachineConfigs(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}
