package cluster

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func ValidateConfig(c *Config) error { _ = "STUB: not implemented"; return nil }

type namespaceObject interface {
	runtime.Object
	GetNamespace() string
}

func validateSameNamespace(c *Config, o namespaceObject) error {
	_ = "STUB: not implemented"
	return nil
}
