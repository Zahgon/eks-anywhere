package validation

// Remediable is an error that provides a possible remediation.
type Remediable interface {
	Remediation() string
}

// remediableError implements Fixable around a generic error.
type remediableError struct {
	error
	remediation string
}

// Remediation returns a possible solution to the error.
func (e *remediableError) Remediation() string { _ = "STUB: not implemented"; return "" }

// NewRemediableErr returns a new [Remediable] error.
func NewRemediableErr(err, remediation string) error { _ = "STUB: not implemented"; return nil }

// WithRemediation makes an error [Remediable].
func WithRemediation(err error, remediation string) error { _ = "STUB: not implemented"; return nil }

// IsRemediable checks if an error has a remediation.
func IsRemediable(err error) bool { _ = "STUB: not implemented"; return false }

// Remediation returns the Remediation message for an error if it has it.
// Otherwise it returns an empty string.
func Remediation(err error) string { _ = "STUB: not implemented"; return "" }
