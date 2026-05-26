package logger

import (
	"go.uber.org/zap"
)

// Init initializes the package logger. Repeat calls will overwrite the package logger which may
// result in unexpected behavior.
func Init(opts Options) error { _ = "STUB: not implemented"; return nil }

// Level 4 and above are used for debugging and we want a different log structure for debug
// logs.

// Because we use negated levels it is necessary to negate the level again so the
// output appears in a V0 format.
//
// See logrAtomicLevel().

// zapcore.Open creates a noop logger if no paths are passed. Using a slice ensures we expand
// the slice to nothing when opts.OutputFilePath is unset.

// Build the encoders and logger.

// Configure package state so the logger can be used by other packages.

// Options represents a set of arguments for initializing the zap logger.
type Options struct {
	// Level is the log level at which to configure the logger from 0 to 9.
	Level int

	// OutputFilePath is an absolute file path. The file will be created if it doesn't exist.
	// All logs available at level 9 will be written to the file.
	OutputFilePath string
}

// logrAtomicLevel creates a zapcore.AtomicLevel compatible with go-logr.
func logrAtomicLevel(level int) zap.AtomicLevel {
	_ = "STUB: not implemented"
	// The go-logr wrapper uses custom Zap log levels. To represent this in Zap, its
	// necessary to negate the level to circumvent Zap level constraints.
	//
	// See https://github.com/go-logr/zapr/blob/master/zapr.go#L50.
	return *new(zap.AtomicLevel)
}
