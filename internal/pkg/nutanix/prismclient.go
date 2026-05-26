package nutanix

import (
	"context"

	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
)

// PrismClient is a interface that provides useful functions for performing Prism operations.
type PrismClient interface {
	GetImageUUIDFromName(ctx context.Context, imageName string) (*string, error)
	GetClusterUUIDFromName(ctx context.Context, clusterName string) (*string, error)
	GetSubnetUUIDFromName(ctx context.Context, subnetName string) (*string, error)
}

type client struct {
	v3.Client
}

// NewPrismClient returns an implementation of the PrismClient interface.
func NewPrismClient(endpoint, port string, insecure bool) (PrismClient, error) {
	_ = "STUB: not implemented"
	return *new(PrismClient), nil
}

// GetImageUUIDFromName retrieves the image uuid from the given image name.
func (c *client) GetImageUUIDFromName(ctx context.Context, imageName string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusterUUIDFromName retrieves the cluster uuid from the given cluster name.
//
//nolint:gocyclo
func (c *client) GetClusterUUIDFromName(ctx context.Context, clusterName string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prism Central is also internally a cluster, but we filter that out here as we only care about prism element clusters

// GetSubnetUUIDFromName retrieves the subnet uuid from the given subnet name.
func (c *client) GetSubnetUUIDFromName(ctx context.Context, subnetName string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
