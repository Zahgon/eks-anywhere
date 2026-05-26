package setupuser

import (
	"context"
)

const (
	vSphereRootPath = "/"
)

// GovcClient specifies govc functions required to configure a vsphere user.
type GovcClient interface {
	CreateUser(ctx context.Context, username string, password string) error
	UserExists(ctx context.Context, username string) (bool, error)
	CreateGroup(ctx context.Context, name string) error
	GroupExists(ctx context.Context, name string) (bool, error)
	AddUserToGroup(ctx context.Context, name string, username string) error
	RoleExists(ctx context.Context, name string) (bool, error)
	CreateRole(ctx context.Context, name string, privileges []string) error
	SetGroupRoleOnObject(ctx context.Context, principal string, role string, object string, domain string) error
}

// SetupGOVCEnv creates appropriate govc environment variables to build govc client.
func SetupGOVCEnv(ctx context.Context, vsuc *VSphereSetupUserConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Run sets up a vSphere user with appropriate group, role, and permissions to create EKS-A kubernetes clusters.
func Run(ctx context.Context, vsuc *VSphereSetupUserConfig, govc GovcClient) error {
	_ = "STUB: not implemented"
	return nil
}

func createGroup(ctx context.Context, vsuc *VSphereSetupUserConfig, govc GovcClient) error {
	_ = "STUB: not implemented"
	return nil
}

func createRoles(ctx context.Context, vsuc *VSphereSetupUserConfig, govc GovcClient) error {
	_ = "STUB: not implemented"
	return nil
}

func associateRolesToObjects(ctx context.Context, vsuc *VSphereSetupUserConfig, govc GovcClient) error {
	_ = "STUB: not implemented"
	return nil
}

func addUserToGroup(ctx context.Context, vsuc *VSphereSetupUserConfig, govc GovcClient) error {
	_ = "STUB: not implemented"
	// associate user to group
	return nil
}

func setGroupRoleOnObjects(ctx context.Context, vsuc *VSphereSetupUserConfig, govc GovcClient, role string, objects []string) error {
	_ = "STUB: not implemented"
	return nil
}

func getPrivsFromFile(privsContent string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type vsphereRole struct {
	name  string
	privs []string
}

func getRoles(vsuc *VSphereSetupUserConfig) ([]vsphereRole, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getUserRoleObjects(vsuc *VSphereSetupUserConfig) []string {
	_ = "STUB: not implemented"
	return nil
}
