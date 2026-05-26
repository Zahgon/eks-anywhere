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
	"crypto/ecdsa"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	"github.com/golang-jwt/jwt/v5"

	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// ValidateSignature validates the signature annotation of the bundles object using KMS public key.
func ValidateSignature(bundle *anywherev1alpha1.Bundles, pubKey string) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ValidateEKSDistroManifestSignature validates the signature annotation of the bundles object using KMS public key.
func ValidateEKSDistroManifestSignature(release *eksdv1alpha1.Release, signature, pubKey, releaseChannel string) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

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

	// Marshal Bundles object to YAML
	return nil, nil, nil
}

// Convert YAML to JSON for easier gojq processing

// Build and execute the gojq filter that deletes excluded fields

// Compute the SHA256 sum of the filtered JSON

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

// We need to escape '.' for certain gojq path usage
// to avoid ambiguities in the path expressions.

// Parse the final gojq query

// Unmarshal the JSON into a generic interface so gojq can operate

// Run the query

// Marshal the filtered result back to JSON

func parsePublicKey(key string) (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseLicense parses licenseKey jwt token using the public key and returns token fields.
func ParseLicense(licenseToken string, key string) (*jwt.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
