package hardware

import (
	"io"
	"os"
)

// TinkerbellManifestYAML is a MachineWriter that writes Tinkerbell manifests to a destination.
type TinkerbellManifestYAML struct {
	writer io.Writer
}

// NewTinkerbellManifestYAML creates a TinkerbellManifestYAML instance that writes its manifests to w.
func NewTinkerbellManifestYAML(w io.Writer) *TinkerbellManifestYAML {
	_ = "STUB: not implemented"
	return nil
}

// Write m as a set of Kubernetes manifests for use with Cluster API Tinkerbell Provider. This includes writing a
// Hardware, BMC and Secret (for the BMC).
func (yw *TinkerbellManifestYAML) Write(m Machine) error { _ = "STUB: not implemented"; return nil }

var yamlSeparatorWithNewline = []byte("---\n")

func (yw *TinkerbellManifestYAML) writeWithPrependedSeparator(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (yw *TinkerbellManifestYAML) write(data []byte) error { _ = "STUB: not implemented"; return nil }

// TODO(chrisdoherty4) Patch these types so we can generate yamls again with the new Hardware
// and BaseboardManagement types.

func marshalTinkerbellHardwareYAML(m Machine) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalTinkerbellBMCYAML(m Machine) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalSecretYAML(m Machine) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateOrStdout will create path and return an *os.File if path is not empty. If path is empty
// os.Stdout is returned.
func CreateOrStdout(path string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }
