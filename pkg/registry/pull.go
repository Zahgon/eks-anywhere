package registry

import (
	"context"
)

// PullBytes a resource from the registry.
func PullBytes(ctx context.Context, sc StorageClient, artifact Artifact) (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
