package registry

import (
	"context"
	"sync"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	orasregistry "oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
)

// OCIRegistryClient storage client for an OCI registry.
type OCIRegistryClient struct {
	StorageContext
	initialized sync.Once
	registry    *remote.Registry
}

var _ StorageClient = (*OCIRegistryClient)(nil)

// NewOCIRegistry create an OCI registry client.
func NewOCIRegistry(context StorageContext) *OCIRegistryClient {
	_ = "STUB: not implemented"
	return nil
}

// Init registry configuration.
func (or *OCIRegistryClient) Init() error { _ = "STUB: not implemented"; return nil }

// #nosec G402

// GetHost for registry host.
func (or *OCIRegistryClient) GetHost() string {
	_ = "STUB: not implemented"

	// SetProject for registry destination.
	return ""
}

func (or *OCIRegistryClient) SetProject(project string) { _ = "STUB: not implemented"; return }

// Destination of this storage registry.
func (or *OCIRegistryClient) Destination(image Artifact) string {
	_ = "STUB: not implemented"
	return ""
}

// GetStorage object based on repository.
func (or *OCIRegistryClient) GetStorage(ctx context.Context, artifact Artifact) (repo orasregistry.Repository, err error) {
	_ = "STUB: not implemented"
	return *new(orasregistry.Repository), nil
}

// Resolve the location of the source repository given the image.
func (or *OCIRegistryClient) Resolve(ctx context.Context, srcStorage orasregistry.Repository, versionedImage string) (desc ocispec.Descriptor, err error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}

// FetchBytes a resource from the registry.
func (or *OCIRegistryClient) FetchBytes(ctx context.Context, srcStorage orasregistry.Repository, artifact Artifact) (ocispec.Descriptor, []byte, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil, nil
}

// FetchBlob get named blob.
func (or *OCIRegistryClient) FetchBlob(ctx context.Context, srcStorage orasregistry.Repository, descriptor ocispec.Descriptor) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CopyGraph copy manifest and all blobs to destination.
func (or *OCIRegistryClient) CopyGraph(ctx context.Context, srcStorage orasregistry.Repository, srcRef string, dstStorage orasregistry.Repository, dstRef string) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}

// Tag an image.
func (or *OCIRegistryClient) Tag(ctx context.Context, dstStorage orasregistry.Repository, desc ocispec.Descriptor, tag string) error {
	_ = "STUB: not implemented"
	return nil
}
