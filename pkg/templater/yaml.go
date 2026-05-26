package templater

import (
	"k8s.io/apimachinery/pkg/runtime"
)

const objectSeparator string = "\n---\n"

func AppendYamlResources(resources ...[]byte) []byte { _ = "STUB: not implemented"; return nil }

func ObjectsToYaml(objs ...runtime.Object) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
