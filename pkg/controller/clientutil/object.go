package clientutil

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ObjectsToClientObjects[T client.Object](objs []T) []client.Object {
	_ = "STUB: not implemented"
	return nil
}
