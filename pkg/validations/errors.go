package validations

type ValidationError struct {
	Errs []string
}

func (v *ValidationError) Error() string { _ = "STUB: not implemented"; return "" }

func (v *ValidationError) String() string { _ = "STUB: not implemented"; return "" }
