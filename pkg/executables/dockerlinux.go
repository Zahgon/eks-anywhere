package executables

import (
	"bytes"
	"context"
)

const containerNamePrefix = "eksa_"

type linuxDockerExecutable struct {
	cli           string
	containerName string
}

// This currently returns a linuxDockerExecutable, but if we support other types of docker executables we can change
// the name of this constructor.
func NewDockerExecutable(cli string, containerName string) Executable {
	_ = "STUB: not implemented"
	return *new(Executable)
}

func (e *linuxDockerExecutable) Execute(ctx context.Context, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (e *linuxDockerExecutable) ExecuteWithStdin(ctx context.Context, in []byte, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (e *linuxDockerExecutable) ExecuteWithEnv(ctx context.Context, envs map[string]string, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (e *linuxDockerExecutable) Command(ctx context.Context, args ...string) *Command {
	_ = "STUB: not implemented"
	return nil
}

func (e *linuxDockerExecutable) Run(cmd *Command) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (e *linuxDockerExecutable) buildCommand(envs map[string]string, cli string, args ...string) []string {
	_ = "STUB: not implemented"
	return nil
}
