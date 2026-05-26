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

	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

func RenameArtifacts(ctx context.Context, r *releasetypes.ReleaseConfig, eksaArtifacts releasetypes.ArtifactsTable) error {
	_ = "STUB: not implemented"
	return nil
}

// Change the name of the archive along with the checksum files

// Override images in the manifest with release URIs

func handleArchiveRename(_ context.Context, r *releasetypes.ReleaseConfig, artifact releasetypes.Artifact) error {
	_ = "STUB: not implemented"
	return nil
}

// Change the names of the checksum files

// Adding a special case for tinkerbell/hook project.
// The project builds linux kernel files that are not stored as tarballs and currently do not have SHA checksums.
// TODO(pokearu): Add logic to generate SHA for hook project

func handleManifestRename(_ context.Context, r *releasetypes.ReleaseConfig, artifact releasetypes.Artifact) error {
	_ = "STUB: not implemented"
	return nil
}
