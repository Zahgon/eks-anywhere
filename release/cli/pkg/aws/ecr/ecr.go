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

package ecr

import (
	"github.com/aws/aws-sdk-go/service/ecr"
	docker "github.com/fsouza/go-dockerclient"
)

func GetImageDigest(imageUri, imageContainerRegistry string, ecrClient *ecr.ECR) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetAuthToken(ecrClient *ecr.ECR) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetAuthConfig(ecrClient *ecr.ECR) (*docker.AuthConfiguration, error) {
	_ = "STUB: not implemented"
	// Get ECR authorization token
	return nil, nil
}

// Decode authorization token to get credential pair

// Get password from credential pair

// Construct docker auth configuration

func DescribeImagesPaginated(ecrClient *ecr.ECR, describeInput *ecr.DescribeImagesInput) ([]*ecr.ImageDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FilterECRRepoByTagPrefix will take a substring, and a repository as input and find the latest pushed image matching that substring.
func FilterECRRepoByTagPrefix(ecrClient *ecr.ECR, repoName, prefix, branchName string, isHelmChart bool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Filter out any tags that don't match our prefix for images with multiple tags

// In case we don't find any tag substring matches, we still want to populate the bundle with the latest version.

// imageTagFilter is used when filtering a list of ECR images for a specific tag or tag substring
func imageTagFilter(details []*ecr.ImageDetail, substring string) []*ecr.ImageDetail {
	_ = "STUB: not implemented"
	return nil
}

// getLatestOCIShaTag is used to find the tag/sha of the latest pushed OCI image from a list.
func getLatestOCIShaTag(details []*ecr.ImageDetail, branchName string, isHelmChart bool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Exclude image if any tag contains "release"

// Exclude image if none of the tags contain the branchName

// removeStringSlice removes a named string from a slice, without knowing it's index or it being ordered.
func removeStringSlice(l []*string, item string) []*string { _ = "STUB: not implemented"; return nil }
