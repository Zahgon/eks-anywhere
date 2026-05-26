package executables

import (
	"bytes"
	"context"
	"sync"

	"github.com/aws/eks-anywhere/pkg/retrier"
)

type DockerClient interface {
	Login(ctx context.Context, endpoint, username, password string) error
	PullImage(ctx context.Context, image string) error
	Execute(ctx context.Context, args ...string) (stdout bytes.Buffer, err error)
}

type dockerContainer struct {
	image               string
	workingDir          string
	mountDirs           []string
	containerName       string
	dockerClient        DockerClient
	initOnce, closeOnce sync.Once
	*retrier.Retrier
}

func newDockerContainer(image, workingDir string, mountDirs []string, dockerClient DockerClient) *dockerContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewDockerContainerCustomBinary(docker DockerClient) *dockerContainer {
	_ = "STUB: not implemented"
	return nil
}

func (d *dockerContainer) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// start container and keep it running in the background

func (d *dockerContainer) ContainerName() string { _ = "STUB: not implemented"; return "" }

func (d *dockerContainer) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
