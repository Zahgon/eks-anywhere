// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filereader

import (
	"sync"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

// Global mutex for synchronizing git operations to prevent concurrent git lock issues
var gitOperationMutex sync.Mutex

type EksDLatestRelease struct {
	Branch               string `json:"branch"`
	KubeVersion          string `json:"kubeVersion"`
	Number               int    `json:"number"`
	Dev                  bool   `json:"dev,omitempty"`
	EndOfStandardSupport string `json:"endOfStandardSupport,omitempty"`
}

type EksDLatestReleases struct {
	Releases []EksDLatestRelease `json:"releases"`
	Latest   string              `json:"latest"`
}

func ReadShaSums(filename string, r *releasetypes.ReleaseConfig) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func readShaFile(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ReadFileContentsTrimmed(filename string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ReadEksDReleases(r *releasetypes.ReleaseConfig) (*EksDLatestReleases, error) {
	_ = "STUB: not implemented"
	// Read the eks-d latest release file to get all the releases
	return nil, nil
}

func GetSupportedK8sVersions(r *releasetypes.ReleaseConfig) ([]string, error) {
	_ = "STUB: not implemented"
	// Read the eks-d latest release file to get all the releases
	return nil, nil
}

func GetBottlerocketSupportedK8sVersionsByFormat(r *releasetypes.ReleaseConfig, imageFormat string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the eks-d latest release file to get all the releases

// new format for BR releases file

func GetBottlerocketContainerMetadata(r *releasetypes.ReleaseConfig, filename string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func GetEksDReleaseManifestUrl(releaseChannel, releaseNumber string, dev bool) string {
	_ = "STUB: not implemented"
	return ""
}

// GetNextEksADevBuildNumber computes next eksa dev build number for the current eks-a dev build
func GetNextEksADevBuildNumber(releaseVersion string, r *releasetypes.ReleaseConfig) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Check if current version and latest version are the same semver

// TODO: remove when we update the pipeline

// NewBuildNumberFromLastVersion bumps the build number for eksa dev build version if found
func NewBuildNumberFromLastVersion(latestEksaBuildVersion, releaseVersion, branchName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
	// TODO: remove when we update the pipeline
}

// TODO: adding vDev case temporally to support old run, remove later
// different semver, reset build number suffix on release version

// Same semver, only bump build number suffix on release version

func GetCurrentEksADevReleaseVersion(releaseVersion string, r *releasetypes.ReleaseConfig, buildNumber int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TODO: remove when we update the pipeline

func PutEksAReleaseVersion(version string, r *releasetypes.ReleaseConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Upload the file to S3

func GetEksdRelease(eksdReleaseURL string) (*eksdv1alpha1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve values from https://github.com/aws/eks-anywhere-build-tooling/blob/main/EKSD_LATEST_RELEASES
func GetEksdReleaseValues(release interface{}) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func ReadHttpFile(uri string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ReadGitTag(projectPath, gitRootPath, branch string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Clean up any potential lock files before git operations

// cleanupGitLockFiles removes git lock files that might prevent concurrent operations
func cleanupGitLockFiles(gitRootPath string) { _ = "STUB: not implemented"; return }
