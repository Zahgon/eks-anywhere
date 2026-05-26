package test

import "time"

var (
	defaultNow = time.Unix(0, 1234567890000*int64(time.Millisecond))
	newNow     = time.Unix(0, 987654321000*int64(time.Millisecond))
)

func FakeNow() time.Time {
	_ = "STUB: not implemented"

	// NewFakeNow sets a dummy value for time.Now in unit tests.
	// This is particularly useful when testing upgrade operations.
	return *new(time.Time)
}

func NewFakeNow() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
