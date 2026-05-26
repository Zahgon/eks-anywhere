package curatedpackages

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// RegistryAccessTestParams struct for testing resigtry access.
type RegistryAccessTestParams struct {
	AccessKey    string
	Secret       string
	SessionToken string
	Region       string
	AwsConfig    string
	Registry     string
}

// RegistryAccessTester test if AWS credentials has valid permission to access an ECR registry.
type RegistryAccessTester interface {
	Test(ctx context.Context, params RegistryAccessTestParams) error
}

// DefaultRegistryAccessTester the default implementation of RegistryAccessTester.
type DefaultRegistryAccessTester struct{}

// Test if the AWS static credential or sharedConfig has valid permission to access an ECR registry.
func (r *DefaultRegistryAccessTester) Test(ctx context.Context, params RegistryAccessTestParams) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TestRegistryWithAuthToken test if the registry can be acccessed with auth token.
func TestRegistryWithAuthToken(authToken, registry string, do Do) error {
	_ = "STUB: not implemented"
	return nil
}

// 404 means the IAM policy is good, so 404 is good here

// GetRegionalRegistry get the regional registry corresponding to defaultRegistry in a specific region.
func GetRegionalRegistry(defaultRegistry, region string) string {
	_ = "STUB: not implemented"
	return ""
}

// RegistryAuthTokenProvider provides auth token for registry access.
type RegistryAuthTokenProvider interface {
	GetTokenByAWSConfig(ctx context.Context, awsConfig string) (string, error)
	GetTokenByAWSKeySecret(ctx context.Context, key, secret, region string) (string, error)
}

// DefaultRegistryAuthTokenProvider provides auth token for AWS ECR registry access.
type DefaultRegistryAuthTokenProvider struct{}

// GetTokenByAWSConfig get auth token by AWS config.
func (d *DefaultRegistryAuthTokenProvider) GetTokenByAWSConfig(ctx context.Context, awsConfig string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseAWSConfig parse AWS config from string.
func ParseAWSConfig(ctx context.Context, awsConfig string) (*aws.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAWSConfigFromKeySecret get AWS config from key, secret and region.
func GetAWSConfigFromKeySecret(ctx context.Context, key, secret, sessionToken, region string) (*aws.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTokenByAWSKeySecret get auth token by AWS key and secret.
func (d *DefaultRegistryAuthTokenProvider) GetTokenByAWSKeySecret(ctx context.Context, key, secret, sessionToken, region string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getAuthorizationToken(cfg aws.Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Do is a function type that takes a http request and returns a http response.
type Do func(req *http.Request) (*http.Response, error)

// TestRegistryAccessWithAWSConfig test if the AWS config has valid permission to access container registry.
func TestRegistryAccessWithAWSConfig(ctx context.Context, awsConfig, registry string, tokenProvider RegistryAuthTokenProvider, do Do) error {
	_ = "STUB: not implemented"
	return nil
}

// TestRegistryAccessWithAWSKeySecret test if the AWS key and secret has valid permission to access container registry.
func TestRegistryAccessWithAWSKeySecret(ctx context.Context, key, secret, region, registry string, tokenProvider RegistryAuthTokenProvider, do Do) error {
	_ = "STUB: not implemented"
	return nil
}
