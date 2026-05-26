package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// EC2Client is an ec2 client that wraps around the aws sdk ec2 client.
type EC2Client interface {
	DescribeImages(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	DescribeKeyPairs(ctx context.Context, params *ec2.DescribeKeyPairsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error)
	DescribeInstanceTypes(ctx context.Context, params *ec2.DescribeInstanceTypesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error)
	ImportKeyPair(ctx context.Context, params *ec2.ImportKeyPairInput, optFns ...func(*ec2.Options)) (*ec2.ImportKeyPairOutput, error)
}

// NewEC2Client builds a new ec2 client.
func NewEC2Client(config aws.Config) *ec2.Client { _ = "STUB: not implemented"; return nil }

// EC2ImageExists calls aws sdk ec2.DescribeImages with filter imageID to fetch a
// specified images (AMIs, AKIs, and ARIs) available.
// Returns (false, nil) if the image does not exist, (true, nil) if image exists,
// and (false, err) if there is an non 400 status code error from ec2.DescribeImages.
func (c *Client) EC2ImageExists(ctx context.Context, imageID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// EC2KeyNameExists calls aws sdk ec2.DescribeKeyPairs with filter keyName to fetch a
// specified key pair available in aws.
// Returns (false, nil) if the key pair does not exist, (true, nil) if key pair exists,
// and (false, err) if there is an error from ec2.DescribeKeyPairs.
func (c *Client) EC2KeyNameExists(ctx context.Context, keyName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// EC2ImportKeyPair calls aws sdk ec2.ImportKeyPair to import a key pair to ec2.
func (c *Client) EC2ImportKeyPair(ctx context.Context, keyName string, keyMaterial []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// EC2InstanceType has the information of an ec2 instance type.
type EC2InstanceType struct {
	Name        string
	DefaultVCPU *int32
}

// EC2InstanceTypes calls aws sdk ec2.DescribeInstanceTypes to get a list of supported instance type for a device.
func (c *Client) EC2InstanceTypes(ctx context.Context) ([]EC2InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
