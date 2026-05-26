package v1alpha1

const (
	AWSIamConfigKind = "AWSIamConfig"
	eksConfigMap     = "EKSConfigMap"
	mountedFile      = "MountedFile"

	DefaultAWSIamConfigPartition = "aws"
)

func GetAndValidateAWSIamConfig(fileName string, refName string, clusterConfig *Cluster) (*AWSIamConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getAWSIamConfig(fileName string) (*AWSIamConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the name is empty, we can assume that they didn't configure their AWS IAM configuration, so return nil

func validateAWSIamConfig(config *AWSIamConfig) error { _ = "STUB: not implemented"; return nil }

func validateMapRoles(mapRoles []MapRoles) error { _ = "STUB: not implemented"; return nil }

func validateMapUsers(mapUsers []MapUsers) error { _ = "STUB: not implemented"; return nil }

func validateAWSIamRefName(config *AWSIamConfig, refName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAWSIamNamespace(config *AWSIamConfig, clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func setDefaultAWSIamPartition(config *AWSIamConfig) { _ = "STUB: not implemented"; return }
