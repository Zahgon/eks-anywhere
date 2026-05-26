package hardware

// NormalizerFunc applies a normalization transformation to the Machine.
type NormalizerFunc func(Machine) Machine

// Normalizer is a decorator for a MachineReader that applies a set of normalization funcs
// to machines.
type Normalizer struct {
	reader      MachineReader
	normalizers []NormalizerFunc
}

// NewNormalizer creates a Normalizer instance that decorates r's Read(). A set of default
// normalization functions are pre-registered.
func NewNormalizer(r MachineReader) *Normalizer { _ = "STUB: not implemented"; return nil }

// NewRawNormalizer returns a Normalizer with default normalizations registered by
// RegisterDefaultNormalizations.
func NewRawNormalizer(r MachineReader) *Normalizer { _ = "STUB: not implemented"; return nil }

// Read reads an Machine from the decorated MachineReader, applies all normalization funcs and
// returns the machine. If the decorated MachineReader errors, it is returned.
func (n Normalizer) Read() (Machine, error) { _ = "STUB: not implemented"; return *new(Machine), nil }

// Register fn to n such that fn is run over each machine read from the wrapped MachineReader.
func (n *Normalizer) Register(fn NormalizerFunc) { _ = "STUB: not implemented"; return }

// LowercaseMACAddress ensures m's MACAddress field has lower chase characters.
func LowercaseMACAddress(m Machine) Machine { _ = "STUB: not implemented"; return *new(Machine) }

// RegisterDefaultNormalizations registers a set of default normalizations on n.
func RegisterDefaultNormalizations(n *Normalizer) { _ = "STUB: not implemented"; return }
