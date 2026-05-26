package cluster

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

func nutanixEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

// We need this conditional check as NutanixDatacenter will be nil for other providers

// We need this conditional check as NutanixMachineConfigs will be nil for other providers

func processNutanixDatacenter(c *Config, objects ObjectLookup) { _ = "STUB: not implemented"; return }

func processNutanixMachineConfig(c *Config, objects ObjectLookup, machineRef *anywherev1.Ref) {
	_ = "STUB: not implemented"
	return
}

func getNutanixDatacenter(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func getNutanixMachineConfigs(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}
