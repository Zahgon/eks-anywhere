package reconciler

import (
	"context"

	"gopkg.in/ini.v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/aws"
	"github.com/aws/eks-anywhere/pkg/providers/snow"
)

type AwsClientBuilder struct {
	client client.Client
}

func NewAwsClientBuilder(client client.Client) *AwsClientBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AwsClientBuilder) Get(ctx context.Context) (snow.AwsClientMap, error) {
	_ = "STUB: not implemented"
	// Setting the aws client map in validator on every reconcile based on the secrets at that point of time
	return *new(snow.AwsClientMap), nil
}

type credentialConfiguration struct {
	AccessKey string `ini:"aws_access_key_id"`
	SecretKey string `ini:"aws_secret_access_key"`
	Region    string `ini:"region"`
}

func createAwsClients(ctx context.Context, credentials []byte, certificates []byte) (aws.Clients, error) {
	_ = "STUB: not implemented"
	return *new(aws.Clients), nil
}

func parseIpConfiguration(credsCfg *ini.File, ip string) (*credentialConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
