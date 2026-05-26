package aflag

import (
	"github.com/spf13/pflag"
)

// MarkRequired is a helper to mark flags required on cmd. If a flag does not exist, it panics.
func MarkRequired(set *pflag.FlagSet, flags ...string) { _ = "STUB: not implemented"; return }

// MarkHidden is a helper to mark flags hidden on cmd. If a flag does not exist, it panics.
func MarkHidden(set *pflag.FlagSet, flags ...string) { _ = "STUB: not implemented"; return }
