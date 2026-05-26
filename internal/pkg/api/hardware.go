package api

import (
	"io"

	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
)

const (
	HardwareVendorDell        = "dell"
	HardwareVendorHP          = "hp"
	HardwareVendorSuperMicro  = "supermicro"
	HardwareVendorUnspecified = "unspecified"
	HardwareLabelTypeKeyName  = "type"
	ControlPlane              = "control-plane"
	Worker                    = "worker"
	ExternalEtcd              = "etcd"
)

// Alias for backwards compatibility.
type Hardware = hardware.Machine

func NewHardwareSlice(r io.Reader) ([]*Hardware, error) { _ = "STUB: not implemented"; return nil, nil }

func NewHardwareSliceFromFile(file string) ([]*Hardware, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHardwareMapFromFile(file string) (map[string]*Hardware, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// converts a hardware slice to a map. The first instance of the slice is used in case slice contains duplicates.
func HardwareSliceToMap(slice []*Hardware) map[string]*Hardware {
	_ = "STUB: not implemented"
	return nil
}

func WriteHardwareSliceToCSV(hardware []*Hardware, csvFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func HardwareMapToSlice(hardware map[string]*Hardware) []*Hardware {
	_ = "STUB: not implemented"
	return nil
}

func WriteHardwareMapToCSV(hardware map[string]*Hardware, csvFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func SplitHardware(slice []*Hardware, chunkSize int) [][]*Hardware {
	_ = "STUB: not implemented"
	return nil
}

// check slice capacity
