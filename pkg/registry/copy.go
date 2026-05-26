package registry

import (
	"context"
)

// Copy an image from a source to a destination.
func Copy(ctx context.Context, srcClient StorageClient, dstClient StorageClient, image Artifact) (err error) {
	_ = "STUB: not implemented"
	return nil
}
