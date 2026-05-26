package tar

import (
	"archive/tar"
)

// Router instructs where to extract a file.
type Router interface {
	// ExtractPath instructs the path where a file should be extracted.
	// Empty strings instructs to omit the file extraction
	ExtractPath(header *tar.Header) string
}

type FolderRouter struct {
	folder string
}

func NewFolderRouter(folder string) FolderRouter {
	_ = "STUB: not implemented"
	return *new(FolderRouter)
}

func (f FolderRouter) ExtractPath(header *tar.Header) string { _ = "STUB: not implemented"; return "" }
