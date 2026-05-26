package framework

import "testing"

// T defines test support functionality, ala the Go stdlib testing.T.
//
// Being able to change its implementation supports logging functionality for
// test support methods that are executed outside of a test, such as when
// bringing up or tearing down test clusters that will be used and re-used
// throughout multiple tests.
//
// Only those methods currently in use are defined. Add more methods from
// stdlib testing.T as necessary.
type T interface {
	Cleanup(func())
	Error(...any)
	Errorf(string, ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(...any)
	Fatalf(string, ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Parallel()
	Run(string, func(*testing.T)) bool
	Setenv(string, string)
	Skip(...any)
	SkipNow()
	Skipf(string, ...any)
	Skipped() bool
	TempDir() string
}

// T ensures that *testing.T implements T, to detect API drift.
var _ T = (*testing.T)(nil)

// LoggingOnlyT implements select logging and error handling functionality of T.
//
// Most non-logging, non-error reporting methods will simply panic.
type LoggingOnlyT struct{}

// NewLoggingOnlyT creates a LoggingOnlyT, which does what its name implies.
func NewLoggingOnlyT() *LoggingOnlyT { _ = "STUB: not implemented"; return nil }

// Cleanup implements T.
func (t LoggingOnlyT) Cleanup(_ func()) { _ = "STUB: not implemented"; return }

// Error implements T.
func (t LoggingOnlyT) Error(_ ...any) { _ = "STUB: not implemented"; return }

// Errorf implements T.
func (t LoggingOnlyT) Errorf(_ string, _ ...any) { _ = "STUB: not implemented"; return }

// Fail implements T.
func (t LoggingOnlyT) Fail() { _ = "STUB: not implemented"; return }

// FailNow implements T.
func (t LoggingOnlyT) FailNow() { _ = "STUB: not implemented"; return }

// Failed implements T.
func (t LoggingOnlyT) Failed() bool { _ = "STUB: not implemented"; return false }

// Fatal implements T.
func (t LoggingOnlyT) Fatal(_ ...any) { _ = "STUB: not implemented"; return }

// panic("LoggingOnlyT implements only the logging methods of T")

// Fatalf implements T.
func (t LoggingOnlyT) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Helper implements T.
func (t LoggingOnlyT) Helper() { _ = "STUB: not implemented"; return }

// Log implements T.
func (t LoggingOnlyT) Log(args ...any) { _ = "STUB: not implemented"; return }

// Logf implements T.
func (t LoggingOnlyT) Logf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Name implements T.
func (t LoggingOnlyT) Name() string { _ = "STUB: not implemented"; return "" }

// Parallel implements T.
func (t LoggingOnlyT) Parallel() { _ = "STUB: not implemented"; return }

// Run implements T.
func (t LoggingOnlyT) Run(_ string, _ func(*testing.T)) bool {
	_ = "STUB: not implemented"
	return false
}

// Setenv implements T.
func (t LoggingOnlyT) Setenv(_ string, _ string) { _ = "STUB: not implemented"; return }

// Skip implements T.
func (t LoggingOnlyT) Skip(_ ...any) { _ = "STUB: not implemented"; return }

// SkipNow implements T.
func (t LoggingOnlyT) SkipNow() { _ = "STUB: not implemented"; return }

// Skipf implements T.
func (t LoggingOnlyT) Skipf(_ string, _ ...any) { _ = "STUB: not implemented"; return }

// Skipped implements T.
func (t LoggingOnlyT) Skipped() bool { _ = "STUB: not implemented"; return false }

// TempDir implements T.
func (t LoggingOnlyT) TempDir() string { _ = "STUB: not implemented"; return "" }
