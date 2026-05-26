package v1alpha1

const SnowDatacenterKind = "SnowDatacenterConfig"

// Used for generating yaml for generate clusterconfig command.
func NewSnowDatacenterConfigGenerate(clusterName string) *SnowDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (s *SnowDatacenterConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (s *SnowDatacenterConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (s *SnowDatacenterConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func GetSnowDatacenterConfig(fileName string) (*SnowDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
