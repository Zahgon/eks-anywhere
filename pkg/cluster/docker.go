package cluster

import (
	"context"
)

func dockerEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

// We need this conditional check as DockerDatacenter will be nil for other providers

func processDockerDatacenter(c *Config, objects ObjectLookup) { _ = "STUB: not implemented"; return }

func getDockerDatacenter(ctx context.Context, client Client, c *Config) error {
	_ = "STUB: not implemented"
	return nil
}
