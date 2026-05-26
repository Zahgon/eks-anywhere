package cluster

import (
	"context"
)

func gitOpsEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

func processGitOps(c *Config, objects ObjectLookup) { _ = "STUB: not implemented"; return }

// GitOpsConfig will be deprecated.
// During the deprecation window, FluxConfig will be used internally
// GitOpsConfig will preserved it was in the original spec

func validateGitOps(c *Config) error { _ = "STUB: not implemented"; return nil }

func validateGitOpsNamespace(c *Config) error { _ = "STUB: not implemented"; return nil }

func setGitOpsDefaults(c *Config) error { _ = "STUB: not implemented"; return nil }

func SetDefaultFluxGitHubConfigPath(c *Config) error { _ = "STUB: not implemented"; return nil }

func getGitOps(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}
