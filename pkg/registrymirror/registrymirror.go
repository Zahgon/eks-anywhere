package registrymirror

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// RegistryMirror configures mirror mappings for artifact registries.
type RegistryMirror struct {
	// BaseRegistry is the address of the registry mirror without namespace. Just the host and the port.
	BaseRegistry string
	// NamespacedRegistryMap stores mirror mappings for artifact registries
	NamespacedRegistryMap map[string]string
	// Auth should be marked as true if authentication is required for the registry mirror
	Auth bool
	// CACertContent defines the contents registry mirror CA certificate
	CACertContent string
	// InsecureSkipVerify skips the registry certificate verification.
	// Only use this solution for isolated testing or in a tightly controlled, air-gapped environment.
	InsecureSkipVerify bool
}

// FromCluster is a constructor for RegistryMirror from a cluster schema.
func FromCluster(cluster *v1alpha1.Cluster) *RegistryMirror { _ = "STUB: not implemented"; return nil }

// FromClusterRegistryMirrorConfiguration is a constructor for RegistryMirror from a RegistryMirrorConfiguration schema.
func FromClusterRegistryMirrorConfiguration(config *v1alpha1.RegistryMirrorConfiguration) *RegistryMirror {
	_ = "STUB: not implemented"
	return nil
}

// add registry mirror base address
// for each namespace, add corresponding endpoint

// for backward compatibility, default mapping for public.ecr.aws is added
// when no namespace mapping is specified

// CoreEKSAMirror returns the configured mirror for public.ecr.aws.
func (r *RegistryMirror) CoreEKSAMirror() string { _ = "STUB: not implemented"; return "" }

// ReplaceRegistry replaces the host in a url with corresponding registry mirror
// It supports full URLs and container image URLs
// If the provided original url is malformed, there are no guarantees
// that the returned value will be valid
// If no corresponding registry mirror, it will return the original URL.
func (r *RegistryMirror) ReplaceRegistry(url string) string { _ = "STUB: not implemented"; return "" }
