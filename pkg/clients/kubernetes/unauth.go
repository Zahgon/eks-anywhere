package kubernetes

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// UnAuthClient is a generic kubernetes API client that takes a kubeconfig
// file on every call in order to authenticate.
type UnAuthClient struct {
	kubectl Kubectl
	scheme  *runtime.Scheme
}

// NewUnAuthClient builds a new UnAuthClient.
func NewUnAuthClient(kubectl Kubectl) *UnAuthClient { _ = "STUB: not implemented"; return nil }

// Init initializes the client internal API scheme
// It has always be invoked at least once before making any API call
// It is not thread safe.
func (c *UnAuthClient) Init() error { _ = "STUB: not implemented"; return nil }

// Get performs a GET call to the kube API server authenticating with a kubeconfig file
// and unmarshalls the response into the provdied Object.
func (c *UnAuthClient) Get(ctx context.Context, name, namespace, kubeconfig string, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// KubeconfigClient returns an equivalent authenticated client.
func (c *UnAuthClient) KubeconfigClient(kubeconfig string) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// BuildClientFromKubeconfig returns an equivalent authenticated client. It will never return
// an error but this helps satisfy a generic factory interface where errors are possible. It's
// basically an alias to KubeconfigClient.
func (c *UnAuthClient) BuildClientFromKubeconfig(kubeconfig string) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// Apply performs an upsert in the form of a client-side apply.
func (c *UnAuthClient) Apply(ctx context.Context, kubeconfig string, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyServerSide creates or patches and object using server side logic.
func (c *UnAuthClient) ApplyServerSide(ctx context.Context, kubeconfig, fieldManager string, obj Object, opts ...ApplyServerSideOption) error {
	_ = "STUB: not implemented"
	return nil
}

// List retrieves list of objects. On a successful call, Items field
// in the list will be populated with the result returned from the server.
func (c *UnAuthClient) List(ctx context.Context, kubeconfig string, list ObjectList, _ ...ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Create saves the object obj in the Kubernetes cluster.
func (c *UnAuthClient) Create(ctx context.Context, kubeconfig string, obj Object) error {
	_ = "STUB: not implemented"
	return nil
}

// Update updates the given obj in the Kubernetes cluster.
func (c *UnAuthClient) Update(ctx context.Context, kubeconfig string, obj Object) error {
	_ = "STUB: not implemented"
	return nil
}

// Patch patches the given object using the given patch and returns the updated object.
// NOTE: This method is not implemented for UnAuthClient and exists only to satisfy the interface.
// Calling this method will return an error indicating it's not supported.
func (c *UnAuthClient) Patch(_ context.Context, _ string, _ Object, _ Patch, _ ...PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes the given obj from Kubernetes cluster.
func (c *UnAuthClient) Delete(ctx context.Context, kubeconfig string, obj Object) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAllOf deletes all objects of the given type matching the given options.
func (c *UnAuthClient) DeleteAllOf(ctx context.Context, kubeconfig string, obj Object, opts ...DeleteAllOfOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *UnAuthClient) resourceTypeForObj(obj runtime.Object) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// if obj is a list, treat it as a request for the "individual" item's resource

func groupVersionToKubectlResourceType(g schema.GroupVersionKind) string {
	_ = "STUB: not implemented"

	// if Group is not set, this probably an obj from "core", which api group is just v1
	return ""
}
