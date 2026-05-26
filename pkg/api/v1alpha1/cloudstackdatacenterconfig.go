package v1alpha1

const CloudStackDatacenterKind = "CloudStackDatacenterConfig"

// Used for generating yaml for generate clusterconfig command.
func NewCloudStackDatacenterConfigGenerate(clusterName string) *CloudStackDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *CloudStackDatacenterConfigGenerate) APIVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CloudStackDatacenterConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *CloudStackDatacenterConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func GetCloudStackDatacenterConfig(fileName string) (*CloudStackDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCloudStackManagementAPIEndpointHostname parses the CloudStackAvailabilityZone's ManagementApiEndpoint URL and returns the hostname.
func GetCloudStackManagementAPIEndpointHostname(az CloudStackAvailabilityZone) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getHostnameFromURL(rawurl string) (string, error) { _ = "STUB: not implemented"; return "", nil }
