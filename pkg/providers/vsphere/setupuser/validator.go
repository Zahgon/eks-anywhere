package setupuser

import (
	"context"
)

const (
	DefaultUsername   = "eksa"
	DefaultGroup      = "EKSAUsers"
	DefaultGlobalRole = "EKSAGlobalRole"
	DefaultUserRole   = "EKSAUserRole"
	DefaultAdminRole  = "EKSACloudAdminRole"
)

type Connection struct {
	Server   string `yaml:"server"`
	Insecure bool   `yaml:"insecure"`
}

type Objects struct {
	Networks      []string `yaml:"networks"`
	Datastores    []string `yaml:"datastores"`
	ResourcePools []string `yaml:"resourcePools"`
	Folders       []string `yaml:"folders"`
	Templates     []string `yaml:"templates"`
}

type VSphereUserSpec struct {
	Datacenter    string     `yaml:"datacenter"`
	VSphereDomain string     `yaml:"vSphereDomain"`
	Connection    Connection `yaml:"connection"`
	Objects       Objects    `yaml:"objects"`
	// Below are optional fields with defaults
	Username   string `yaml:"username"`
	GroupName  string `yaml:"group"`
	GlobalRole string `yaml:"globalRole"`
	UserRole   string `yaml:"userRole"`
	AdminRole  string `yaml:"adminRole"`
}

type VSphereSetupUserConfig struct {
	ApiVersion string          `yaml:"apiVersion"`
	Kind       string          `yaml:"kind"`
	Spec       VSphereUserSpec `yaml:"spec"`
}

func GenerateConfig(ctx context.Context, filepath string) (*VSphereSetupUserConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readConfig(ctx context.Context, filepath string) (*VSphereSetupUserConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validate(c *VSphereSetupUserConfig) error { _ = "STUB: not implemented"; return nil }

func setDefaults(c *VSphereSetupUserConfig) { _ = "STUB: not implemented"; return }

// ValidateVSphereObjects validates objects do not exist before configuring user.
func ValidateVSphereObjects(ctx context.Context, c *VSphereSetupUserConfig, govc GovcClient) error {
	_ = "STUB: not implemented"
	return nil
}
