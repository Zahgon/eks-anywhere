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

package bundles

import (
	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

func GetBottlerocketHostContainersBundle(r *releasetypes.ReleaseConfig, eksDReleaseChannel string, imageDigests releasetypes.ImageDigestsTable) (anywherev1alpha1.BottlerocketHostContainersBundle, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.BottlerocketHostContainersBundle), nil
}

func bottlerocketDefaultArtifact(r *releasetypes.ReleaseConfig, metadataFile, imageName string) (anywherev1alpha1.Image, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.Image), nil
}

// getBottlerocketBootstrapArtifact is a shared helper function that retrieves a specific
// bottlerocket bootstrap artifact by asset name from the bundle artifacts table.
func getBottlerocketBootstrapArtifact(r *releasetypes.ReleaseConfig, eksDReleaseChannel string, imageDigests releasetypes.ImageDigestsTable, assetName string) (anywherev1alpha1.Image, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.Image), nil
}

func bottlerocketKubeadmBootstrapArtifact(r *releasetypes.ReleaseConfig, eksDReleaseChannel string, imageDigests releasetypes.ImageDigestsTable) (anywherev1alpha1.Image, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.Image), nil
}

func GetBottlerocketBootstrapContainersBundle(r *releasetypes.ReleaseConfig, eksDReleaseChannel string, imageDigests releasetypes.ImageDigestsTable) (anywherev1alpha1.BottlerocketBootstrapContainersBundle, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.BottlerocketBootstrapContainersBundle), nil
}

// VSphere multi-network bootstrap container (optional)

// Note: We don't return an error if the artifact is not found since bootstrap containers are optional
// and may not be available for all release channels or configurations.

func bottlerocketMultiNetworkArtifact(r *releasetypes.ReleaseConfig, eksDReleaseChannel string, imageDigests releasetypes.ImageDigestsTable) (anywherev1alpha1.Image, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.Image), nil
}
