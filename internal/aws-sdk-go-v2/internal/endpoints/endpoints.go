package endpoints

import (
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	defaultProtocol = "https"
	defaultSigner   = "v4"
)

var (
	protocolPriority = []string{"https", "http"}
	signerPriority   = []string{"v4"}
)

// Options provide configuration needed to direct how endpoints are resolved.
type Options struct {
	// Disable usage of HTTPS (TLS / SSL)
	DisableHTTPS bool
}

// Partitions is a slice of partition
type Partitions []Partition

// ResolveEndpoint resolves a service endpoint for the given region and options.
func (ps Partitions) ResolveEndpoint(region string, opts Options) (aws.Endpoint, error) {
	_ = "STUB: not implemented"
	return *new(aws.Endpoint), nil
}

// fallback to first partition format to use when resolving the endpoint.

// Partition is an AWS partition description for a service and its' region endpoints.
type Partition struct {
	ID                string
	RegionRegex       *regexp.Regexp
	PartitionEndpoint string
	IsRegionalized    bool
	Defaults          Endpoint
	Endpoints         Endpoints
}

func (p Partition) canResolveEndpoint(region string) bool { _ = "STUB: not implemented"; return false }

// ResolveEndpoint resolves and service endpoint for the given region and options.
func (p Partition) ResolveEndpoint(region string, options Options) (resolved aws.Endpoint, err error) {
	_ = "STUB: not implemented"
	return *new(aws.Endpoint), nil
}

func (p Partition) endpointForRegion(region string) (Endpoint, bool) {
	_ = "STUB: not implemented"
	return *new(Endpoint), false
}

// Unable to find any matching endpoint, return
// blank that will be used for generic endpoint creation.

// Endpoints is a map of service config regions to endpoints
type Endpoints map[string]Endpoint

// CredentialScope is the credential scope of a region and service
type CredentialScope struct {
	Region  string
	Service string
}

// Endpoint is a service endpoint description
type Endpoint struct {
	// True if the endpoint cannot be resolved for this partition/region/service
	Unresolveable aws.Ternary

	Hostname  string
	Protocols []string

	CredentialScope CredentialScope

	SignatureVersions []string `json:"signatureVersions"`
}

func (e Endpoint) resolve(partition, region string, def Endpoint, options Options) aws.Endpoint {
	_ = "STUB: not implemented"
	return *new(aws.Endpoint)
}

// Only attempt to resolve the endpoint if it can be resolved.

func (e *Endpoint) mergeIn(other Endpoint) { _ = "STUB: not implemented"; return }

func getEndpointScheme(protocols []string, disableHTTPS bool) string {
	_ = "STUB: not implemented"
	return ""
}

func getByPriority(s []string, p []string, def string) string { _ = "STUB: not implemented"; return "" }

// MapFIPSRegion extracts the intrinsic AWS region from one that may have an
// embedded FIPS microformat.
func MapFIPSRegion(region string) string { _ = "STUB: not implemented"; return "" }
