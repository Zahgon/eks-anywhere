package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Client provides the single API client to make operations call to aws services.
type Client struct {
	ec2            EC2Client
	imds           IMDSClient
	snowballDevice SnowballDeviceClient
}

// Clients are a map between aws profile and its aws client.
type Clients map[string]*Client

type ServiceEndpoint struct {
	ServiceID     string
	URL           string
	SigningRegion string
}

type AwsConfigOpt = config.LoadOptionsFunc

func AwsConfigOptSet(opts ...AwsConfigOpt) AwsConfigOpt {
	_ = "STUB: not implemented"
	return *new(AwsConfigOpt)
}

// LoadConfig reads the optional aws configurations, and populates an AWS Config
// with the values from the configurations.
func LoadConfig(ctx context.Context, opts ...AwsConfigOpt) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

// ClientOpt updates an aws.Client.
type ClientOpt func(*Client)

// WithEC2 returns a ClientOpt that sets the ec2 client.
func WithEC2(ec2 EC2Client) ClientOpt { _ = "STUB: not implemented"; return *new(ClientOpt) }

// WithIMDS returns a ClientOpt that sets the imds client.
func WithIMDS(imds IMDSClient) ClientOpt { _ = "STUB: not implemented"; return *new(ClientOpt) }

// WithSnowballDevice returns a ClientOpt that sets the snowballdevice client.
func WithSnowballDevice(snowballdevice SnowballDeviceClient) ClientOpt {
	_ = "STUB: not implemented"
	return *new(ClientOpt)
}

// NewClient builds an aws Client.
func NewClient(opts ...ClientOpt) *Client { _ = "STUB: not implemented"; return nil }

// NewClientFromConfig builds an aws client with ec2 and snowballdevice apis from aws config.
func NewClientFromConfig(cfg aws.Config) *Client { _ = "STUB: not implemented"; return nil }
