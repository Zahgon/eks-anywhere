package clientutil

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func YamlToClientObjects(yamlObjects []byte) ([]client.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use a numbered loop to avoid problems when retrieving the pointer
