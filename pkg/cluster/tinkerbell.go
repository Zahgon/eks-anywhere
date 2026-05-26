package cluster

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// tinkerbellEntry is unimplemented. Its boiler plate to mute warnings that could confuse the customer until we
// get round to implementing it.
func tinkerbellEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

// We need this conditional check as TinkerbellDatacenter will be nil for other providers

// We need this conditional check as TinkerbellMachineConfigs will be nil for other providers

func processTinkerbellDatacenter(c *Config, objects ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

func processTinkerbellTemplateConfigs(c *Config, objects ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

func processTinkerbellMachineConfig(c *Config, objects ObjectLookup, machineRef *anywherev1.Ref) {
	_ = "STUB: not implemented"
	return
}

func getTinkerbellDatacenter(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func getTinkerbellMachineAndTemplateConfigs(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}
