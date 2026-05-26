package logger

import (
	"sync"

	"github.com/go-logr/logr"
)

// MaxLogLevel denotes the maximum log level supported by the logger package.
const MaxLogLevel = 9

const (
	markPass    = "✅ "
	markSuccess = "🎉 "
	markFailed  = "❌ "
	markWarning = "⚠️"
)

var (
	pkgLogger    logr.Logger = logr.Discard()
	pkgLoggerMtx sync.RWMutex
)

func setLogger(l logr.Logger) { _ = "STUB: not implemented"; return }

// Get returns the logger instance that has been previously set.
// If no logger has been set, it returns a null logger.
func Get() logr.Logger { _ = "STUB: not implemented"; return *new(logr.Logger) }

// MaxLogging determines if the package logger is configured to log at MaxLogLevel.
func MaxLogging() bool { _ = "STUB: not implemented"; return false }

// Fatal is equivalent to Get().Error() followed by a call to os.Exit(1).
func Fatal(err error, msg string) { _ = "STUB: not implemented"; return }

// Info logs a non-error message with the given key/value pairs as context.
//
// The msg argument should be used to add some constant description to
// the log line. The key/value pairs can then be used to add additional
// variable information. The key/value pairs should alternate string
// keys and arbitrary values.
func Info(msg string, keysAndValues ...interface{}) { _ = "STUB: not implemented"; return }

// V returns an Logger value for a specific verbosity level, relative to
// this Logger. In other words, V values are additive.  V higher verbosity
// level means a log message is less important.  It's illegal to pass a log
// level less than zero.
func V(level int) logr.Logger {
	_ = "STUB: not implemented"
	return *

	// Error logs an error message using the package logger.
	new(logr.Logger)
}

func Error(err error, msg string, keysAndValues ...interface{}) { _ = "STUB: not implemented"; return }

// MarkPass logs a message prefixed with a green check emoji.
func MarkPass(msg string, keysAndValues ...interface{}) { _ = "STUB: not implemented"; return }

// MarkSuccess logs a message prefixed with a popper emoji.
func MarkSuccess(msg string, keysAndValues ...interface{}) { _ = "STUB: not implemented"; return }

// MarkFail logs a message prefixed with a cross emoji.
func MarkFail(msg string, keysAndValues ...interface{}) { _ = "STUB: not implemented"; return }

// MarkWarning logs a message prefixed with a warning mark.
func MarkWarning(msg string, keysAndValues ...interface{}) { _ = "STUB: not implemented"; return }
