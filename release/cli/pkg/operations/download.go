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

	s3sdk "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/eks-anywhere/release/cli/pkg/retrier"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

func DownloadArtifacts(ctx context.Context, r *releasetypes.ReleaseConfig, eksaArtifacts releasetypes.ArtifactsTable) error {
	_ = "STUB: not implemented"
	// Retrier for downloading source S3 objects. This retrier has a max timeout of 60 minutes. It
	// checks whether the error occured during download is an ObjectNotFound error and retries the
	// download operation for a maximum of 60 retries, with a wait time of 30 seconds per retry.
	return nil
}

// Check if there is an archive to be downloaded

// Check if there is a manifest to be downloaded

func handleArchiveDownload(_ context.Context, r *releasetypes.ReleaseConfig, artifact releasetypes.Artifact, s3Retrier *retrier.Retrier, s3Client *s3sdk.S3, s3Downloader *s3manager.Downloader) error {
	_ = "STUB: not implemented"
	return nil
}

// Download checksum files for the archive

// Adding a special case for tinkerbell/hook project.
// The project builds linux kernel files that are not stored as tarballs and currently do not have SHA checksums.
// TODO(pokearu): Add logic to generate SHA for hook project

func handleManifestDownload(_ context.Context, r *releasetypes.ReleaseConfig, artifact releasetypes.Artifact, s3Retrier *retrier.Retrier, s3Client *s3sdk.S3, s3Downloader *s3manager.Downloader) error {
	_ = "STUB: not implemented"
	return nil
}
