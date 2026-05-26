package clientutil

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// Kubeclient implements kubernetes.Client interface using a
// client.Client as the underlying implementation.
type KubeClient struct {
	client client.Client
}

func NewKubeClient(client client.Client) *KubeClient { _ = "STUB: not implemented"; return nil }

// Get retrieves an obj for the given name and namespace from the Kubernetes Cluster.
func (c *KubeClient) Get(ctx context.Context, name, namespace string, obj kubernetes.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// List retrieves list of objects. On a successful call, Items field
// in the list will be populated with the result returned from the server.
func (c *KubeClient) List(ctx context.Context, list kubernetes.ObjectList, opts ...kubernetes.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Create saves the object obj in the Kubernetes cluster.
func (c *KubeClient) Create(ctx context.Context, obj kubernetes.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// Update updates the given obj in the Kubernetes cluster.
func (c *KubeClient) Update(ctx context.Context, obj kubernetes.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// Patch updates the given obj in the Kubernetes cluster.
// This method is used only in unit tests currently.
func (c *KubeClient) Patch(ctx context.Context, obj kubernetes.Object, patch kubernetes.Patch, opts ...kubernetes.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyServerSide creates or patches and object using server side logic.
func (c *KubeClient) ApplyServerSide(ctx context.Context, fieldManager string, obj kubernetes.Object, opts ...kubernetes.ApplyServerSideOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes the given obj from Kubernetes cluster.
func (c *KubeClient) Delete(ctx context.Context, obj kubernetes.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAllOf deletes all objects of the given type matching the given options.
func (c *KubeClient) DeleteAllOf(ctx context.Context, obj kubernetes.Object, opts ...kubernetes.DeleteAllOfOption) error {
	_ = "STUB: not implemented"
	return nil
}
