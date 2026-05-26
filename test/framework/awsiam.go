package framework

import (
	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	"github.com/aws/eks-anywhere/internal/pkg/api"
)

const (
	AWSIamRoleArn = "T_AWS_IAM_ROLE_ARN"
)

var awsIamRequiredEnvVars = []string{
	AWSIamRoleArn,
}

func RequiredAWSIamEnvVars() []string { _ = "STUB: not implemented"; return nil }

func WithAWSIam() ClusterE2ETestOpt { _ = "STUB: not implemented"; return *new(ClusterE2ETestOpt) }

func withArnFromEnv(envVar string) string { _ = "STUB: not implemented"; return "" }

func (e *ClusterE2ETest) ValidateAWSIamAuth() { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) downloadAwsIamAuthClient() error { _ = "STUB: not implemented"; return nil }

func (e *ClusterE2ETest) setIamAuthClientPATH() error { _ = "STUB: not implemented"; return nil }

func (e *ClusterE2ETest) getEksdReleaseManifest() (*eksdv1alpha1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ClusterE2ETest) iamAuthKubeconfigFilePath() string { _ = "STUB: not implemented"; return "" }

// WithAwsIamEnvVarCheck returns a ClusterE2ETestOpt that checks for the required env vars.
func WithAwsIamEnvVarCheck() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithAwsIamConfig sets aws iam in cluster config.
func WithAwsIamConfig() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}
