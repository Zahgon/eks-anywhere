package validations

import (
	"context"
)

const (
	recommendedTotalMemory = 6200000000
	requiredMajorVersion   = 20
)

type DockerExecutable interface {
	Version(ctx context.Context) (int, error)
	AllocatedMemory(ctx context.Context) (uint64, error)
}

func CheckMinimumDockerVersion(ctx context.Context, dockerExecutable DockerExecutable) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckDockerAllocatedMemory(ctx context.Context, dockerExecutable DockerExecutable) {
	_ = "STUB: not implemented"
	return
}

func ValidateDockerExecutable(ctx context.Context, docker DockerExecutable, os string) error {
	_ = "STUB: not implemented"
	return nil
}
