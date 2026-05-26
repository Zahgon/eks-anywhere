package files

import (
	"archive/tar"
)

// GzipFileDownloadExtract downloads and extracts a specific file to destination.
func GzipFileDownloadExtract(uri, fileToExtract, destination string) error {
	_ = "STUB: not implemented"
	return nil
}

type singleFileRouter struct {
	folder   string
	fileName string
}

func (s singleFileRouter) ExtractPath(header *tar.Header) string {
	_ = "STUB: not implemented"
	return ""
}
