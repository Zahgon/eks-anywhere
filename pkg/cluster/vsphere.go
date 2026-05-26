package cluster

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

func vsphereEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

// We need this conditional check as VSphereDatacenter will be nil for other providers

// We need this conditional check as VSphereMachineConfigs will be nil for other providers

func processVSphereDatacenter(c *Config, objects ObjectLookup) { _ = "STUB: not implemented"; return }

func processVSphereMachineConfig(c *Config, objects ObjectLookup, machineRef *anywherev1.Ref) {
	_ = "STUB: not implemented"
	return
}

func getVSphereDatacenter(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func getVSphereMachineConfigs(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}
