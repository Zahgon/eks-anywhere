package executables

import "context"

type DockerContainer interface {
	Init(ctx context.Context) error
	Close(ctx context.Context) error
	ContainerName() string
}

func NewDockerExecutableBuilder(dockerContainer DockerContainer) *dockerExecutableBuilder {
	_ = "STUB: not implemented"
	return nil
}

type dockerExecutableBuilder struct {
	container DockerContainer
}

func (d *dockerExecutableBuilder) Build(binaryName string) Executable {
	_ = "STUB: not implemented"
	return *new(Executable)
}

func (b *dockerExecutableBuilder) Init(ctx context.Context) (Closer, error) {
	_ = "STUB: not implemented"
	return *new(Closer), nil
}
