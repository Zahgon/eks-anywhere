package nutanix

import (
	"github.com/nutanix-cloud-native/prism-go-client/environment/credentials"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// ClientCache is a map of NutanixDatacenterConfig name to Nutanix client.
type ClientCache struct {
	clients map[string]Client
}

// NewClientCache returns a new ClientCache.
func NewClientCache() *ClientCache { _ = "STUB: not implemented"; return nil }

// GetNutanixClient returns a Nutanix client for the given NutanixDatacenterConfig.
func (cb *ClientCache) GetNutanixClient(datacenterConfig *anywherev1.NutanixDatacenterConfig, creds credentials.BasicAuthCredential) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}
