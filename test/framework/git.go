package framework

import (
	"context"
	_ "embed"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	gitFactory "github.com/aws/eks-anywhere/pkg/git/factory"
)

func (e *ClusterE2ETest) NewGitTools(ctx context.Context, cluster *v1alpha1.Cluster, fluxConfig *v1alpha1.FluxConfig, writer filewriter.FileWriter, repoPath string) (*gitFactory.GitTools, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
