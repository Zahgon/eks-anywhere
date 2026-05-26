package executables

import (
	"context"
)

const awsCliPath = "aws"

type AwsCli struct {
	Executable
}

func NewAwsCli(executable Executable) *AwsCli { _ = "STUB: not implemented"; return nil }

func (ac *AwsCli) CreateAccessKey(ctx context.Context, username string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
