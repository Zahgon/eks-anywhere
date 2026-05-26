package cluster

import (
	"context"
)

func fluxEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

func processFlux(c *Config, objects ObjectLookup) { _ = "STUB: not implemented"; return }

func validateFlux(c *Config) error { _ = "STUB: not implemented"; return nil }

func validateFluxNamespace(c *Config) error { _ = "STUB: not implemented"; return nil }

func setFluxDefaults(c *Config) error { _ = "STUB: not implemented"; return nil }

func SetDefaultFluxConfigPath(c *Config) error { _ = "STUB: not implemented"; return nil }

func getFluxConfig(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}
