package test

import (
	"io"
	"os"
	"testing"

	"k8s.io/client-go/rest"
)

// WithFakeFile returns a throwaway file in a test-specific directory.
//
// The file is automatically closed and removed when the test ends.
func WithFakeFile(t *testing.T) (f *os.File) { _ = "STUB: not implemented"; return nil }

// WithFakeFileContents returns a throwaway file containing some data.
//
// The file is automatically closed and removed when the test ends.
func WithFakeFileContents(t *testing.T, r io.Reader) (f *os.File) {
	_ = "STUB: not implemented"
	return nil
}

// RemoveFileIfExists is a helper for ValidateFilename tests.
func RemoveFileIfExists(t *testing.T, filename string) { _ = "STUB: not implemented"; return }

// UseEnvTest sets up the controller-runtime EnvTest framework.
//
// The test will be skipped if EnvTest framework isn't detected.
//
// EnvTest provides fake k8s control plane components for testing
// purposes. The process of bringing up and tearing down the EnvTest framework
// involves running a few binaries, and is not integrated into a vanilla "go
// test" run, but rather can be run via the unit-test target in Makefile. See
// https://book.kubebuilder.io/reference/envtest.html for details.
//
// TODO: What could be done to integrate EnvTest with go test, so that "go
// test" would work?
func UseEnvTest(t *testing.T) *rest.Config {
	_ = "STUB: not implemented"
	// Detect if EnvTest has been set up.
	return nil
}

// By skipping this test, we allow traditional runs of go test from
// the command-line (or your editor) to work.
