package e2e

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/git"
)

func (e *E2ESession) TestGithubClient(ctx context.Context, githubToken string, owner string, repository string, personal bool) (git.ProviderClient, error) {
	_ = "STUB: not implemented"
	return *new(git.ProviderClient), nil
}
