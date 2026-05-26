package hardware

import (
	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/rufio"
)

// GofishProviderOption is the provider name for Redfish Provider in Rufio.
const GofishProviderOption = "gofish"

// IndexBMCs indexes BMC instances on index by extracfting the key using fn.
func (c *Catalogue) IndexBMCs(index string, fn KeyExtractorFunc) { _ = "STUB: not implemented"; return }

// InsertBMC inserts BMCs into the catalogue. If any indexes exist, the BMC is indexed.
func (c *Catalogue) InsertBMC(bmc *v1alpha1.Machine) error { _ = "STUB: not implemented"; return nil }

// AllBMCs retrieves a copy of the catalogued BMC instances.
func (c *Catalogue) AllBMCs() []*v1alpha1.Machine { _ = "STUB: not implemented"; return nil }

// LookupBMC retrieves BMC instances on index with a key of key. Multiple BMCs _may_
// have the same key hence it can return multiple BMCs.
func (c *Catalogue) LookupBMC(index, key string) ([]*v1alpha1.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalBMCs returns the total BMCs registered in the catalogue.
func (c *Catalogue) TotalBMCs() int { _ = "STUB: not implemented"; return 0 }

const BMCNameIndex = ".ObjectMeta.Name"

// WithBMCNameIndex creates a BMC index using BMCNameIndex on .ObjectMeta.Name.
func WithBMCNameIndex() CatalogueOption { _ = "STUB: not implemented"; return *new(CatalogueOption) }

// BMCCatalogueWriter converts Machine instances to Tinkerbell Machine and inserts them
// in a catalogue.
type BMCCatalogueWriter struct {
	catalogue *Catalogue
}

var _ MachineWriter = &BMCCatalogueWriter{}

// NewBMCCatalogueWriter creates a new BMCCatalogueWriter instance.
func NewBMCCatalogueWriter(catalogue *Catalogue) *BMCCatalogueWriter {
	_ = "STUB: not implemented"
	return nil
}

// Write converts m to a Tinkerbell Machine and inserts it into w's Catalogue.
func (w *BMCCatalogueWriter) Write(m Machine) error { _ = "STUB: not implemented"; return nil }

func toRufioMachine(m Machine) *v1alpha1.Machine {
	_ = "STUB: not implemented"
	// TODO(chrisdoherty4)
	//   - Set the namespace to the CAPT namespace.
	//   - Patch through insecure TLS.
	return nil
}

// Redfish bmc client generally seems to be more reliable in bmc interactions
// compared to other clients. Prefer Redfish client if available

func toRPCOptions(r *RPCOpts, m Machine) *v1alpha1.RPCOptions {
	_ = "STUB: not implemented"
	return nil
}

func toRequestOpts(r RequestOpts) *v1alpha1.RequestOpts { _ = "STUB: not implemented"; return nil }

func toSignatureOpts(s SignatureOpts) *v1alpha1.SignatureOpts {
	_ = "STUB: not implemented"
	return nil
}

func toHMACOpts(h HMACOpts, m Machine) *v1alpha1.HMACOpts { _ = "STUB: not implemented"; return nil }

func toExperimentalOpts(e ExperimentalOpts) *v1alpha1.ExperimentalOpts {
	_ = "STUB: not implemented"
	return nil
}
