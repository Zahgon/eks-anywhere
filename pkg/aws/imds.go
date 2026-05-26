package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

// IMDSClient is an imds client that wraps around the aws sdk imds client.
type IMDSClient interface {
	GetMetadata(ctx context.Context, params *imds.GetMetadataInput, optFns ...func(*imds.Options)) (*imds.GetMetadataOutput, error)
}

// NewIMDSClient builds a new imds client.
func NewIMDSClient(config aws.Config) *imds.Client { _ = "STUB: not implemented"; return nil }

// BuildIMDS builds or overrides the imds client in the Client with default aws config.
func (c *Client) BuildIMDS(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// EC2InstanceIP calls aws sdk imds.GetMetadata with public-ipv4 path to fetch the instance ip from metadata service.
func (c *Client) EC2InstanceIP(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
