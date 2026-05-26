package helm

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/registrymirror"
)

// EnvClientFactory provides a Helm client for a cluster.
type EnvClientFactory struct {
	helmClient Client
	builder    ClientBuilder
}

// NewEnvClientFactory returns a new EnvClientFactory.
func NewEnvClientFactory(builder ClientBuilder) *EnvClientFactory {
	_ = "STUB: not implemented"
	return nil
}

// buildClient returns a new helm executable client.
func (f *EnvClientFactory) buildClient(opts ...Opt) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// Get returns the helm registry client.
// The parameters here are not used and it only returns the client that is initialized using Init.
func (f *EnvClientFactory) Get(_ context.Context, _ *anywherev1.Cluster) (Client, error) {
	_ = "STUB: not implemented"
	return *

	// Init builds the helm registry client once using the registry mirror information from the cluster information.
	// It should be called at least once first, before trying to retrieving and using the client using Get.
	// It only builds the helm registry client once.
	// This is not thread safe and the caller should guarantee that it does not get called from multiple threads.
	new(Client), nil
}

func (f *EnvClientFactory) Init(ctx context.Context, r *registrymirror.RegistryMirror, opts ...Opt) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO (cxbrowne): The registry credentials should be injected on construction through environment variables REGISTRY_USERNAME
// and REGISTRY_PASSWORD, or passed to this method as arguments.
// Issue: https://github.com/aws/eks-anywhere-internal/issues/2115
