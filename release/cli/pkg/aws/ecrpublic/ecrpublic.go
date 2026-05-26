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

package ecrpublic

import (
	"github.com/aws/aws-sdk-go/service/ecrpublic"
	docker "github.com/fsouza/go-dockerclient"
)

func GetImageDigest(imageUri, imageContainerRegistry string, ecrPublicClient *ecrpublic.ECRPublic) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetAuthToken(ecrPublicClient *ecrpublic.ECRPublic) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetAuthConfig(ecrPublicClient *ecrpublic.ECRPublic) (*docker.AuthConfiguration, error) {
	_ = "STUB: not implemented"
	// Get ECR Public authorization token
	return nil, nil
}

// Decode authorization token to get credential pair

// Get password from credential pair

// Construct docker auth configuration

func GetAllImagesCount(imageRepository string, ecrPublicClient *ecrpublic.ECRPublic) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetTagsCountForImage(imageRepository, imageDigest string, ecrPublicClient *ecrpublic.ECRPublic) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func CheckImageExistence(imageUri, imageContainerRegistry string, ecrPublicClient *ecrpublic.ECRPublic) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
