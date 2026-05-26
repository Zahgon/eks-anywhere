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

func GetPackagesBundle(r *releasetypes.ReleaseConfig, imageDigests releasetypes.ImageDigestsTable) (anywherev1alpha1.PackageBundle, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.PackageBundle), nil
}

// Find latest Packages dev build for the Helm chart and Image which will always start with `0.0.0` for helm chart and `v0.0.0` for images respectively.
// If we can't find the build starting with our substring, we default to the original dev tag.
// If we do find the Tag in Private ECR, but it doesn't exist in Public ECR Copy the image over so the helm chart will work correctly.

// replaceTag is used to replace the tag of an Image URI with a string.
func replaceTag(uri, tag string) string { _ = "STUB: not implemented"; return "" }
