package endpoints

import (
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go/logging"
)

// DefaultKey is a compound map key of a variant and other values.
type DefaultKey struct {
	Variant        EndpointVariant
	ServiceVariant ServiceVariant
}

// EndpointKey is a compound map key of a region and associated variant value.
type EndpointKey struct {
	Region         string
	Variant        EndpointVariant
	ServiceVariant ServiceVariant
}

// EndpointVariant is a bit field to describe the endpoints attributes.
type EndpointVariant uint64

const (
	// FIPSVariant indicates that the endpoint is FIPS capable.
	FIPSVariant EndpointVariant = 1 << (64 - 1 - iota)

	// DualStackVariant indicates that the endpoint is DualStack capable.
	DualStackVariant
)

// ServiceVariant is a bit field to describe the service endpoint attributes.
type ServiceVariant uint64

const (
	defaultProtocol = "https"
	defaultSigner   = "v4"
)

var (
	protocolPriority = []string{"https", "http"}
	signerPriority   = []string{"v4", "s3v4"}
)

// Options provide configuration needed to direct how endpoints are resolved.
type Options struct {
	// Logger is a logging implementation that log events should be sent to.
	Logger logging.Logger

	// LogDeprecated indicates that deprecated endpoints should be logged to the provided logger.
	LogDeprecated bool

	// ResolvedRegion is the resolved region string. If provided (non-zero length) it takes priority
	// over the region name passed to the ResolveEndpoint call.
	ResolvedRegion string

	// Disable usage of HTTPS (TLS / SSL)
	DisableHTTPS bool

	// Instruct the resolver to use a service endpoint that supports dual-stack.
	// If a service does not have a dual-stack endpoint an error will be returned by the resolver.
	UseDualStackEndpoint aws.DualStackEndpointState

	// Instruct the resolver to use a service endpoint that supports FIPS.
	// If a service does not have a FIPS endpoint an error will be returned by the resolver.
	UseFIPSEndpoint aws.FIPSEndpointState

	// ServiceVariant is a bitfield of service specified endpoint variant data.
	ServiceVariant ServiceVariant
}

// GetEndpointVariant returns the EndpointVariant for the variant associated options.
func (o Options) GetEndpointVariant() (v EndpointVariant) {
	_ = "STUB: not implemented"
	return *new(EndpointVariant)
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
	Defaults          map[DefaultKey]Endpoint
	Endpoints         Endpoints
}

func (p Partition) canResolveEndpoint(region string, opts Options) bool {
	_ = "STUB: not implemented"
	return false
}

// ResolveEndpoint resolves and service endpoint for the given region and options.
func (p Partition) ResolveEndpoint(region string, options Options) (resolved aws.Endpoint, err error) {
	_ = "STUB: not implemented"
	return *new(aws.Endpoint), nil
}

func (p Partition) endpointForRegion(region string, variant EndpointVariant, serviceVariant ServiceVariant, endpoints Endpoints) Endpoint {
	_ = "STUB: not implemented"
	return *new(Endpoint)
}

// Unable to find any matching endpoint, return
// blank that will be used for generic endpoint creation.

// Endpoints is a map of service config regions to endpoints
type Endpoints map[EndpointKey]Endpoint

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

	SignatureVersions []string

	// Indicates that this endpoint is deprecated.
	Deprecated aws.Ternary
}

// IsZero returns whether the endpoint structure is an empty (zero) value.
func (e Endpoint) IsZero() bool { _ = "STUB: not implemented"; return false }

func (e Endpoint) resolve(partition, region string, def Endpoint, options Options) (aws.Endpoint, error) {
	_ = "STUB: not implemented"
	return *new(aws.Endpoint), nil
}

// Only attempt to resolve the endpoint if it can be resolved.

func (e *Endpoint) mergeIn(other Endpoint) { _ = "STUB: not implemented"; return }

func getEndpointScheme(protocols []string, disableHTTPS bool) string {
	_ = "STUB: not implemented"
	return ""
}

func getByPriority(s []string, p []string, def string) string { _ = "STUB: not implemented"; return "" }
