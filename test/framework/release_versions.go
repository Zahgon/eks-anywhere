package framework

import (
	"github.com/aws/eks-anywhere/pkg/semver"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	prodReleasesManifest = "https://anywhere-assets.eks.amazonaws.com/releases/eks-a/manifest.yaml"
	releaseBinaryName    = "eksctl-anywhere"
	BranchNameEnvVar     = "T_BRANCH_NAME"
	defaultTestBranch    = "main"
)

// GetLatestMinorReleaseFromTestBranch inspects the T_BRANCH_NAME environment variable for a
// branch to retrieve the latest released CLI version. If T_BRANCH_NAME is main, it returns
// the latest minor release.
//
// If T_BRANCH_NAME is not main, it expects it to be of the format release-<major>.<minor>
// and will use the <major>.<minor> to retrieve the previous minor release. For example, if the
// release branch is release-0.2 it will retrieve the latest 0.1 release.
func GetLatestMinorReleaseFromTestBranch() (*releasev1alpha1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPreviousMinorReleaseFromTestBranch inspects the T_BRANCH_NAME environment variable for a
// branch to retrieve the previous latest minor released CLI version. If T_BRANCH_NAME is main, it returns
// the previous minor version which is one minor version below the latest minor version that is released.
//
// If T_BRANCH_NAME is not main, it expects it to be of the format release-<major>.<minor>
// and will use the <major>.<minor> to retrieve the previous minor release. For example, if the
// release branch is release-0.2 it will retrieve the latest 0.1 release.
func GetPreviousMinorReleaseFromTestBranch() (*releasev1alpha1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For release branch just return the latest minor release

// EKSAVersionForTestBinary returns the "future" EKS-A version for the tested binary based on the TEST_BRANCH name.
// For main, it returns the next minor version.
// For a release branch, it returns the next path version for that release minor version.
func EKSAVersionForTestBinary() (string, error) { _ = "STUB: not implemented"; return "", nil }

func eksaVersionForMain() (string, error) { _ = "STUB: not implemented"; return "", nil }

func eksaVersionForReleaseBranch(branch string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// if no patch version for the release branch, this is an unreleased minor version
// so the next version will be x.x.0

func GetLatestMinorReleaseBinaryFromMain() (binaryPath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetLatestMinorReleaseFromMain() (*releasev1alpha1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func semverForReleaseBranch(branch string) (*semver.Version, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func latestRelease(releases *releasev1alpha1.Release) (*releasev1alpha1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// localEksaCLIDevVersionRelease returns the EKS-A release for the local eks-a CLI version.
// It reads the version from the local eks-a CLI by running `eksctl anywhere version` command
// and follows the same logic as the CLI to extract the release.
func localEksaCLIDevVersionRelease() (*releasev1alpha1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPreviousMinorReleaseFromVersion calculates the previous minor release by decrementing the
// version minor number, then retrieves the latest <major>.<minor>.<patch>  for the calculated
// version.
func GetPreviousMinorReleaseFromVersion(version *semver.Version) (*releasev1alpha1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetReleaseBinaryFromVersion(version *semver.Version) (binaryPath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewEKSAReleasePackagedBinary builds a new EKSAReleasePackagedBinary.
func NewEKSAReleasePackagedBinary(release *releasev1alpha1.EksARelease) *EKSAReleasePackagedBinary {
	_ = "STUB: not implemented"
	return nil
}

// EKSAReleasePackagedBinary decorates an EKSA release with extra functionality.
type EKSAReleasePackagedBinary struct {
	*releasev1alpha1.EksARelease
}

// BinaryPath implements EKSAPackagedBinary.
func (b *EKSAReleasePackagedBinary) BinaryPath() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Version returns the eks-a release version.
func (b *EKSAReleasePackagedBinary) Version() string { _ = "STUB: not implemented"; return "" }

func getBinary(release *releasev1alpha1.EksARelease) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getBinaryFromRelease(release *releasev1alpha1.EksARelease, chainedErr error) (binaryPath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

type platformAwareRelease struct {
	*releasev1alpha1.EksARelease
}

func (p *platformAwareRelease) binaryUri() (binaryUri string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func prodReleases() (release *releasev1alpha1.Release, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getReleases(url string) (release *releasev1alpha1.Release, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getLatestPrevMinorRelease(releases *releasev1alpha1.Release, releaseBranchVersion *semver.Version) (*releasev1alpha1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLatestPatchRelease return the latest patch version for the major.minor release version.
// If releases doesn't contain a major.minor for version, it returns nil.
func GetLatestPatchRelease(releases *releasev1alpha1.Release, version *semver.Version) *releasev1alpha1.EksARelease {
	_ = "STUB: not implemented"
	return nil
}

// GetLatestProductionPatchRelease retrieves the latest patch release for version from the
// production release manifest. If the production release manifest does not contain a release for
// the major.minor of version it errors.
func GetLatestProductionPatchRelease(version *semver.Version) (*releasev1alpha1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMajorMinorFromTestBranch(testBranch string) string { _ = "STUB: not implemented"; return "" }

func devReleaseURL() string { _ = "STUB: not implemented"; return "" }

func testBranch() string { _ = "STUB: not implemented"; return "" }
