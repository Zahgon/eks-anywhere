package flux

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/git"
	gitFactory "github.com/aws/eks-anywhere/pkg/git/factory"
	"github.com/aws/eks-anywhere/pkg/retrier"
)

type gitClient struct {
	git         git.Client
	gitProvider git.ProviderClient
	*retrier.Retrier
}

func newGitClient(gitTools *gitFactory.GitTools) *gitClient { _ = "STUB: not implemented"; return nil }

func (c *gitClient) GetRepo(ctx context.Context) (repo *git.Repository, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *gitClient) CreateRepo(ctx context.Context, opts git.CreateRepoOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *gitClient) Clone(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *gitClient) Push(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *gitClient) Pull(ctx context.Context, branch string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *gitClient) PathExists(ctx context.Context, owner, repo, branch, path string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *gitClient) Add(filename string) error { _ = "STUB: not implemented"; return nil }

func (c *gitClient) Remove(filename string) error { _ = "STUB: not implemented"; return nil }

func (c *gitClient) Commit(message string) error { _ = "STUB: not implemented"; return nil }

func (c *gitClient) Branch(name string) error { _ = "STUB: not implemented"; return nil }

func (c *gitClient) Init() error { _ = "STUB: not implemented"; return nil }
