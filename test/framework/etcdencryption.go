package framework

import (
	"context"
	_ "embed"

	"github.com/aws/aws-sdk-go/aws/session"
	jose "github.com/go-jose/go-jose/v3"
)

const (
	irsaS3BucketVar            = "T_IRSA_S3_BUCKET"
	kmsIAMRoleVar              = "T_KMS_IAM_ROLE"
	kmsImageVar                = "T_KMS_IMAGE"
	podIdentityWebhookImageVar = "T_POD_IDENTITY_WEBHOOK_IMAGE"
	kmsKeyArn                  = "T_KMS_KEY_ARN"
	kmsKeyRegion               = "T_KMS_KEY_REGION"
	kmsSocketVar               = "T_KMS_SOCKET"

	defaultRegion       = "us-west-2"
	keysFilename        = "keys.json"
	keyIDFilenameFormat = "%s-oidc-keyid"

	// SSHKeyPath is the path where the SSH private key is stored on the test-runner instance.
	SSHKeyPath = "/tmp/ssh_key"
)

//go:embed config/pod-identity-webhook.yaml
var podIdentityWebhookManifest string

//go:embed config/aws-kms-encryption-provider.yaml
var kmsProviderManifest string

type keyResponse struct {
	Keys []jose.JSONWebKey `json:"keys"`
}

// etcdEncryptionTestVars stores all the environment variables needed by etcd encryption tests.
type etcdEncryptionTestVars struct {
	KmsKeyRegion            string
	S3Bucket                string
	KmsIamRole              string
	KmsImage                string
	PodIdentityWebhookImage string
	KmsKeyArn               string
	KmsSocket               string
}

// RequiredEtcdEncryptionEnvVars returns the environment variables required .
func RequiredEtcdEncryptionEnvVars() []string { _ = "STUB: not implemented"; return nil }

func getEtcdEncryptionVarsFromEnv() *etcdEncryptionTestVars { _ = "STUB: not implemented"; return nil }

// WithPodIamConfig is a ClusterE2ETestOpt that adds pod IAM config to the cluster.
func WithPodIamConfig() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithEtcdEncrytion is a ClusterE2ETestOpt that adds etcd encryption config to the Cluster.
func WithEtcdEncrytion() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// ValidateEtcdEncryption validates that etcd encryption is properly configured by creating a secret
// and SSHing into the ETCD nodes and ensuring the secret is not stored in plaintext.
func (e *ClusterE2ETest) ValidateEtcdEncryption() { _ = "STUB: not implemented"; return }

func getSSHUsernameByProvider(provider string) string { _ = "STUB: not implemented"; return "" }

// PostClusterCreateEtcdEncryptionSetup performs operations on the cluster to prepare it for etcd encryption.
// These operations include:
// - Adding Cluster SA cert to the OIDC provider's keys.
// - Deploying Pod Identity Webhook.
// - Deploying AWS KMS Provider.
func (e *ClusterE2ETest) PostClusterCreateEtcdEncryptionSetup() { _ = "STUB: not implemented"; return }

// register cleanup step to remove the keys from s3 after the test is done

// cleanup removes the cluster's key from the IAM OIDC config.
func (e *ClusterE2ETest) cleanupKeysFromOIDCConfig() { _ = "STUB: not implemented"; return }

// download the current keys json from S3 to add the current cluster's cert

// upload the modified keys json to s3 with the public read access

func getIssuerURL() string { _ = "STUB: not implemented"; return "" }

func (e *ClusterE2ETest) deployPodIdentityWebhook(ctx context.Context, envVars *etcdEncryptionTestVars) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) deployKMSProvider(ctx context.Context, envVars *etcdEncryptionTestVars) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) addClusterCertToIrsaOidcProvider(ctx context.Context, envVars *etcdEncryptionTestVars, awsSession *session.Session) error {
	_ = "STUB: not implemented"
	return nil
}

// Fetch the cluster's service account cert

// download the current keys json from S3 to add the current cluster's cert

// upload the modified keys json to s3 with the public read access

func (e *ClusterE2ETest) getClusterSACert(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getJSONWebKeyFromCertFile(cert []byte) (*jose.JSONWebKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func keyIDFromPublicKey(publicKey interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
