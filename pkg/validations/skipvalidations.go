package validations

// string values of supported validation names that can be skipped.
const (
	PDB             = "pod-disruption"
	VSphereUserPriv = "vsphere-user-privilege"
	EksaVersionSkew = "eksa-version-skew"
)

// ValidSkippableValidationsMap returns a map for all valid skippable validations as keys, defaulting values to false.
// Defaulting to False means these validations won't be skipped unless set to True.
func validSkippableValidationsMap(skippableValidations []string) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

// ValidateSkippableValidation validates if provided validations are supported by EKSA to skip for upgrades.
func ValidateSkippableValidation(skippedValidations []string, skippableValidations []string) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
