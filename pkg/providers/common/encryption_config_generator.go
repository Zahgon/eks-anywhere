package common

import (
	apiserverv1 "k8s.io/apiserver/pkg/apis/apiserver/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	encryptionConfigurationKind  = "EncryptionConfiguration"
	encryptionProviderVersion    = "v1"
	encryptionProviderNamePrefix = "aws-encryption-provider"
)

var identityProvider = apiserverv1.ProviderConfiguration{
	Identity: &apiserverv1.IdentityConfiguration{},
}

// GenerateKMSEncryptionConfiguration takes a list of the EtcdEncryption configs and generates the corresponding Kubernetes Encryptionapiserverv1.
func GenerateKMSEncryptionConfiguration(confs *[]v1alpha1.EtcdEncryption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
