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

	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

func GenerateBundleArtifactsTable(r *releasetypes.ReleaseConfig) (releasetypes.ArtifactsTable, error) {
	_ = "STUB: not implemented"
	return *new(releasetypes.ArtifactsTable), nil
}

func BundleArtifactsRelease(r *releasetypes.ReleaseConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func GenerateImageDigestsTable(ctx context.Context, r *releasetypes.ReleaseConfig) (releasetypes.ImageDigestsTable, error) {
	_ = "STUB: not implemented"
	return *new(releasetypes.ImageDigestsTable), nil
}

func SignImagesNotation(r *releasetypes.ReleaseConfig, imageDigests releasetypes.ImageDigestsTable) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip signing image if it is already signed.

// Sign public ECR image using AWS signer and notation CLI
// notation sign <registry>/<repository>@<sha256:shasum> --plugin com.amazonaws.signer.notation.plugin --id <signer_profile_arn>

// Copy image signatures to production account from staging account
func CopyImageSignatureUsingOras(r *releasetypes.ReleaseConfig, imageDigests releasetypes.ImageDigestsTable) error {
	_ = "STUB: not implemented"
	return nil
}

// Get imageRespository name since we have a different source and release registry.

// Compute image digest for source and destination images from manifest contents. The digest will be in the
// form sha256:digest, so we are changing it to Notation's image index and signatures format, sha256-digest.

// Form releaseImageURI in the form <source-registry>/<repository>:<sha256-digest>

// Form releaseImageURI in the form <release-registry>/<repository>:<sha256-digest>

func GenerateBundleSpec(r *releasetypes.ReleaseConfig, bundle *anywherev1alpha1.Bundles, imageDigests releasetypes.ImageDigestsTable) error {
	_ = "STUB: not implemented"
	return nil
}

// SignBundleManifest is the top-level function that computes the Bundles
// manifest signature using AWS KMS and attaches that signature as an
// annotation on the Bundles object.
func SignBundleManifest(ctx context.Context, bundle *anywherev1alpha1.Bundles) error {
	_ = "STUB: not implemented"
	return nil
}

// SignEKSDistroManifest is the top-level function that computes the EKS Distro
// manifest signature using AWS KMS and attaches that signature as an
// annotation on the Bundles object for each supported kubernetes version.
func SignEKSDistroManifest(ctx context.Context, bundle *anywherev1alpha1.Bundles) error {
	_ = "STUB: not implemented"
	return nil
}

func getImageDigest(_ context.Context, r *releasetypes.ReleaseConfig, artifact releasetypes.Artifact) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
