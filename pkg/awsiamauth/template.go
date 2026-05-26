package awsiamauth

import (
	_ "embed"

	"github.com/google/uuid"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/crypto"
)

//go:embed config/aws-iam-authenticator.yaml
var awsIamAuthTemplate string

//go:embed config/aws-iam-authenticator-ca-secret.yaml
var awsIamAuthCaSecretTemplate string

//go:embed config/aws-iam-authenticator-kubeconfig.yaml
var awsIamAuthKubeconfigTemplate string

// TemplateBuilder generates manifest files from templates.
type TemplateBuilder struct{}

// GenerateManifest generates a YAML Kubernetes manifest for deploying the AWS IAM Authenticator.
func (t *TemplateBuilder) GenerateManifest(clusterSpec *cluster.Spec, clusterID uuid.UUID) ([]byte, error) {
	_ = "STUB: not implemented"
	// Give uuid.Nil semantics that result in no ConfigMap being generated containing the cluster ID
	return nil, nil
}

func (t *TemplateBuilder) mapRolesToYaml(m []v1alpha1.MapRoles) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *TemplateBuilder) mapUsersToYaml(m []v1alpha1.MapUsers) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *TemplateBuilder) setControlPlaneNodeSelector(kubeVersion v1alpha1.KubernetesVersion) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GenerateCertKeyPairSecret generates a YAML Kubernetes Secret for deploying the AWS IAM Authenticator.
func (t *TemplateBuilder) GenerateCertKeyPairSecret(certgen crypto.CertificateGenerator, managementClusterName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateKubeconfig generates a Kubeconfig in yaml format to authenticate with AWS IAM Authenticator.
func (t *TemplateBuilder) GenerateKubeconfig(clusterSpec *cluster.Spec, clusterID uuid.UUID, serverURL, tlsCert string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
