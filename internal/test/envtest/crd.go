package envtest

import (
	"regexp"
)

type moduleWithCRD struct {
	pkg          string
	crdPaths     []string
	requireRegex *regexp.Regexp
	replaceRegex *regexp.Regexp
}

func mustBuildModuleWithCRDs(p string, opts ...moduleOpt) moduleWithCRD {
	_ = "STUB: not implemented"
	return *new(moduleWithCRD)
}

func withAdditionalCustomCRDPath(customCRDPath string) moduleOpt {
	_ = "STUB: not implemented"
	return *new(moduleOpt)
}

func withMainCustomCRDPath(customCRDPath string) moduleOpt {
	_ = "STUB: not implemented"
	return *new(moduleOpt)
}

type moduleOpt func(*moduleWithCRD)

func buildModuleWithCRD(pkg string, opts ...moduleOpt) (*moduleWithCRD, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type moduleInDisk struct {
	moduleWithCRD
	path, name, version string
}

func (m moduleInDisk) pathsToCRDs() []string { _ = "STUB: not implemented"; return nil }

func pathToCRDs(path, name, version, crdPath string) string { _ = "STUB: not implemented"; return "" }

func getPathsToPackagesCRDs(rootFolder string, packages ...moduleWithCRD) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the package has already been mapped to disk, it was
// probably by a replace, don't overwrite

func envOrDefault(envKey, defaultValue string) string { _ = "STUB: not implemented"; return "" }

func buildModulesMappedToDisk(modules []moduleWithCRD) map[string]*moduleInDisk {
	_ = "STUB: not implemented"
	return nil
}
