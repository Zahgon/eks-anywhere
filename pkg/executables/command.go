package executables

import (
	"bytes"
	"context"
)

type commandRunner interface {
	Run(cmd *Command) (stdout bytes.Buffer, err error)
}

type Command struct {
	commandRunner commandRunner
	ctx           context.Context
	args          []string
	stdIn         []byte
	envVars       map[string]string
}

func NewCommand(ctx context.Context, commandRunner commandRunner, args ...string) *Command {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) WithEnvVars(envVars map[string]string) *Command {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) WithStdIn(stdIn []byte) *Command { _ = "STUB: not implemented"; return nil }

func (c *Command) Run() (out bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}
