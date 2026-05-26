package collection

// ToMap is a utility function that converts the slice s to a map using key to retrieve map
// keys.
func ToMap[K comparable, V any](s []V, key func(V) K) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// ToSlice is a utility function that converts the map m to a slice of m's values. The returned
// slices order is undefined.
func ToSlice[K comparable, V any](m map[K]V) []V { _ = "STUB: not implemented"; return nil }
