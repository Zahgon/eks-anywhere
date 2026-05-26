package hardware

import (
	tinkv1alpha1 "github.com/tinkerbell/tink/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"

	eksav1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// serializeHardwareSelector returns a key for use in a map unique selector.
func serializeHardwareSelector(selector eksav1alpha1.HardwareSelector) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// ErrDiskNotFound indicates a disk was not found for a given selector.
		nil
}

type ErrDiskNotFound struct {
	// A unique identifier for the selector, preferrably something useful to an end-user.
	SelectorID string
}

func (e ErrDiskNotFound) Error() string { _ = "STUB: not implemented"; return "" }

func (ErrDiskNotFound) Is(t error) bool { _ = "STUB: not implemented"; return false }

// IndexHardware indexes Hardware instances on index by extracfting the key using fn.
func (c *Catalogue) IndexHardware(index string, fn KeyExtractorFunc) {
	_ = "STUB: not implemented"
	return
}

// InsertHardware inserts Hardware into the catalogue. If any indexes exist, the hardware is
// indexed.
func (c *Catalogue) InsertHardware(hardware *tinkv1alpha1.Hardware) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveHardwares removes a slice of hardwares from the catalogue.
func (c *Catalogue) RemoveHardwares(hardware []tinkv1alpha1.Hardware) error {
	_ = "STUB: not implemented"
	return nil
}

// getRemoveKey returns key used to search and remove hardware.
func getRemoveKey(hardware tinkv1alpha1.Hardware) string { _ = "STUB: not implemented"; return "" }

// AllHardware retrieves a copy of the catalogued Hardware instances.
func (c *Catalogue) AllHardware() []*tinkv1alpha1.Hardware { _ = "STUB: not implemented"; return nil }

// LookupHardware retrieves Hardware instances on index with a key of key. Multiple hardware _may_
// have the same key hence it can return multiple Hardware.
func (c *Catalogue) LookupHardware(index, key string) ([]*tinkv1alpha1.Hardware, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalHardware returns the total hardware registered in the catalogue.
func (c *Catalogue) TotalHardware() int { _ = "STUB: not implemented"; return 0 }

const HardwareIDIndex = ".Spec.Metadata.Instance.ID"

// WithHardwareIDIndex creates a Hardware index using HardwareIDIndex on .Spec.Metadata.Instance.ID
// values.
func WithHardwareIDIndex() CatalogueOption { _ = "STUB: not implemented"; return *new(CatalogueOption) }

const HardwareBMCRefIndex = ".Spec.BmcRef"

// WithHardwareBMCRefIndex creates a Hardware index using HardwareBMCRefIndex on .Spec.BmcRef.
func WithHardwareBMCRefIndex() CatalogueOption {
	_ = "STUB: not implemented"
	return *new(CatalogueOption)
}

// HardwareCatalogueWriter converts Machine instances to Tinkerbell Hardware and inserts them
// in a catalogue.
type HardwareCatalogueWriter struct {
	catalogue *Catalogue
}

var _ MachineWriter = &HardwareCatalogueWriter{}

// NewHardwareCatalogueWriter creates a new HardwareCatalogueWriter instance.
func NewHardwareCatalogueWriter(catalogue *Catalogue) *HardwareCatalogueWriter {
	_ = "STUB: not implemented"
	return nil
}

// Write converts m to a Tinkerbell Hardware and inserts it into w's Catalogue.
func (w *HardwareCatalogueWriter) Write(m Machine) error { _ = "STUB: not implemented"; return nil }

func hardwareFromMachine(m Machine) *tinkv1alpha1.Hardware {
	_ = "STUB: not implemented"
	// allow is necessary to allocate memory so we can get a bool pointer required by
	// the hardware.
	return nil
}

// TODO(chrisdoherty4) Set the namespace to the CAPT namespace.

// TODO(chrisdoherty4) Fix upstream. The OperatingSystem is used in boots to
// detect what iPXE scripts should be served. The Kubernetes back-end nilifies
// its response to retrieving the OS data and the handling code doesn't check
// for nil resulting in a segfault.
//
// Upstream needs patching but this will suffice for now.

// set LeaseTime to the max value so it effectively hands out max duration leases (~136 years)
// This value gets ignored for Ubuntu because we set static IPs for it
// It's only temporarily needed for Bottlerocket until Bottlerocket supports static IPs

// newBMCRefFromMachine returns a BMCRef pointer for Hardware.
func newBMCRefFromMachine(m Machine) *corev1.TypedLocalObjectReference {
	_ = "STUB: not implemented"
	return nil
}
