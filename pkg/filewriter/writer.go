package filewriter

import (
	"io"
	"io/fs"
)

type writer struct {
	dir     string
	tempDir string
}

func NewWriter(dir string) (FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(FileWriter), nil
}

func (w *writer) Write(fileName string, content []byte, opts ...FileOptionsFunc) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (w *writer) WithDir(dir string) (FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(FileWriter), nil
}

func (w *writer) Dir() string { _ = "STUB: not implemented"; return "" }

func (w *writer) TempDir() string { _ = "STUB: not implemented"; return "" }

func (w *writer) CleanUp() { _ = "STUB: not implemented"; return }

func (w *writer) CleanUpTemp() { _ = "STUB: not implemented"; return }

// Create creates a file with the given name rooted at w's base directory.
func (w *writer) Create(name string, opts ...FileOptionsFunc) (_ io.WriteCloser, path string, _ error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), "", nil
}

// Delete removes a file with the given name from w's base directory.
func (w *writer) Delete(fileName string) error { _ = "STUB: not implemented"; return nil }

// Check if file exists first

// File doesn't exist, nothing to delete

type options struct {
	BasePath    string
	Permissions fs.FileMode
}

// buildOptions converts a set of FileOptionsFunc's to a single options struct.
func buildOptions(w *writer, opts []FileOptionsFunc) options {
	_ = "STUB: not implemented"
	return *new(options)
}
