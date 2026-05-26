package validations

type ValidationResult struct {
	Name        string
	Err         error
	Remediation string
	Silent      bool
}

func (v *ValidationResult) Report() { _ = "STUB: not implemented"; return }

func (v *ValidationResult) LogPass() { _ = "STUB: not implemented"; return }

func capitalize(s string) string { _ = "STUB: not implemented"; return "" }
