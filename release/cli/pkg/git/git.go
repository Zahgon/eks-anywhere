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

package git

func CloneRepo(cloneUrl, destination string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func CheckoutRepo(gitRoot, branch string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func DescribeTag(gitRoot string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetRepoTagsDescending retrieves all Git tags in the specified repository root that match the pattern "v*"
// and returns them as a single string sorted in descending semantic version order (e.g., v3.0.0, v2.1.0, v1.0.0).
func GetRepoTagsDescending(gitRoot string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetHead(gitRoot string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetRepoRoot() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetCurrentBranch(gitRoot string) (string, error) { _ = "STUB: not implemented"; return "", nil }
