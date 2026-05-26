package configsources

import (
	"context"
)

// ServiceBaseEndpointProvider is needed to search for all providers
// that provide a configured service endpoint
type ServiceBaseEndpointProvider interface {
	GetServiceBaseEndpoint(ctx context.Context, sdkID string) (string, bool, error)
}

// IgnoreConfiguredEndpointsProvider is needed to search for all providers
// that provide a flag to disable configured endpoints.
//
// Currently duplicated from github.com/aws/aws-sdk-go-v2/config because
// service packages cannot import github.com/aws/aws-sdk-go-v2/config
// due to result import cycle error.
type IgnoreConfiguredEndpointsProvider interface {
	GetIgnoreConfiguredEndpoints(ctx context.Context) (bool, bool, error)
}

// GetIgnoreConfiguredEndpoints is used in knowing when to disable configured
// endpoints feature.
//
// Currently duplicated from github.com/aws/aws-sdk-go-v2/config because
// service packages cannot import github.com/aws/aws-sdk-go-v2/config
// due to result import cycle error.
func GetIgnoreConfiguredEndpoints(ctx context.Context, configs []interface{}) (value bool, found bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// ResolveServiceBaseEndpoint is used to retrieve service endpoints from configured sources
// while allowing for configured endpoints to be disabled
func ResolveServiceBaseEndpoint(ctx context.Context, sdkID string, configs []interface{}) (value string, found bool, err error) {
	_ = "STUB: not implemented"
	return "", false, nil
}
