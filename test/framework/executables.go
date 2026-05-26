package framework

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/helm"
)

func buildKubectl(t T) *executables.Kubectl { _ = "STUB: not implemented"; return nil }

func buildLocalKubectl() *executables.Kubectl { _ = "STUB: not implemented"; return nil }

func executableBuilder(ctx context.Context, t T) *executables.ExecutablesBuilder {
	_ = "STUB: not implemented"
	return nil
}

func buildGovc(t T) *executables.Govc { _ = "STUB: not implemented"; return nil }

func buildDocker(t T) *executables.Docker { _ = "STUB: not implemented"; return nil }

func buildHelm(t T) helm.Client { _ = "STUB: not implemented"; return *new(helm.Client) }

func buildSSH(t T) *executables.SSH { _ = "STUB: not implemented"; return nil }

func buildCmk(t T) *executables.Cmk { _ = "STUB: not implemented"; return nil }
