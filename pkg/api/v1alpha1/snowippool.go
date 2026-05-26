package v1alpha1

const (
	// SnowIPPoolKind is the object kind name for SnowIPPool.
	SnowIPPoolKind = "SnowIPPool"
)

// SnowIPPoolsSliceEqual compares and returns whether two snow IPPool objects are equal.
func SnowIPPoolsSliceEqual(a, b []IPPool) bool { _ = "STUB: not implemented"; return false }

func generateKeyForIPPool(pool IPPool) string { _ = "STUB: not implemented"; return "" }

func validateSnowIPPool(pool *SnowIPPool) error {
	_ = "STUB: not implemented" //nolint:gocyclo
	return nil
}
