package manifests

import (
	"context"

	eksdv1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type FileReader interface {
	ReadFile(url string) ([]byte, error)
}

type Reader struct {
	FileReader
	releasesManifestURL string
}

type ReaderOpt func(*Reader)

func WithReleasesManifest(manifestURL string) ReaderOpt {
	_ = "STUB: not implemented"
	return *new(ReaderOpt)
}

func NewReader(filereader FileReader, opts ...ReaderOpt) *Reader {
	_ = "STUB: not implemented"
	return nil
}

// ReadReleaseForVersion returns an EksaRelease based on a version.
func (r *Reader) ReadReleaseForVersion(version string) (*releasev1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadBundlesForVersion returns a Bundle based on the version.
func (r *Reader) ReadBundlesForVersion(version string) (*releasev1.Bundles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) ReadEKSD(eksaVersion, kubeVersion string) (*eksdv1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) ReadImages(eksaVersion string) ([]releasev1.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) ReadImagesFromBundles(_ context.Context, b *releasev1.Bundles) ([]releasev1.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) ReadCharts(eksaVersion string) ([]releasev1.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) ReadChartsFromBundles(ctx context.Context, b *releasev1.Bundles) []releasev1.Image {
	_ = "STUB: not implemented"
	return nil
}
