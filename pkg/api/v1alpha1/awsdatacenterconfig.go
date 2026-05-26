package v1alpha1

const AWSDatacenterKind = "AWSDatacenterConfig"

// Used for generating yaml for generate clusterconfig command.
func NewAWSDatacenterConfigGenerate(clusterName string) *AWSDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *AWSDatacenterConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *AWSDatacenterConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *AWSDatacenterConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func GetAWSDatacenterConfig(fileName string) (*AWSDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
