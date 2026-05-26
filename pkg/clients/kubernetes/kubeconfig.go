package kubernetes

import (
	"context"
)

// KubeconfigClient is an authenticated kubernetes API client
// it authenticates using the credentials of a kubeconfig file.
type KubeconfigClient struct {
	client     *UnAuthClient
	kubeconfig string
}

func NewKubeconfigClient(client *UnAuthClient, kubeconfig string) *KubeconfigClient {
	_ = "STUB: not implemented"
	return nil
}

// Get performs a GET call to the kube API server
// and unmarshalls the response into the provided Object.
func (c *KubeconfigClient) Get(ctx context.Context, name, namespace string, obj Object) error {
	_ = "STUB: not implemented"
	return nil
}

// List retrieves list of objects. On a successful call, Items field
// in the list will be populated with the result returned from the server.
func (c *KubeconfigClient) List(ctx context.Context, list ObjectList, _ ...ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Create saves the object obj in the Kubernetes cluster.
func (c *KubeconfigClient) Create(ctx context.Context, obj Object) error {
	_ = "STUB: not implemented"
	return nil
}

// Update updates the given obj in the Kubernetes cluster.
func (c *KubeconfigClient) Update(ctx context.Context, obj Object) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyServerSide creates or patches and object using server side logic.
func (c *KubeconfigClient) ApplyServerSide(ctx context.Context, fieldManager string, obj Object, opts ...ApplyServerSideOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes the given obj from Kubernetes cluster.
func (c *KubeconfigClient) Delete(ctx context.Context, obj Object) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAllOf deletes all objects of the given type matching the given options.
func (c *KubeconfigClient) DeleteAllOf(ctx context.Context, obj Object, opts ...DeleteAllOfOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Patch patches the given obj in the Kubernetes cluster. obj must be a
// struct pointer so that obj can be updated with the content returned by the Server.
func (c *KubeconfigClient) Patch(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}
