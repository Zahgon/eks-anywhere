package api

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CleanupPathsFromYaml(yamlContent []byte, paths []string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CleanupPathsInObject unsets or nullifies the provided paths in a give kubernetes Object.
func CleanupPathsInObject[T any, PT interface {
	*T
	client.Object
}](obj PT, paths []string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func deletePaths(m map[string]interface{}, paths []string) { _ = "STUB: not implemented"; return }

func deletePath(m map[string]interface{}, path []string) { _ = "STUB: not implemented"; return }
