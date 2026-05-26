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

package archives

import (
	assettypes "github.com/aws/eks-anywhere/release/cli/pkg/assets/types"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

func EksDistroArtifactPathGetter(rc *releasetypes.ReleaseConfig, archive *assettypes.Archive, projectPath, gitTag, eksDReleaseChannel, eksDReleaseNumber, kubeVersion, latestPath, arch string) (string, string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil
}

func TarballArtifactPathGetter(rc *releasetypes.ReleaseConfig, archive *assettypes.Archive, projectPath, gitTag, eksDReleaseChannel, eksDReleaseNumber, kubeVersion, latestPath, arch string) (string, string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil
}

func ReleaseBranchedTarballArtifactPathGetter(rc *releasetypes.ReleaseConfig, archive *assettypes.Archive, projectPath, gitTag, eksDReleaseChannel, eksDReleaseNumber, kubeVersion, latestPath, arch string) (string, string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil
}

func KernelArtifactPathGetter(rc *releasetypes.ReleaseConfig, archive *assettypes.Archive, projectPath, gitTag, eksDReleaseChannel, eksDReleaseNumber, kubeVersion, latestPath, arch string) (string, string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil
}

func GetArchiveAssets(rc *releasetypes.ReleaseConfig, archive *assettypes.Archive, projectPath, gitTag, eksDReleaseChannel, eksDReleaseNumber, kubeVersion string) (*releasetypes.ArchiveArtifact, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getArtifactPathGenerator(archive *assettypes.Archive) assettypes.ArchiveS3PathGenerator {
	_ = "STUB: not implemented"
	return *new(assettypes.ArchiveS3PathGenerator)
}
