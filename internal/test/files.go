package test

import (
	"flag"
	"regexp"
	"testing"

	"github.com/aws/eks-anywhere/pkg/files"
	"github.com/aws/eks-anywhere/pkg/filewriter"
)

var UpdateGoldenFiles = flag.Bool("update", false, "update golden files")

func AssertFilesEquals(t *testing.T, gotPath, wantPath string) { _ = "STUB: not implemented"; return }

func AssertContentToFile(t *testing.T, gotContent, wantFile string) {
	_ = "STUB: not implemented"
	return
}

func contentEqualToFile(gotContent []byte, wantFile string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func computeDiffBetweenContentAndFile(content []byte, file string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func processUpdate(t *testing.T, filePath, content string) { _ = "STUB: not implemented"; return }

func ReadFileAsBytes(t *testing.T, file string) []byte { _ = "STUB: not implemented"; return nil }

func ReadFile(t *testing.T, file string) string { _ = "STUB: not implemented"; return "" }

func NewWriter(t *testing.T) (dir string, writer filewriter.FileWriter) {
	_ = "STUB: not implemented"
	return "", *new(filewriter.FileWriter)
}

func cleanupDir(t *testing.T, dir string) func() { _ = "STUB: not implemented"; return nil }

var sanitizePathChars = regexp.MustCompile(`[^\w-]`)

const sanitizePathReplacementChar = "_"

// SanitizePath sanitizes s so its usable as a path name. For safety, it assumes all characters that are not
// A-Z, a-z, 0-9, _ or - are illegal and replaces them with _.
func SanitizePath(s string) string { _ = "STUB: not implemented"; return "" }

// NewFileReader builds a file reader with a proper user-agent.
// Unit tests should never make network call to the internet, but just in case we
// set the user-agent to be able to pin-point them here.
func NewFileReader() *files.Reader { _ = "STUB: not implemented"; return nil }
