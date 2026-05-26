package executables

import "context"

type localExecutableBuilder struct{}

func newLocalExecutableBuilder() localExecutableBuilder {
	_ = "STUB: not implemented"
	return *new(localExecutableBuilder)
}

func (b localExecutableBuilder) Build(binaryPath string) Executable {
	_ = "STUB: not implemented"
	return *new(Executable)
}

func (b localExecutableBuilder) Init(_ context.Context) (Closer, error) {
	_ = "STUB: not implemented"
	return *new(Closer), nil
}

func NoOpClose(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
