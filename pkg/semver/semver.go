package semver

import (
	"regexp"
)

const semverRegex = `^v?(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

var semverRegexp = regexp.MustCompile(semverRegex)

type Version struct {
	Major, Minor, Patch       int64
	Prerelease, Buildmetadata string
}

func New(version string) (*Version, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *Version) SameMajor(v2 *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) SameMinor(v2 *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) SamePatch(v2 *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) SamePrerelease(v2 *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) Equal(v2 *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) GreaterThan(v2 *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) LessThan(v2 *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) Compare(v2 *Version) int { _ = "STUB: not implemented"; return 0 }

// CompareBuildMetadata compares the build metadata of v and v2.
// The metadata is split in its identifiers and these compared one by one.
// Number identifiers are considered lower than strings.
// If one build metadata is a prefix of the other, the longer one is considered greater.
// -1 == v is less than v2.
// 0 == v is equal to v2.
// 1 == v is greater than v2.
// 2 == v is different than v2 (it is not possible to identify if lower or greater).
func (v *Version) CompareBuildMetadata(v2 *Version) int { _ = "STUB: not implemented"; return 0 }

func (v *Version) buildIdentifiers() identifiers {
	_ = "STUB: not implemented"
	return *new(identifiers)
}

func (v *Version) String() string { _ = "STUB: not implemented"; return "" }

func compare(i, i2 int64) int { _ = "STUB: not implemented"; return 0 }
