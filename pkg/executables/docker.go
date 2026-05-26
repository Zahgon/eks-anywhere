package executables

import (
	"context"
)

// Temporary: Curated packages dev and prod accounts are currently hard coded
// This is because there is no mechanism to extract these values as of now.
const (
	dockerPath        = "docker"
	defaultRegistry   = "public.ecr.aws"
	packageProdDomain = "783794618700.dkr.ecr.us-west-2.amazonaws.com"
	packageDevDomain  = "067575901363.dkr.ecr.us-west-2.amazonaws.com"
)

type Docker struct {
	Executable
}

func NewDocker(executable Executable) *Docker { _ = "STUB: not implemented"; return nil }

func (d *Docker) GetDockerLBPort(ctx context.Context, clusterName string) (port string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *Docker) PullImage(ctx context.Context, image string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Docker) Version(ctx context.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Docker) AllocatedMemory(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Docker) TagImage(ctx context.Context, image string, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Docker) PushImage(ctx context.Context, image string, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Docker) Login(ctx context.Context, endpoint, username, password string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Docker) LoadFromFile(ctx context.Context, filepath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Docker) SaveToFile(ctx context.Context, filepath string, images ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Docker) Run(ctx context.Context, image string, name string, cmd []string, flags ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Docker) ForceRemove(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckContainerExistence checks whether a Docker container with the provided name exists
// It returns true if a container with the name exists, false if it doesn't and an error if it encounters some other error.
func (d *Docker) CheckContainerExistence(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
