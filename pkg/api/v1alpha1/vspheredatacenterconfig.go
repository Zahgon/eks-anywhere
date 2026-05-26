package v1alpha1

const VSphereDatacenterKind = "VSphereDatacenterConfig"

type folderType string

const (
	networkFolderType folderType = "network"
)

// Used for generating yaml for generate clusterconfig command.
func NewVSphereDatacenterConfigGenerate(clusterName string) *VSphereDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *VSphereDatacenterConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *VSphereDatacenterConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *VSphereDatacenterConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func GetVSphereDatacenterConfig(fileName string) (*VSphereDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateFullVCenterPath(foldType folderType, folderPath string, datacenter string) string {
	_ = "STUB: not implemented"
	return ""
}

func validatePath(foldType folderType, folderPath string, datacenter string) error {
	_ = "STUB: not implemented"
	return nil
}
