package docker

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/retrier"
)

// These constants are temporary since currently there is a limitation on harbor
// Harbor requires root level projects but curated packages private account currently
// doesn't have support for root level.
const (
	packageProdDomain = "783794618700.dkr.ecr.us-west-2.amazonaws.com"
	packageDevDomain  = "067575901363.dkr.ecr.us-west-2.amazonaws.com"
	publicProdECRName = "eks-anywhere"
	publicDevECRName  = "x3k6m8v0"
)

// ImageRegistryDestination implements the ImageDestination interface, writing images and tags from
// from the local docker cache to an external registry.
type ImageRegistryDestination struct {
	client    ImageTaggerPusher
	endpoint  string
	processor *ConcurrentImageProcessor
}

func NewRegistryDestination(client ImageTaggerPusher, registryEndpoint string) *ImageRegistryDestination {
	_ = "STUB: not implemented"
	return nil
}

// Write pushes images and tags from from the local docker cache to an external registry.
func (d *ImageRegistryDestination) Write(ctx context.Context, images ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// ImageOriginalRegistrySource implements the ImageSource interface, pulling images and tags from
// their original registry into the local docker cache.
type ImageOriginalRegistrySource struct {
	client    ImagePuller
	processor *ConcurrentImageProcessor
	Retrier   retrier.Retrier
}

func NewOriginalRegistrySource(client ImagePuller) *ImageOriginalRegistrySource {
	_ = "STUB: not implemented"
	return nil
}

// Load pulls images and tags from their original registry into the local docker cache.
func (s *ImageOriginalRegistrySource) Load(ctx context.Context, images ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Currently private curated packages don't have a root level project
// This method adds a root level projectName to the endpoint.
func getUpdatedEndpoint(originalEndpoint, image string) string {
	_ = "STUB: not implemented"
	return ""
}

// Curated packages are currently referenced by digest
// Docker doesn't support tagging images with digest
// This method extracts any @ in the image tag.
func removeDigestReference(image string) string { _ = "STUB: not implemented"; return "" }
