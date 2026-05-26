package yaml

import (
	"io"
)

// K8sEncoder leverages the Kubernetes YAML package (sigs.k8s.io/yaml) to provide an Encoder data
// structure that is friendlier to the io package.
type K8sEncoder struct {
	out io.Writer
}

// NewK8sEncoder creates a K8sEncoder instance that writes to out.
func NewK8sEncoder(out io.Writer) K8sEncoder {
	_ = "STUB: not implemented"
	return *

	// Encode marshals v into YAML and writes it to e's output stream.
	new(K8sEncoder)
}

func (e K8sEncoder) Encode(v any) error { _ = "STUB: not implemented"; return nil }
