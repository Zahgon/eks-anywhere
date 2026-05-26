package cluster

import (
	eksdv1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/manifests/bundles"
	"github.com/aws/eks-anywhere/pkg/version"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// FileSpecBuilder allows to build [Spec] by reading from files.
type FileSpecBuilder struct {
	// TODO(g-gaston): this is very much a CLI thing. Move to `pkg/cli` when available.
	reader              manifests.FileReader
	cliVersion          version.Info
	releasesManifestURL string
	bundlesManifestURL  string
}

// FileSpecBuilderOpt allows to configure [FileSpecBuilder].
type FileSpecBuilderOpt func(*FileSpecBuilder)

// WithReleasesManifest configures the URL to read the Releases manifest.
func WithReleasesManifest(url string) FileSpecBuilderOpt {
	_ = "STUB: not implemented"
	return *new(FileSpecBuilderOpt)
}

// WithOverrideBundlesManifest configures the URL to read the Bundles manifest.
// This overrides the Bundles declared in the Releases so reading the Releases
// manifest is skipped.
func WithOverrideBundlesManifest(url string) FileSpecBuilderOpt {
	_ = "STUB: not implemented"
	return *new(FileSpecBuilderOpt)
}

// NewFileSpecBuilder builds a new [FileSpecBuilder].
// cliVersion is used to chose the right Bundles from the the Release manifest.
func NewFileSpecBuilder(reader manifests.FileReader, cliVersion version.Info, opts ...FileSpecBuilderOpt) FileSpecBuilder {
	_ = "STUB: not implemented"
	return *new(FileSpecBuilder)
}

// Build constructs a new [Spec] by reading the cluster config in yaml from a file and
// Releases, Bundles and EKS-D manifests from the configured URLs.
func (b FileSpecBuilder) Build(clusterConfigURL string) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getAllEksdReleases(cluster *v1alpha1.Cluster, bundlesManifest *releasev1.Bundles, reader bundles.Reader) ([]eksdv1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getEksdReleases(version v1alpha1.KubernetesVersion, bundlesManifest *releasev1.Bundles, reader bundles.Reader) (*eksdv1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b FileSpecBuilder) getConfig(clusterConfigURL string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b FileSpecBuilder) createManifestReader() *manifests.Reader {
	_ = "STUB: not implemented"
	return nil
}

func (b FileSpecBuilder) getBundles(manifestReader *manifests.Reader) (*releasev1.Bundles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b FileSpecBuilder) getEksaRelease(mReader *manifests.Reader) (*releasev1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// this shouldn't return an error at this point due to getBundles performing similar operations prior to this call
}

// When using bundles-override or a custom bundle, a fake EksaRelease can be used since using a custom bundle
// is like creating a new EKS-A version.

func buildEKSARelease(release *releasev1.EksARelease, bundle *releasev1.Bundles) *releasev1.EKSARelease {
	_ = "STUB: not implemented"
	return nil
}
