package hardware

// MachineReader reads single Machine configuration at a time. When there are no more Machine entries
// to be read, Read() returns io.EOF.
type MachineReader interface {
	Read() (Machine, error)
}

// MachineWriter writes Machine entries.
type MachineWriter interface {
	Write(Machine) error
}

// MachineValidator validates an instance of Machine.
type MachineValidator interface {
	Validate(Machine) error
}

// TranslateAll reads entries 1 at a time from reader and writes them to writer. When reader returns io.EOF,
// TranslateAll returns nil. Failure to return io.EOF from reader will result in an infinite loop.
func TranslateAll(reader MachineReader, writer MachineWriter, validator MachineValidator) error {
	_ = "STUB: not implemented"
	return nil
}

// Translate reads 1 entry from reader and writes it to writer. When reader returns io.EOF Translate
// returns io.EOF to the caller.
func Translate(reader MachineReader, writer MachineWriter, validator MachineValidator) error {
	_ = "STUB: not implemented"
	return nil
}
