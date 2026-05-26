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

package operations

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"

	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

func UploadArtifacts(ctx context.Context, r *releasetypes.ReleaseConfig, eksaArtifacts releasetypes.ArtifactsTable, isBundleRelease bool) error {
	_ = "STUB: not implemented"
	return nil
}

func handleArchiveUpload(_ context.Context, r *releasetypes.ReleaseConfig, artifact releasetypes.Artifact) error {
	_ = "STUB: not implemented"
	return nil
}

// Adding a special case for tinkerbell/hook project.
// The project builds linux kernel files that are not stored as tarballs and currently do not have SHA checksums.
// TODO(pokearu): Add logic to generate SHA for hook project

func handleManifestUpload(_ context.Context, r *releasetypes.ReleaseConfig, artifact releasetypes.Artifact) error {
	_ = "STUB: not implemented"
	return nil
}

func handleImageUpload(_ context.Context, r *releasetypes.ReleaseConfig, packagesArtifacts map[string][]releasetypes.Artifact, artifact releasetypes.Artifact, defaultSourceEcrAuthConfig, packagesSourceEcrAuthConfig, defaultReleaseEcrAuthConfig, packagesReleaseEcrAuthConfig *docker.AuthConfiguration) error {
	_ = "STUB: not implemented"
	// If the artifact is a helm chart, skip the skopeo copy. Instead, modify the Chart.yaml to match the release tag
	// and then use Helm package and push commands to upload chart to ECR Public
	// Packages Helm chart modification for dev-release is handled elsewhere, so we are checking for that case and skipping
	return nil
}

// Trim -helm on the packages helm chart, but don't need to trim tinkerbell-helm since the AssetName is the same as the repoName
