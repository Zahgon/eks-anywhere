package gogithub

import (
	"context"
	"net/http"

	goGithub "github.com/google/go-github/v35/github"

	"github.com/aws/eks-anywhere/pkg/git"
)

type GoGithub struct {
	Opts   Options
	Client Client
}

type Options struct {
	Auth git.TokenAuth
}

func New(ctx context.Context, opts Options) *GoGithub { _ = "STUB: not implemented"; return nil }

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client interface {
	CreateRepo(ctx context.Context, org string, repo *goGithub.Repository) (*goGithub.Repository, *goGithub.Response, error)
	AddDeployKeyToRepo(ctx context.Context, owner, repo string, key *goGithub.Key) error
	Repo(ctx context.Context, owner, repo string) (*goGithub.Repository, *goGithub.Response, error)
	User(ctx context.Context, user string) (*goGithub.User, *goGithub.Response, error)
	Organization(ctx context.Context, org string) (*goGithub.Organization, *goGithub.Response, error)
	GetContents(ctx context.Context, owner, repo, path string, opt *goGithub.RepositoryContentGetOptions) (
		fileContent *goGithub.RepositoryContent, directoryContent []*goGithub.RepositoryContent, resp *goGithub.Response, err error,
	)
	DeleteRepo(ctx context.Context, owner, repo string) (*goGithub.Response, error)
}

type githubClient struct {
	client *goGithub.Client
}

var HttpClient HTTPClient

func init() {
	HttpClient = &http.Client{}
}

func (ggc *githubClient) CreateRepo(ctx context.Context, org string, repo *goGithub.Repository) (*goGithub.Repository, *goGithub.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (ggc *githubClient) Repo(ctx context.Context, owner, repo string) (*goGithub.Repository, *goGithub.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (ggc *githubClient) User(ctx context.Context, user string) (*goGithub.User, *goGithub.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (ggc *githubClient) Organization(ctx context.Context, org string) (*goGithub.Organization, *goGithub.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (ggc *githubClient) GetContents(ctx context.Context, owner, repo, path string, opt *goGithub.RepositoryContentGetOptions) (fileContent *goGithub.RepositoryContent, directoryContent []*goGithub.RepositoryContent, resp *goGithub.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (ggc *githubClient) DeleteRepo(ctx context.Context, owner, repo string) (*goGithub.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ggc *githubClient) AddDeployKeyToRepo(ctx context.Context, owner, repo string, key *goGithub.Key) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateRepo creates an empty Github Repository. The repository must be initialized locally or
// file must be added to it via the github api before it can be successfully cloned.
func (g *GoGithub) CreateRepo(ctx context.Context, opts git.CreateRepoOpts) (repository *git.Repository, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoGithub) GetAccessTokenPermissions(accessToken string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (g *GoGithub) CheckAccessTokenPermissions(checkPATPermission string, allPermissionScopes string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetRepo describes a remote repository, return the repo name if it exists.
// If the repo does not exist, resulting in a 404 exception, it returns a `RepoDoesNotExist` error.
func (g *GoGithub) GetRepo(ctx context.Context, opts git.GetRepoOpts) (*git.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoGithub) AuthenticatedUser(ctx context.Context) (*goGithub.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// passing the empty string will fetch the authenticated

func (g *GoGithub) Organization(ctx context.Context, org string) (*goGithub.Organization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PathExists checks if a path exists in the remote repository. If the owner, repository or branch doesn't exist,
// it returns false and no error.
func (g *GoGithub) PathExists(ctx context.Context, owner, repo, branch, path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (g *GoGithub) AddDeployKeyToRepo(ctx context.Context, opts git.AddDeployKeyOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteRepo deletes a Github repository.
func (g *GoGithub) DeleteRepo(ctx context.Context, opts git.DeleteRepoOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func newClient(ctx context.Context, opts Options) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

func isNotFound(err error) bool { _ = "STUB: not implemented"; return false }
