package clientutil

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func DeleteYaml(ctx context.Context, c client.Client, yaml []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteObjects(ctx context.Context, c client.Client, objs []client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteObject(ctx context.Context, c client.Client, obj client.Object) error {
	_ = "STUB: not implemented"
	return nil
}
