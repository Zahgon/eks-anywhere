package tar

import (
	"archive/tar"
	"io"
	"os"
)

func TarFolder(sourceFolder, dstFile string) error { _ = "STUB: not implemented"; return nil }

func tarFolderToWriter(sourceFolder string, dst io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

type TarFunc func(file string, info os.FileInfo, header *tar.Header) error

type Walker interface {
	Walk(TarFunc) error
}

func Tar(source Walker, dst io.Writer) error { _ = "STUB: not implemented"; return nil }

func addToTar(tw *tar.Writer) TarFunc { _ = "STUB: not implemented"; return *new(TarFunc) }
