package curatedpackages

import (
	"context"

	packagesv1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/registry"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type PackageReader struct {
	cache           *registry.Cache
	credentialStore *registry.CredentialStore
	awsRegion       string
}

// NewPackageReader create a new package reader with storage client.
func NewPackageReader(cache *registry.Cache, credentialStore *registry.CredentialStore, awsRegion string) *PackageReader {
	_ = "STUB: not implemented"
	return nil
}

// ReadImagesFromBundles and return a list of image artifacts.
func (r *PackageReader) ReadImagesFromBundles(ctx context.Context, b *releasev1.Bundles) ([]registry.Artifact, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadChartsFromBundles and return a list of chart artifacts.
func (r *PackageReader) ReadChartsFromBundles(ctx context.Context, b *releasev1.Bundles) []registry.Artifact {
	_ = "STUB: not implemented"
	return nil
}

func (r *PackageReader) getBundle(ctx context.Context, vb releasev1.VersionsBundle) (string, *packagesv1.PackageBundle, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (r *PackageReader) fetchPackagesHelmChart(bundleURI string, bundle *packagesv1.PackageBundle) []registry.Artifact {
	_ = "STUB: not implemented"
	return nil
}

func (r *PackageReader) fetchImagesFromBundle(bundleURI string, bundle *packagesv1.PackageBundle) []registry.Artifact {
	_ = "STUB: not implemented"
	return nil
}

// each package will have at least one version

// We do not have the tag right now

func removeDuplicateImages(images []registry.Artifact) []registry.Artifact {
	_ = "STUB: not implemented"
	return nil
}

func getChartRegistry(uri string) string { _ = "STUB: not implemented"; return "" }

func getImageRegistry(uri, awsRegion string) string { _ = "STUB: not implemented"; return "" }
