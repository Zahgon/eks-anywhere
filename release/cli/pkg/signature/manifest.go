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

package signature

import (
	"context"
	"strings"
	"text/template"

	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
)

// GojqTemplate is used to build a gojq filter expression that deletes the desired fields.
var GojqTemplate = template.Must(template.New("gojq_query").Funcs(
	template.FuncMap{
		"StringsJoin": strings.Join,
		"Escape": func(in string) string {
			// We need to escape '.' for certain gojq path usage
			// to avoid ambiguities in the path expressions.
			return strings.ReplaceAll(in, ".", "\\\\.")
		},
	},
).Parse(`
del({{ StringsJoin .Excludes ", " }})
`))

// GetBundleSignature calls KMS and retrieves a signature, then base64-encodes it
// to store in the Bundles manifest annotation.
func GetBundleSignature(ctx context.Context, bundle *anywherev1alpha1.Bundles, key string) (string, error) {
	_ = "STUB: not implemented"
	// Compute the digest from the Bundles manifest, excluding certain fields.
	return "", nil
}

// Create KMS Client for bundle manifest signing

// The KMS Sign API requires the raw hash as the Message when MessageType is DIGEST.

// Return the base64-encoded signature.

// GetEKSDistroManifestSignature calls KMS and retrieves a signature, then base64-encodes it
// to store in the Bundles manifest annotation.
func GetEKSDistroManifestSignature(ctx context.Context, bundle *anywherev1alpha1.Bundles, key, releaseUrl string) (string, error) {
	_ = "STUB: not implemented"
	// Retrieve the eks-distro release from the release URL.
	return "", nil
}

// Compute the digest for the eks-distro release, excluding certain fields.

// Create KMS Client for eks distro manifest signing

// The KMS Sign API requires the raw hash as the Message when MessageType is DIGEST.

// Return the base64-encoded signature.

// getEksdDigest computes the SHA256 digest for an EKS Distro release object.
// It follows similar steps as getBundleDigest() for Bundles by marshalling the object,
// converting it to JSON, filtering out undesired fields, and then computing the hash.
func getEKSDistroReleaseDigest(release *eksdv1alpha1.Release) ([32]byte, []byte, error) {
	_ = "STUB: not implemented"

	// Marshal the eks-distro release object to YAML.
	return nil, nil, nil
}

// Convert the YAML to JSON for easier gojq processing.

// Build and execute the gojq filter that deletes excluded fields.

// Compute the SHA256 digest of the filtered JSON.

// getBundleDigest converts the Bundles manifest to JSON, excludes certain fields, then
// computes the SHA256 hash of the filtered manifest. It returns the digest and
// the final bytes used to produce that digest.
func getBundleDigest(bundle *anywherev1alpha1.Bundles) ([32]byte, []byte, error) {
	_ = "STUB: not implemented"

	// Marshal Bundles object to YAML.
	return nil, nil, nil
}

// Convert YAML to JSON for easier gojq processing.

// Build and execute the gojq filter that deletes excluded fields.

// Compute the SHA256 digest of the filtered JSON.

// filterExcludes applies the default and user-specified excludes to the JSON
// representation of the Bundles object using gojq.
func filterExcludes(jsonBytes []byte, excludes string) ([]byte, error) {
	_ = "STUB: not implemented"
	// Decode the base64-encoded excludes
	return nil, nil
}

// Convert them into slice of strings

// Combine AlwaysExcluded with userExcludes

// Build the argument to the gojq template

// Parse the final gojq query

// Unmarshal the JSON into a generic interface so gojq can operate

// Run the query

// Marshal the filtered result back to JSON
