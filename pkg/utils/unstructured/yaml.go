package unstructured

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func YamlToUnstructured(yamlObjects []byte) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	// Using this CAPI util for now, not sure if we want to depend on it but it's well written
	return nil, nil
}

func UnstructuredToYaml(yamlObjects []unstructured.Unstructured) ([]byte, error) {
	_ = "STUB: not implemented"
	// Using this CAPI util for now, not sure if we want to depend on it but it's well written
	return nil, nil
}

// StripNull removes all null fields from the provided yaml.
func StripNull(resources []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func stripNull(m map[string]interface{}) { _ = "STUB: not implemented"; return }
