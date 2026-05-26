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

package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	ecrsdk "github.com/aws/aws-sdk-go/service/ecr"
	ecrpublicsdk "github.com/aws/aws-sdk-go/service/ecrpublic"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	docker "github.com/fsouza/go-dockerclient"
)

type SourceClients struct {
	S3       *SourceS3Clients
	ECR      *SourceECRClient
	Packages *SourceECRClient
}

type ReleaseClients struct {
	S3        *ReleaseS3Clients
	ECRPublic *ReleaseECRPublicClient
	Packages  *ReleaseECRPublicClient
}

type SourceS3Clients struct {
	Client     *s3.S3
	Downloader *s3manager.Downloader
}

type ReleaseS3Clients struct {
	Client   *s3.S3
	Uploader *s3manager.Uploader
}

type SourceECRClient struct {
	EcrClient       *ecrsdk.ECR
	EcrPublicClient *ecrpublicsdk.ECRPublic
	AuthConfig      *docker.AuthConfiguration
}

type ReleaseECRPublicClient struct {
	Client     *ecrpublicsdk.ECRPublic
	AuthConfig *docker.AuthConfiguration
}

// Function to create release clients for dev release.
func CreateDevReleaseClients(dryRun bool) (*SourceClients, *ReleaseClients, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PDX session for eks-a-build-prod-pdx

// IAD session for eks-a-build-prod-pdx

// PDX Session for beta-pdx-packages

// IAD Session for beta-pdx-packages

// S3 client and uploader

// Get source ECR auth config

// Get packages source ECR auth config

// Get release ECR Public auth config

// Get packages release ECR Public auth config

// Constructing source clients

// Constructing release clients

// Function to create clients for staging release.
func CreateStagingReleaseClients(bundleRelease bool) (*SourceClients, *ReleaseClients, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Session for eks-a-build-prod-pdx

// Session for eks-a-artifact-beta-iad

// Source S3 client

// Release S3 client and uploader

// Get source ECR auth config

// Get release ECR Public auth config

// Session for beta-pdx-packages

// Get packages source ECR auth config

// Constructing source clients

// Constructing release clients

// Function to create clients for production release.
func CreateProdReleaseClients() (*SourceClients, *ReleaseClients, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Session for eks-a-artifact-beta-iad

// Session for eks-a-artifact-prod-iad

// Source S3 client

// Release S3 client and uploader

// Get source ECR Public auth config

// Get release ECR Public auth config

// Constructing release clients

// Constructing release clients

// Function to create KMS client for bundle manifest signing and eks distro manifest signing.
func CreateKMSClient(ctx context.Context) (*kms.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
