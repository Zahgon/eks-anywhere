// Package aflag is the eks anywhere flag handling package.
package aflag

import (
	"github.com/spf13/pflag"
)

// Flag defines a CLI flag.
type Flag[T any] struct {
	Name    string
	Short   string
	Usage   string
	Default T
}

// String applies f to fs and writes the value to dst.
func String(f Flag[string], dst *string, fs *pflag.FlagSet) {
	_ = "STUB: not implemented"

	// With short form
	return
}

// Without short form

// Bool applies f to fs and writes the value to dst.
func Bool(f Flag[bool], dst *bool, fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// StringSlice applies f to fs and writes the value to dst.
func StringSlice(f Flag[[]string], dst *[]string, fs *pflag.FlagSet) {
	_ = "STUB: not implemented"
	return
}

// StringString applies f to fs and writes the value to dst.
func StringString(f Flag[map[string]string], dst *map[string]string, fs *pflag.FlagSet) {
	_ = "STUB: not implemented"
	return
}

// HTTPHeader applies f to fs and writes the value to dst.
func HTTPHeader(f Flag[Header], dst *Header, fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }
