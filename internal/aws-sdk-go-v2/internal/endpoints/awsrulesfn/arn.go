package awsrulesfn

// ARN provides AWS ARN components broken out into a data structure.
type ARN struct {
	Partition  string
	Service    string
	Region     string
	AccountId  string
	ResourceId OptionalStringSlice
}

const (
	arnDelimiters      = ":"
	resourceDelimiters = "/:"
	arnSections        = 6
	arnPrefix          = "arn:"

	// zero-indexed
	sectionPartition = 1
	sectionService   = 2
	sectionRegion    = 3
	sectionAccountID = 4
	sectionResource  = 5
)

// ParseARN returns an [ARN] value parsed from the input string provided. If
// the ARN cannot be parsed nil will be returned, and error added to
// [ErrorCollector].
func ParseARN(input string) *ARN { _ = "STUB: not implemented"; return nil }

// splitResource splits the resource components by the ARN resource delimiters.
func splitResource(v string) []string { _ = "STUB: not implemented"; return nil }

// OptionalStringSlice provides a helper to safely get the index of a string
// slice that may be out of bounds. Returns pointer to string if index is
// valid. Otherwise returns nil.
type OptionalStringSlice []string

// Get returns a string pointer of the string at index i if the index is valid.
// Otherwise returns nil.
func (s OptionalStringSlice) Get(i int) *string { _ = "STUB: not implemented"; return nil }
