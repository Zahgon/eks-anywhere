package tar

type Packager struct{}

func NewPackager() Packager { _ = "STUB: not implemented"; return *new(Packager) }

func (Packager) Package(sourceFolder, dstFile string) error { _ = "STUB: not implemented"; return nil }

func (Packager) UnPackage(orgFile, dstFolder string) error { _ = "STUB: not implemented"; return nil }
