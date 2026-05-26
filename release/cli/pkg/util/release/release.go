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

package release

import (
	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

type EksAReleases []anywherev1alpha1.EksARelease

func GetPreviousReleaseIfExists(r *releasetypes.ReleaseConfig) (*anywherev1alpha1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppendOrUpdateRelease appends a new release to the manifest if it does not exist, or updates the existing release.
func (releases EksAReleases) AppendOrUpdateRelease(r anywherev1alpha1.EksARelease) EksAReleases {
	_ = "STUB: not implemented"
	return *new(EksAReleases)
}

// Trim removes the oldest releases if the manifest size exceeds the maxSize.
// If maxSize is -1, no releases are removed.
func Trim(releases EksAReleases, maxSize int) EksAReleases {
	_ = "STUB: not implemented"
	return *new(EksAReleases)
}
