package cmd

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/dependencies"
	"github.com/aws/eks-anywhere/pkg/types"
)

func cleanup(deps *dependencies.Dependencies, commandErr *error) { _ = "STUB: not implemented"; return }

func close(ctx context.Context, closer types.Closer) { _ = "STUB: not implemented"; return }

func cleanupDirectory(directory string) { _ = "STUB: not implemented"; return }
