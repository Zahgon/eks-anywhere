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

package images

import (
	ecrpublicsdk "github.com/aws/aws-sdk-go/service/ecrpublic"
	docker "github.com/fsouza/go-dockerclient"

	assettypes "github.com/aws/eks-anywhere/release/cli/pkg/assets/types"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

// PollForExistence checks if a Docker image exists in a container registry by polling the registry's manifest endpoint.
// It is primarily used to ensure that a given image (identified by URI and tag) is available before proceeding with operations like image copying.
// For development environments or dev releases, it targets the private registry; otherwise, it targets the public ECR registry.
// If the image is not found (manifest returns "MANIFEST_UNKNOWN"), it retries the request up to 60 times over a 60-minute window (every 30 seconds),
// but only for the "main" branch and if the error is specifically an "requested image not found" error.
func PollForExistence(devRelease bool, authConfig *docker.AuthConfiguration, imageUri, imageContainerRegistry, releaseEnvironment, branchName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Creating new GET request

// Retrier for downloading source ECR images. This retrier has a max timeout of 60 minutes. It
// checks whether the error occured during download is an ImageNotFound error and retries the
// download operation for a maximum of 60 retries, with a wait time of 30 seconds per retry.

func CopyToDestination(sourceAuthConfig, releaseAuthConfig *docker.AuthConfiguration, sourceImageUri, releaseImageUri string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetSourceImageURI(r *releasetypes.ReleaseConfig, name, repoName string, tagOptions map[string]string, imageTagConfiguration assettypes.ImageTagConfiguration, trimVersionSignifier, hasSeparateTagPerReleaseBranch bool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func GetReleaseImageURI(r *releasetypes.ReleaseConfig, name, repoName string, tagOptions map[string]string, imageTagConfiguration assettypes.ImageTagConfiguration, trimVersionSignifier, hasSeparateTagPerReleaseBranch bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// generateFormattedTagPrefix replaces placeholders in the imageTagFormat string with corresponding values from tagOptions.
// Placeholders are denoted by angle brackets (e.g., <env>, <version>).
// Example:
//
//	imageTagFormat: "release-<env>-<version>"
//	tagOptions: map[string]string{"env": "prod", "version": "1.2.3"}
//	returns: "release-prod-1.2.3"
func generateFormattedTagPrefix(imageTagFormat string, tagOptions map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func CompareHashWithPreviousBundle(r *releasetypes.ReleaseConfig, currentSourceImageUri, previousReleaseImageUri string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetPreviousReleaseImageSemver(r *releasetypes.ReleaseConfig, releaseImageUri string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ComputeImageDigestFromManifest(ecrPublicClient *ecrpublicsdk.ECRPublic, registry, repository, tag string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func CheckRepositoryImagesAndTagsCountLimit(sourceImageUri, releaseImageUri, sourceContainerRegistry, releaseContainerRegistry string, ecrClient interface{}, ecrPublicClient *ecrpublicsdk.ECRPublic) error {
	_ = "STUB: not implemented"
	return nil
}
