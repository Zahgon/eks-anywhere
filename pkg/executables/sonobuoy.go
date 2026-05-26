package executables

import (
	"context"
)

const sonobuoyPath = "./sonobuoy"

type Sonobuoy struct {
	Executable
}

func NewSonobuoy(executable Executable) *Sonobuoy { _ = "STUB: not implemented"; return nil }

func (k *Sonobuoy) Run(ctx context.Context, contextName string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *Sonobuoy) GetResults(ctx context.Context, contextName string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
