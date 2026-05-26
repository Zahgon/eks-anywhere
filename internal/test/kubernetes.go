package test

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	_ "github.com/aws/eks-anywhere/internal/test/envtest"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// testKubeClient implements a kubernetes.Client that uses
// a fake client.Client under the hood. It reimplements server-side
// apply since Fake client doesn't support it.
type testKubeClient struct {
	kubernetes.Client
	fakeClient client.Client
}

// ApplyServerSide creates or patches an object using server side apply logic.
// Uses client.Apply patch type like production code. Requires GVK to be set on object.
func (t *testKubeClient) ApplyServerSide(ctx context.Context, fieldManager string, obj kubernetes.Object, opts ...kubernetes.ApplyServerSideOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure GVK is set - client.Apply requires it

// Use client.Apply with proper options

// NewKubeClient builds a new kubernetes.Client by using client.Client.
func NewKubeClient(client client.Client) kubernetes.Client {
	_ = "STUB: not implemented"
	return *new(kubernetes.Client)
}

// NewFakeKubeClient returns a kubernetes.Client that uses a fake client.Client under the hood.
func NewFakeKubeClient(objs ...client.Object) kubernetes.Client {
	_ = "STUB: not implemented"
	return *new(kubernetes.Client)
}

// NewFakeKubeClientAlwaysError returns a kubernetes.Client  that will always fail in any operation
// This is achieved by injecting an empty Scheme, which will make the underlying client.Client
// incapable of determining the resource type for a particular client.Object.
func NewFakeKubeClientAlwaysError(objs ...client.Object) kubernetes.Client {
	_ = "STUB: not implemented"
	return *new(kubernetes.Client)
}
