package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/aws/eks-anywhere/pkg/utils/ptr"
)

var (
	// DefaultKMSCacheSize is the default cache size for KMS provider (1000).
	DefaultKMSCacheSize = ptr.Int32(1000)
	// DefaultKMSTimeout is the default timeout for KMS provider (3s).
	DefaultKMSTimeout = metav1.Duration{Duration: time.Second * 3}
)

// ValidateEtcdEncryptionConfig validates the etcd encryption configuration.
func ValidateEtcdEncryptionConfig(config *[]EtcdEncryption) error {
	_ = "STUB: not implemented"
	return nil
}

func validateKMSConfig(kms *KMS) error { _ = "STUB: not implemented"; return nil }

func setEtcdEncryptionConfigDefaults(cluster *Cluster) error { _ = "STUB: not implemented"; return nil }

func setKMSConfigDefauts(kms *KMS) { _ = "STUB: not implemented"; return }
