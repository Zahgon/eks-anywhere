package yaml

// Join joins YAML resources into a single YAML document. It does not validate individual
// resources.
func Join(resources [][]byte) []byte { _ = "STUB: not implemented"; return nil }

// Serialize serializes objects into YAML documents.
func Serialize[T any](objs ...T) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }
