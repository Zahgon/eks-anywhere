package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	snowEC2Port        = 8243
	snowballDevicePort = 9092
)

func BuildClients(ctx context.Context) (Clients, error) {
	_ = "STUB: not implemented"
	return *new(Clients), nil
}

func snowEndpoints(deviceIP string) []ServiceEndpoint { _ = "STUB: not implemented"; return nil }

// WithCustomCABundleFile is a helper function to construct functional options
// that reads an aws certificates file and sets CustomCABundle on config's LoadOptions.
func WithCustomCABundleFile(certsFile string) AwsConfigOpt {
	_ = "STUB: not implemented"
	return *new(AwsConfigOpt)
}

// WithSnowEndpointAccess gathers all the config's LoadOptions for snow,
// which includes snowball ec2 endpoint, snow credentials for a specific profile,
// and CA bundles for accessing the https endpoint.
func WithSnowEndpointAccess(deviceIP string, certsFile, credsFile string) AwsConfigOpt {
	_ = "STUB: not implemented"
	return *new(AwsConfigOpt)
}

func SnowEndpointResolver(deviceIP string) aws.EndpointResolverWithOptionsFunc {
	_ = "STUB: not implemented"
	return *new(aws.EndpointResolverWithOptionsFunc)
}

// returning EndpointNotFoundError allows the service to fallback to it's default resolution
