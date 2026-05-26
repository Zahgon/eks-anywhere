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

package version

import (
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

const FakeComponentChecksum = "abcdef1"

type ProjectVersioner interface {
	patchVersion() (string, error)
}

func BuildComponentVersion(versioner ProjectVersioner, componentCheckSum string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type Versioner struct {
	repoSource    string
	pathToProject string
}

func NewVersioner(pathToProject string) *Versioner { _ = "STUB: not implemented"; return nil }

func (v *Versioner) patchVersion() (string, error) { _ = "STUB: not implemented"; return "", nil }

type VersionerWithGITTAG struct {
	Versioner
	folderWithGITTAG  string
	sourcedFromBranch string
	releaseConfig     *releasetypes.ReleaseConfig
}

func NewVersionerWithGITTAG(repoSource, pathToProject, sourcedFromBranch string, releaseConfig *releasetypes.ReleaseConfig) *VersionerWithGITTAG {
	_ = "STUB: not implemented"
	return nil
}

func NewMultiProjectVersionerWithGITTAG(repoSource, pathToRootFolder, pathToMainProject, sourcedFromBranch string, releaseConfig *releasetypes.ReleaseConfig) *VersionerWithGITTAG {
	_ = "STUB: not implemented"
	return nil
}

func (v *VersionerWithGITTAG) patchVersion() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type cliVersioner struct {
	Versioner
	cliVersion string
}

func NewCliVersioner(cliVersion, pathToProject string) *cliVersioner {
	_ = "STUB: not implemented"
	return nil
}

func (v *cliVersioner) patchVersion() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GenerateComponentHash(hashes []string, dryRun bool) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateManifestHash(r *releasetypes.ReleaseConfig, manifestArtifact *releasetypes.ManifestArtifact) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
