package helm

import (
	"context"
)

type ChartRegistryImporter struct {
	client             Client
	registry           string
	username, password string
	srcFolder          string
}

func NewChartRegistryImporter(client Client, srcFolder, registry, username, password string) *ChartRegistryImporter {
	_ = "STUB: not implemented"
	return nil
}

func (i *ChartRegistryImporter) Import(ctx context.Context, charts ...string) error {
	_ = "STUB: not implemented"
	return nil
}
