package api

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type AWSIamConfigOpt func(c *v1alpha1.AWSIamConfig)

func NewAWSIamConfig(name string, opts ...AWSIamConfigOpt) *v1alpha1.AWSIamConfig {
	_ = "STUB: not implemented"
	return nil
}

func WithAWSIamAWSRegion(awsRegion string) AWSIamConfigOpt {
	_ = "STUB: not implemented"
	return *new(AWSIamConfigOpt)
}

func WithAWSIamBackendMode(backendMode string) AWSIamConfigOpt {
	_ = "STUB: not implemented"
	return *new(AWSIamConfigOpt)
}

func AddAWSIamRole(arn, username string, groups []string) *v1alpha1.MapRoles {
	_ = "STUB: not implemented"
	return nil
}

func WithAWSIamMapRoles(mapRoles *v1alpha1.MapRoles) AWSIamConfigOpt {
	_ = "STUB: not implemented"
	return *new(AWSIamConfigOpt)
}

func AddAWSIamUser(arn, username string, groups []string) *v1alpha1.MapUsers {
	_ = "STUB: not implemented"
	return nil
}

func WithAWSIamMapUsers(mapUsers *v1alpha1.MapUsers) AWSIamConfigOpt {
	_ = "STUB: not implemented"
	return *new(AWSIamConfigOpt)
}

func WithAWSIamPartition(partition string) AWSIamConfigOpt {
	_ = "STUB: not implemented"
	return *new(AWSIamConfigOpt)
}
