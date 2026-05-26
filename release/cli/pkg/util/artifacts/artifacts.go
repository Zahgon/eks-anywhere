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

package artifacts

func IsObjectNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

func IsImageNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

func GetFakeSHA(hashType int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetLatestUploadDestination(sourcedFromBranch string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetURI returns an full URL for the given path.
func GetURI(cdn, path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func SplitImageUri(imageUri, imageContainerRegistry string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}
