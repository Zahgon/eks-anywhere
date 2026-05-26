package snow

import (
	"context"
)

type ClientRegistry interface {
	Get(ctx context.Context) (AwsClientMap, error)
}

type AwsClientRegistry struct {
	deviceClientMap AwsClientMap
}

func NewAwsClientRegistry() *AwsClientRegistry { _ = "STUB: not implemented"; return nil }

// Build creates the device client map based on the filepaths specified.
// This method must be called before any Get operations.
func (b *AwsClientRegistry) Build(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *AwsClientRegistry) Get(ctx context.Context) (AwsClientMap, error) {
	_ = "STUB: not implemented"
	return *new(AwsClientMap), nil
}
