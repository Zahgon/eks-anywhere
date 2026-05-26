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

package types

import (
	"sync"
	"time"

	"github.com/aws/eks-anywhere/release/cli/pkg/clients"
)

// ReleaseConfig contains metadata fields for a release.
type ReleaseConfig struct {
	ReleaseVersion                   string
	DevReleaseUriVersion             string
	BundleNumber                     int
	CliMinVersion                    string
	CliMaxVersion                    string
	CliRepoUrl                       string
	CliRepoSource                    string
	CliRepoHead                      string
	CliRepoBranchName                string
	BuildRepoUrl                     string
	BuildRepoSource                  string
	BuildRepoHead                    string
	BuildRepoBranchName              string
	ArtifactDir                      string
	SourceBucket                     string
	ReleaseBucket                    string
	SourceContainerRegistry          string
	PackagesSourceContainerRegistry  string
	ReleaseContainerRegistry         string
	PackagesReleaseContainerRegistry string
	CDN                              string
	ReleaseNumber                    int
	ReleaseDate                      string
	ReleaseTime                      time.Time
	DevRelease                       bool
	BundleRelease                    bool
	DryRun                           bool
	Weekly                           bool
	ReleaseEnvironment               string
	SourceClients                    *clients.SourceClients
	ReleaseClients                   *clients.ReleaseClients
	BundleArtifactsTable             ArtifactsTable
	EksAArtifactsTable               ArtifactsTable
	AwsSignerProfileArn              string
	MaxReleasesInManifest            int
}

type ImageTagOverride struct {
	Repository string
	ReleaseUri string
}

type ArchiveArtifact struct {
	SourceS3Key       string
	SourceS3Prefix    string
	ArtifactPath      string
	ReleaseName       string
	ReleaseS3Path     string
	ReleaseCdnURI     string
	OS                string
	OSName            string
	Arch              []string
	GitTag            string
	ProjectPath       string
	SourcedFromBranch string
	ImageFormat       string
	Private           bool
}

type ImageArtifact struct {
	AssetName         string
	SourceImageURI    string
	ReleaseImageURI   string
	OS                string
	Arch              []string
	GitTag            string
	ProjectPath       string
	SourcedFromBranch string
}

type ManifestArtifact struct {
	SourceS3Prefix    string // S3 uri till base to download artifact
	SourceS3Key       string
	ArtifactPath      string
	ReleaseName       string
	ReleaseS3Path     string
	ReleaseCdnURI     string
	ImageTagOverrides []ImageTagOverride
	GitTag            string
	ProjectPath       string
	SourcedFromBranch string
	Component         string
	Private           bool
}

type Artifact struct {
	Archive  *ArchiveArtifact
	Image    *ImageArtifact
	Manifest *ManifestArtifact
}

type ArtifactsTable struct {
	artifactsMap sync.Map
}

type ImageDigestsTable struct {
	imageDigestsMap sync.Map
}

// BundlesManifestFilepath returns the filepath for the bundles manifest.
func (r *ReleaseConfig) BundlesManifestFilepath() string { _ = "STUB: not implemented"; return "" }

// ReleaseManifestFilepath returns the filepath for the release manifest.
func (r *ReleaseConfig) ReleaseManifestFilepath() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactsTable) Load(projectName string) ([]Artifact, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ArtifactsTable) Store(projectName string, artifacts []Artifact) {
	_ = "STUB: not implemented"
	return
}

func (a *ArtifactsTable) Range(f func(key, value any) bool) { _ = "STUB: not implemented"; return }

func (i *ImageDigestsTable) Load(imageURI string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (i *ImageDigestsTable) Store(imageURI, imageDigest string) { _ = "STUB: not implemented"; return }

func (i *ImageDigestsTable) Range(f func(key, value any) bool) { _ = "STUB: not implemented"; return }
