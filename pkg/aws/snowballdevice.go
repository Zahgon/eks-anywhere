package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/eks-anywhere/internal/aws-sdk-go-v2/service/snowballdevice"
)

type SnowballDeviceClient interface {
	DescribeDevice(ctx context.Context, params *snowballdevice.DescribeDeviceInput, optFns ...func(*snowballdevice.Options)) (*snowballdevice.DescribeDeviceOutput, error)
	DescribeDeviceSoftware(ctx context.Context, params *snowballdevice.DescribeDeviceSoftwareInput, optFns ...func(*snowballdevice.Options)) (*snowballdevice.DescribeDeviceSoftwareOutput, error)
}

func NewSnowballClient(config aws.Config) *snowballdevice.Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) IsSnowballDeviceUnlocked(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Client) SnowballDeviceSoftwareVersion(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
