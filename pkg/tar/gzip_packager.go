package tar

type GzipPackager struct{}

func NewGzipPackager() GzipPackager { _ = "STUB: not implemented"; return *new(GzipPackager) }

func (GzipPackager) Package(sourceFolder, dstFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (GzipPackager) UnPackage(orgFile, dstFolder string) error {
	_ = "STUB: not implemented"
	return nil
}
