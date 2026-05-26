package collection

// Set is a collection that only contains unique elements.
type Set[T comparable] map[T]struct{}

// NewSet creates an empty Set.
func NewSet[T comparable]() Set[T] { _ = "STUB: not implemented"; return nil }

// NewSetFrom creates a Set from a list of elements.
func NewSetFrom[T comparable](elements ...T) Set[T] { _ = "STUB: not implemented"; return nil }

func newSet[T comparable](size int) Set[T] { _ = "STUB: not implemented"; return nil }

// Add stores a new element in the Set if wasn't contained yet.
func (s Set[T]) Add(e T) {
	_ = "STUB: not implemented"

	// Delete removes an element from the Set if it existed.
	return
}

func (s Set[T]) Delete(e T) {
	_ = "STUB: not implemented"

	// Contains checks if an element is contained in the Set.
	return
}

func (s Set[T]) Contains(e T) bool { _ = "STUB: not implemented"; return false }

// ToSlice generates a new slice with all elements in the Set.
// Order is non deterministic.
func (s Set[T]) ToSlice() []T { _ = "STUB: not implemented"; return nil }

// MapSet converts c to a new set. f is used to extract the value for representing each element
// of c.
func MapSet[G any, T comparable](c []G, f func(G) T) Set[T] { _ = "STUB: not implemented"; return nil }
