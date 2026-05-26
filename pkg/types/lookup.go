package types

type Lookup map[string]struct{}

func (l Lookup) IsPresent(v string) bool { _ = "STUB: not implemented"; return false }

func (l Lookup) ToSlice() []string { _ = "STUB: not implemented"; return nil }

func SliceToLookup(slice []string) Lookup { _ = "STUB: not implemented"; return *new(Lookup) }
