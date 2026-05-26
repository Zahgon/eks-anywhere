package executables

import (
	"context"
)

const clusterAwsAdminPath = "clusterawsadm"

type Clusterawsadm struct {
	Executable
}

func NewClusterawsadm(executable Executable) *Clusterawsadm { _ = "STUB: not implemented"; return nil }

func (c *Clusterawsadm) BootstrapIam(ctx context.Context, envs map[string]string, configFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Clusterawsadm) BootstrapCreds(ctx context.Context, envs map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Clusterawsadm) ListAccessKeys(ctx context.Context, userName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Clusterawsadm) DeleteCloudformationStack(ctx context.Context, envs map[string]string, fileName string) error {
	_ = "STUB: not implemented"
	return nil
}
