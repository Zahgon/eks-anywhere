package cluster

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

func snowEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

func processSnowDatacenter(c *Config, objects ObjectLookup) { _ = "STUB: not implemented"; return }

func snowIPPoolsProcessor(c *Config, o ObjectLookup) { _ = "STUB: not implemented"; return }

func processSnowIPPool(c *Config, objects ObjectLookup, ipPoolRef *anywherev1.Ref) {
	_ = "STUB: not implemented"
	return
}

func processSnowMachineConfig(c *Config, objects ObjectLookup, machineRef *anywherev1.Ref) {
	_ = "STUB: not implemented"
	return
}

func SetSnowMachineConfigsAnnotations(c *Config) error { _ = "STUB: not implemented"; return nil }

func getSnowDatacenter(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func getSnowMachineConfigsAndIPPools(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func getSnowIPPools(ctx context.Context, client Client, c *Config, machine *anywherev1.SnowMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func getSnowIdentitySecret(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}

// SetSnowDatacenterIndentityRefDefault sets a default secret as the identity reference
// The secret will need to be created by the CLI flow as it's not provided by the user
// This only runs in CLI. snowDatacenterConfig.SetDefaults() will run in both CLI and webhook.
func SetSnowDatacenterIndentityRefDefault(s *anywherev1.SnowDatacenterConfig) {
	_ = "STUB: not implemented"
	return
}

// ValidateSnowMachineRefExists checks the cluster spec machine refs and makes sure
// the snowmachineconfig object exists for each ref with kind == snowmachineconfig.
func ValidateSnowMachineRefExists(c *Config) error { _ = "STUB: not implemented"; return nil }

func validateSnowUnstackedEtcd(c *Config) error { _ = "STUB: not implemented"; return nil }
