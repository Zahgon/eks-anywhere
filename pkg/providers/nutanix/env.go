package nutanix

import (
	"os"

	"github.com/nutanix-cloud-native/prism-go-client/environment/credentials"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	nutanixEndpointKey       = "NUTANIX_ENDPOINT"
	expClusterResourceSetKey = "EXP_CLUSTER_RESOURCE_SET"
)

var osSetenv = os.Setenv

func setupEnvVars(datacenterConfig *anywherev1.NutanixDatacenterConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCredsFromEnv returns nutanix credentials based on the environment.
func GetCredsFromEnv() credentials.BasicAuthCredential {
	_ = "STUB: not implemented"
	return *new(credentials.BasicAuthCredential)
}
