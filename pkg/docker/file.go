package docker

import (
	"context"
)

// ImageDiskSource implements the ImageSource interface, loading images and tags from
// a tarbal into the local docker cache.
type ImageDiskSource struct {
	client ImageDiskLoader
	file   string
}

func NewDiskSource(client ImageDiskLoader, file string) *ImageDiskSource {
	_ = "STUB: not implemented"
	return nil
}

// Load reads images and tags from a tarbal into the local docker cache.
func (s *ImageDiskSource) Load(ctx context.Context, images ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// ImageDiskDestination implements the ImageDestination interface, writing images and tags from
// from the local docker cache into a tarbal.
type ImageDiskDestination struct {
	client ImageDiskWriter
	file   string
}

func NewDiskDestination(client ImageDiskWriter, file string) *ImageDiskDestination {
	_ = "STUB: not implemented"
	return nil
}

// Write creates a tarball including images and tags from the the local docker cache.
func (s *ImageDiskDestination) Write(ctx context.Context, images ...string) error {
	_ = "STUB: not implemented"
	return nil
}
