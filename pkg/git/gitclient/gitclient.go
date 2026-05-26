package gitclient

import (
	"context"
	"time"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport"

	"github.com/aws/eks-anywhere/pkg/retrier"
)

const (
	gitTimeout     = 30 * time.Second
	maxRetries     = 5
	backOffPeriod  = 5 * time.Second
	emptyRepoError = "remote repository is empty"
)

type GitClient struct {
	Auth          transport.AuthMethod
	Client        GoGit
	RepoUrl       string
	RepoDirectory string
	Retrier       *retrier.Retrier
}

type Opt func(*GitClient)

func New(opts ...Opt) *GitClient { _ = "STUB: not implemented"; return nil }

func WithAuth(auth transport.AuthMethod) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithRepositoryUrl(repoUrl string) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithRepositoryDirectory(repoDir string) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func (g *GitClient) Clone(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (g *GitClient) Add(filename string) error { _ = "STUB: not implemented"; return nil }

func (g *GitClient) Remove(filename string) error { _ = "STUB: not implemented"; return nil }

func (g *GitClient) Commit(message string) error { _ = "STUB: not implemented"; return nil }

func (g *GitClient) Push(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (g *GitClient) Pull(ctx context.Context, branch string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GitClient) Init() error { _ = "STUB: not implemented"; return nil }

func (g *GitClient) Branch(name string) error { _ = "STUB: not implemented"; return nil }

func (g *GitClient) ValidateRemoteExists(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if we are able to make a connection to the remote by attempting to list refs

func (g *GitClient) pullIfRemoteExists(r *gogit.Repository, w *gogit.Worktree, branchName string, localBranchRef plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GitClient) remoteBranchExists(r *gogit.Repository, localBranchRef plumbing.ReferenceName) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type GoGit interface {
	AddGlob(f string, w *gogit.Worktree) error
	Checkout(w *gogit.Worktree, opts *gogit.CheckoutOptions) error
	Clone(ctx context.Context, dir string, repoUrl string, auth transport.AuthMethod) (*gogit.Repository, error)
	Commit(m string, sig *object.Signature, w *gogit.Worktree) (plumbing.Hash, error)
	CommitObject(r *gogit.Repository, h plumbing.Hash) (*object.Commit, error)
	Create(r *gogit.Repository, url string) (*gogit.Remote, error)
	CreateBranch(r *gogit.Repository, config *config.Branch) error
	Head(r *gogit.Repository) (*plumbing.Reference, error)
	NewRemote(url, remoteName string) *gogit.Remote
	Init(dir string) (*gogit.Repository, error)
	OpenDir(dir string) (*gogit.Repository, error)
	OpenWorktree(r *gogit.Repository) (*gogit.Worktree, error)
	PushWithContext(ctx context.Context, r *gogit.Repository, auth transport.AuthMethod) error
	PullWithContext(ctx context.Context, w *gogit.Worktree, auth transport.AuthMethod, ref plumbing.ReferenceName) error
	ListRemotes(r *gogit.Repository, auth transport.AuthMethod) ([]*plumbing.Reference, error)
	ListWithContext(ctx context.Context, r *gogit.Remote, auth transport.AuthMethod) ([]*plumbing.Reference, error)
	Remove(f string, w *gogit.Worktree) (plumbing.Hash, error)
	SetRepositoryReference(r *gogit.Repository, p *plumbing.Reference) error
}

type goGit struct{}

func (gg *goGit) Clone(ctx context.Context, dir string, repourl string, auth transport.AuthMethod) (*gogit.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gg *goGit) OpenDir(dir string) (*gogit.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gg *goGit) OpenWorktree(r *gogit.Repository) (*gogit.Worktree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gg *goGit) AddGlob(f string, w *gogit.Worktree) error { _ = "STUB: not implemented"; return nil }

func (gg *goGit) Commit(m string, sig *object.Signature, w *gogit.Worktree) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (gg *goGit) CommitObject(r *gogit.Repository, h plumbing.Hash) (*object.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gg *goGit) PushWithContext(ctx context.Context, r *gogit.Repository, auth transport.AuthMethod) error {
	_ = "STUB: not implemented"
	return nil
}

func (gg *goGit) PullWithContext(ctx context.Context, w *gogit.Worktree, auth transport.AuthMethod, ref plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}

func (gg *goGit) Head(r *gogit.Repository) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gg *goGit) Init(dir string) (*gogit.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ggc *goGit) NewRemote(url, remoteName string) *gogit.Remote {
	_ = "STUB: not implemented"
	return nil
}

func (gg *goGit) Checkout(worktree *gogit.Worktree, opts *gogit.CheckoutOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (gg *goGit) Create(r *gogit.Repository, url string) (*gogit.Remote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gg *goGit) CreateBranch(repo *gogit.Repository, config *config.Branch) error {
	_ = "STUB: not implemented"
	return nil
}

func (gg *goGit) ListRemotes(r *gogit.Repository, auth transport.AuthMethod) ([]*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gg *goGit) Remove(f string, w *gogit.Worktree) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (ggc *goGit) ListWithContext(ctx context.Context, r *gogit.Remote, auth transport.AuthMethod) ([]*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gg *goGit) SetRepositoryReference(r *gogit.Repository, p *plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}
