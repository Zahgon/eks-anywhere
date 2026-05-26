package validations

import (
	"context"
	"time"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	"github.com/golang-jwt/jwt/v5"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// LicensePublicKey is the public key for verifying license token signature.
// this is injected at build time.
var LicensePublicKey string

// ValidateExtendedK8sVersionSupport validates all the validations needed for the support of extended kubernetes support.
func ValidateExtendedK8sVersionSupport(ctx context.Context, clusterSpec anywherev1.Cluster, bundle *v1alpha1.Bundles, releaseManifest *eksdv1alpha1.Release, k kubernetes.Client) error {
	_ = "STUB: not implemented"
	// Validate EKS-A bundle has not been modified by verifying the signature in the bundle annotation
	return nil
}

// Check whether the kubernetes version for the cluster is currently under extended support by comparing the endOfStandardSupport date from the bundle with the current date.

// Validate EKS Distro manifest has not been modified by verifying the signature in the EKS-A bundle annotation

// Validate that the claims in the license token have not been modified by verifying the signature in the token

// Validate that the license token has not expired yet

// Validate that the same license token is not being used by multiple clusters

// validateBundleSignature validates bundles signature with the KMS public key.
func validateBundleSignature(bundle *v1alpha1.Bundles) error { _ = "STUB: not implemented"; return nil }

// validateEKSDistroManifestSignature validates eks distro manifest signature with the KMS public key.
func validateEKSDistroManifestSignature(eksdReleaseManifest *eksdv1alpha1.Release, sig, releaseChannel string) error {
	_ = "STUB: not implemented"
	return nil
}

func isExtendedSupport(versionsBundle *v1alpha1.VersionsBundle) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getLicense(licenseToken string) (*jwt.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateLicense(token *jwt.Token) error { _ = "STUB: not implemented"; return nil }

func isPastDateThanToday(dateToCompare time.Time) bool { _ = "STUB: not implemented"; return false }

func validateLicenseKeyIsUnique(ctx context.Context, clusterName string, licenseToken string, k kubernetes.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// ShouldSkipBundleSignatureValidation returns true if the eksa version is less than v0.22.0
// and false otherwise. This is to skip signature validation for older eksa versions.
func ShouldSkipBundleSignatureValidation(eksaVersion *string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// v0.22.0 is the first version to support extended kubernetes support
