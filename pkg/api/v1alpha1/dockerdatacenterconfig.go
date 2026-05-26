package v1alpha1

const DockerDatacenterKind = "DockerDatacenterConfig"

// Used for generating yaml for generate clusterconfig command.
func NewDockerDatacenterConfigGenerate(clusterName string) *DockerDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *DockerDatacenterConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *DockerDatacenterConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *DockerDatacenterConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func GetDockerDatacenterConfig(fileName string) (*DockerDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
