package file

import (
	"io"
)

// ReadFile function reads the contents of a file and provides them as an `io.Reader`.
func ReadFile(fileName string) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
