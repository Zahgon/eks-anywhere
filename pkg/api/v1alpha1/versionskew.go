package v1alpha1

import (
	"k8s.io/apimachinery/pkg/util/version"
)

// SupportedMinorVersionIncrement represents the minor version skew for kubernetes version upgrades.
const SupportedMinorVersionIncrement = 1

// ValidateVersionSkew validates Kubernetes version skew between valid non-nil versions.
func ValidateVersionSkew(oldVersion, newVersion *version.Version) error {
	_ = "STUB: not implemented"
	return nil
}
