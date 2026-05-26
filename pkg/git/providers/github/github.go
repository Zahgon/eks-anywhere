package github

import (
	"context"

	goGithub "github.com/google/go-github/v35/github"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/git"
)

const (
	GitProviderName    = "github"
	EksaGithubTokenEnv = "EKSA_GITHUB_TOKEN"
	GithubTokenEnv     = "GITHUB_TOKEN"
	githubUrlTemplate  = "https://github.com/%v/%v.git"
	patRegex           = "^ghp_[a-zA-Z0-9]{36}|github_pat_[a-zA-Z0-9]{22}_[a-zA-Z0-9]{59}$"
	repoPermissions    = "repo"
)

type githubProvider struct {
	githubProviderClient GithubClient
	config               *v1alpha1.GithubProviderConfig
	auth                 git.TokenAuth
}

type Options struct {
	Repository string
	Owner      string
	Personal   bool
}

// GithubClient represents the attributes that the Github provider requires of a library to directly connect to and interact with the Github API.
type GithubClient interface {
	GetRepo(ctx context.Context, opts git.GetRepoOpts) (repo *git.Repository, err error)
	CreateRepo(ctx context.Context, opts git.CreateRepoOpts) (repo *git.Repository, err error)
	AddDeployKeyToRepo(ctx context.Context, opts git.AddDeployKeyOpts) error
	AuthenticatedUser(ctx context.Context) (*goGithub.User, error)
	Organization(ctx context.Context, org string) (*goGithub.Organization, error)
	GetAccessTokenPermissions(accessToken string) (string, error)
	CheckAccessTokenPermissions(checkPATPermission string, allPermissionScopes string) error
	PathExists(ctx context.Context, owner, repo, branch, path string) (bool, error)
	DeleteRepo(ctx context.Context, opts git.DeleteRepoOpts) error
}

func New(githubProviderClient GithubClient, config *v1alpha1.GithubProviderConfig, auth git.TokenAuth) (*githubProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRepo creates an empty Github Repository. The repository must be initialized locally or
// file must be added to it via the github api before it can be successfully cloned.
func (g *githubProvider) CreateRepo(ctx context.Context, opts git.CreateRepoOpts) (repository *git.Repository, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRepo describes a remote repository, return the repo name if it exists.
// If the repo does not exist, a nil repo is returned.
func (g *githubProvider) GetRepo(ctx context.Context) (*git.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *githubProvider) AddDeployKeyToRepo(ctx context.Context, opts git.AddDeployKeyOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// validates the github setup and access.
func (g *githubProvider) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// for now only checks if user belongs to the org

func validateGithubAccessToken() error { _ = "STUB: not implemented"; return nil }

func GetGithubAccessTokenFromEnv() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (g *githubProvider) PathExists(ctx context.Context, owner, repo, branch, path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (g *githubProvider) DeleteRepo(ctx context.Context, opts git.DeleteRepoOpts) error {
	_ = "STUB: not implemented"
	return nil
}

type GitProviderNotFoundError struct {
	Provider string
}

func (e *GitProviderNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func RepoUrl(owner string, repo string) string { _ = "STUB: not implemented"; return "" }
