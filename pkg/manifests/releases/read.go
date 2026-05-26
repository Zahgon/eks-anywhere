package releases

import (
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// manifestURL holds the url to the eksa releases manifest
// this is injected at build time, this is just a sane default for development.
var manifestURL = "https://dev-release-assets.eks-anywhere.model-rocket.aws.dev/eks-a-release.yaml"

func ManifestURL() string { _ = "STUB: not implemented"; return "" }

type Reader interface {
	ReadFile(url string) ([]byte, error)
}

func ReadReleases(reader Reader) (*releasev1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadReleasesFromURL(reader Reader, url string) (*releasev1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBundleManifestURL fetches the bundle manifest URL pertaining to this
// release version of EKS Anywhere.
func GetBundleManifestURL(reader Reader, version string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// BundleManifestURL returns the  Bundles manifest URL for the release matched by the provided
// version. If no release is found for the version, an error is returned.
func BundleManifestURL(releases *releasev1.Release, version string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ReadBundlesForRelease(reader Reader, release *releasev1.EksARelease) (*releasev1.Bundles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReleaseForVersion(releases *releasev1.Release, version string) (*releasev1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We treat "latest" as a special case to be able to get the latest build without requiring an exact match.
// We will look for an exact match at the pre-release level and then compare the build metadata.
// This allows a locally built CLI to get the latest dev build without needing to know the exact build number.

// If we are looking for the latest pre-release, we need to compare the build metadata
// Else we continue to look for an exact match.

// If we are looking for the latest pre-release, we need to compare the build metadata
// to find the latest one. CompareBuildMetadata will compare the build identifiers
// in order. For example: v0.19.0-dev+build.10 > v0.19.0-dev+build.9
