package serverside

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const fieldManager = "eks-a-controller"

func ReconcileYaml(ctx context.Context, c client.Client, yaml []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func ReconcileObjects(ctx context.Context, c client.Client, objs []client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func ReconcileObject(ctx context.Context, c client.Client, obj client.Object) error {
	_ = "STUB: not implemented"
	// Server side apply
	return nil
}

// UpdateObject updates the existing object during reconciliation.
// This is intended for special use cases only as the preferred method to reconcile objects is server-side apply.
func UpdateObject(ctx context.Context, c client.Client, obj client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteYaml deletes Kubernetes objects from YAML content.
func DeleteYaml(ctx context.Context, c client.Client, yaml []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteObjects deletes multiple Kubernetes objects.
func DeleteObjects(ctx context.Context, c client.Client, objs []client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteObject deletes a single Kubernetes object. It's idempotent - if the object doesn't exist, no error is returned.
func DeleteObject(ctx context.Context, c client.Client, obj client.Object) error {
	_ = "STUB: not implemented"
	return nil
}
