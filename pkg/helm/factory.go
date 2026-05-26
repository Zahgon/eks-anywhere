package helm

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// ClientBuilder builds a Helm Client.
type ClientBuilder interface {
	BuildHelm(...Opt) Client
}

// ClientFactory provides a helm client for a cluster.
type ClientFactory struct {
	client  client.Client
	builder ClientBuilder
}

// NewClientForClusterFactory returns a new helm ClientFactory.
func NewClientForClusterFactory(client client.Client, builder ClientBuilder) *ClientFactory {
	_ = "STUB: not implemented"
	return nil
}

// Get returns a new Helm client configured using information from the provided cluster's management cluster.
func (f *ClientFactory) Get(ctx context.Context, clus *anywherev1.Cluster) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}
