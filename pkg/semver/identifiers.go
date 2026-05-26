package semver

// identifier represents a single identifier in a semver version.
// identifiers are dot separated strings or numbers in the pre release
// and build metadata.
type identifier struct {
	intValue uint64
	strValue string
	isNum    bool
}

func newIdentifier(s string) identifier { _ = "STUB: not implemented"; return *new(identifier) }

// compare compares v and o.
// -1 == v is less than o.
// 0 == v is equal to o.
// 1 == v is greater than o.
// 2 == v is different than o (it is not possible to identify if lower or greater).
// Number is considered lower than string.
func (v identifier) compare(o identifier) int { _ = "STUB: not implemented"; return 0 }

// both are numbers

// both are strings

// In order to support random identifiers, like commit hashes,
// we return 2 when the strings are different to signal the
// identifiers are different but we can't determine the precedence

type identifiers []identifier

func newIdentifiers(ids []string) identifiers { _ = "STUB: not implemented"; return *new(identifiers) }

// compare compares 2 identifiers v and o.
// -1 == v is less than o.
// 0 == v is equal to o.
// 1 == v is greater than o.
// If everything else is equal the longer identifier is greater.
func (v identifiers) compare(o identifiers) int { _ = "STUB: not implemented"; return 0 }

// if everything is equal until now the longer is greater
