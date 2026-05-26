package hardware

import (
	"io"

	tinkv1alpha1 "github.com/tinkerbell/tink/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"

	rufiov1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/rufio"
)

// Indexer provides indexing behavior for objects.
type Indexer interface {
	// Lookup retrieves objects associated with the index => value pair.
	Lookup(index, value string) ([]interface{}, error)
	// Insert inserts v int the index.
	Insert(v interface{}) error
	// IndexField associated index with fn such that Lookup may be used to retrieve objects.
	IndexField(index string, fn KeyExtractorFunc)
	// Remove deletes v from the index.
	Remove(v interface{}) error
}

// Catalogue represents a catalogue of Tinkerbell hardware manifests to be used with Tinkerbells
// Kubefied back-end.
type Catalogue struct {
	hardware      []*tinkv1alpha1.Hardware
	hardwareIndex Indexer

	bmcs     []*rufiov1alpha1.Machine
	bmcIndex Indexer

	secrets     []*corev1.Secret
	secretIndex Indexer
}

// CatalogueOption defines an option to be applied in Catalogue instantiation.
type CatalogueOption func(*Catalogue)

// NewCatalogue creates a new Catalogue instance.
func NewCatalogue(opts ...CatalogueOption) *Catalogue { _ = "STUB: not implemented"; return nil }

// ParseYAMLCatalogueFromFile parses filename, a YAML document, using ParseYamlCatalogue.
func ParseYAMLCatalogueFromFile(catalogue *Catalogue, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// ParseYAMLCatalogue parses a YAML document, r, that represents a set of Kubernetes manifests.
// Manifests parsed include CAPT Hardware, PBnJ BMCs and associated Core API Secret.
func ParseYAMLCatalogue(catalogue *Catalogue, r io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func catalogueSerializedHardware(catalogue *Catalogue, manifest []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func catalogueSerializedBMC(catalogue *Catalogue, manifest []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func catalogueSerializedSecret(catalogue *Catalogue, manifest []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalCatalogue marshals c into YAML that can be submitted to a Kubernetes cluster.
func MarshalCatalogue(c *Catalogue) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NewMachineCatalogueWriter creates a MachineWriter instance that writes Machine instances to
// catalogue including its Machine and Secret data.
func NewMachineCatalogueWriter(catalogue *Catalogue) MachineWriter {
	_ = "STUB: not implemented"
	return *new(MachineWriter)
}
