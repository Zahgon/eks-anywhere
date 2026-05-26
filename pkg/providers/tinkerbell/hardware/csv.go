package hardware

import (
	"io"

	csv "github.com/gocarina/gocsv"
)

// CSVReader reads a CSV file and provides Machine instances. It satisfies the MachineReader interface. The ID field of
// the Machine is optional in the CSV. If unspecified, CSVReader will generate a UUID and apply it to the machine.
type CSVReader struct {
	reader *csv.Unmarshaller
	// BMCOptions used in a Machine that do not have a corresponding column in the CSV.
	BMCOptions *BMCOptions
}

// NewCSVReader returns a new CSVReader instance that consumes csv data from r. r should return io.EOF when no more
// records are available.
func NewCSVReader(r io.Reader, opts *BMCOptions) (CSVReader, error) {
	_ = "STUB: not implemented"
	return *new(CSVReader), nil
}

// Read reads a single entry from the CSV data source and returns a new Machine representation.
func (cr CSVReader) Read() (Machine, error) { _ = "STUB: not implemented"; return *new(Machine), nil }

// NewNormalizedCSVReaderFromFile creates a MachineReader instance backed by a CSVReader reading from path
// that applies default normalizations to machines.
func NewNormalizedCSVReaderFromFile(path string, opts *BMCOptions) (MachineReader, error) {
	_ = "STUB: not implemented"
	return *new(MachineReader), nil
}

// requiredColumns matches the csv tags on the Machine struct. These must remain in sync with
// the struct. We may consider an alternative that uses reflection to interpret whether a field
// is required in the future.
var requiredColumns = map[string]struct{}{
	"hostname":    {},
	"ip_address":  {},
	"netmask":     {},
	"gateway":     {},
	"nameservers": {},
	"mac":         {},
	"disk":        {},
	"labels":      {},
}

func ensureRequiredColumnsInCSV(unmatched []string) error { _ = "STUB: not implemented"; return nil }

// BuildHardwareYAML builds a hardware yaml from the csv at the provided path.
func BuildHardwareYAML(path string, opts *BMCOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
