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

package tagger

import (
	assettypes "github.com/aws/eks-anywhere/release/cli/pkg/assets/types"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

// BuildToolingGitTagAssigner reads the Git tag from the eks-anywhere-build-tooling repository using the branch name.
// If overrideBranch is provided, it takes precedence over the default branch from the release config.
func BuildToolingGitTagAssigner(rc *releasetypes.ReleaseConfig, gitTagPath, overrideBranch string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CliGitTagAssigner determines the Git tag to use for the CLI repository based on the release configuration.
// If the release is a development release (DevRelease is true), it fetches the list of version tags from the repository
// and returns the most recent (highest) tag in descending semantic version order.
// Otherwise, it uses the explicitly defined ReleaseVersion from the release configuration.
func CliGitTagAssigner(rc *releasetypes.ReleaseConfig, gitTagPath, overrideBranch string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NonExistentTagAssigner is a placeholder GitTagAssigner that always returns the tag "non-existent".
// This can be used in scenarios where Git tagging is irrelevant.
func NonExistentTagAssigner(rc *releasetypes.ReleaseConfig, gitTagPath, overrideBranch string) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// GetGitTagAssigner returns the GitTagAssigner function to be used for the asset configuration.
		// If a custom GitTagAssigner is defined in the AssetConfig, it returns that.
		// Otherwise, it defaults to using the BuildToolingGitTagAssigner.
		nil
}

func GetGitTagAssigner(ac *assettypes.AssetConfig) assettypes.GitTagAssigner {
	_ = "STUB: not implemented"
	return *new(assettypes.GitTagAssigner)
}
