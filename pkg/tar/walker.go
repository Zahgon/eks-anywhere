package tar

import (
	"path/filepath"
)

func NewFolderWalker(folder string) FolderWalker {
	_ = "STUB: not implemented"
	return *new(FolderWalker)
}

type FolderWalker struct {
	folder, folderPrefix string
}

func (f FolderWalker) Walk(fn TarFunc) error { _ = "STUB: not implemented"; return nil }

func (f FolderWalker) trimFolder(fn TarFunc) filepath.WalkFunc {
	_ = "STUB: not implemented"
	return *new(filepath.WalkFunc)
}
