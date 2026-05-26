package v1alpha1

const NutanixDatacenterKind = "NutanixDatacenterConfig"

// NewNutanixDatacenterConfigGenerate is used for generating yaml for generate clusterconfig command.
func NewNutanixDatacenterConfigGenerate(clusterName string) *NutanixDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *NutanixDatacenterConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *NutanixDatacenterConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *NutanixDatacenterConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

// GetNutanixDatacenterConfig parses config in a yaml file and returns a NutanixDatacenterConfig object.
func GetNutanixDatacenterConfig(fileName string) (*NutanixDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
