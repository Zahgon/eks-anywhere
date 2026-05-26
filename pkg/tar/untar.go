package tar

import (
	"io"
)

func UntarFile(tarFile, dstFolder string) error { _ = "STUB: not implemented"; return nil }

func Untar(source io.Reader, router Router) error { _ = "STUB: not implemented"; return nil }

// Prevent malicous directory traversals.
// https://cwe.mitre.org/data/definitions/22.html
