package gitfactory

import (
	"context"

	"github.com/go-git/go-git/v5/plumbing/transport"
	gogitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/git"
	"github.com/aws/eks-anywhere/pkg/git/gitclient"
)

type GitTools struct {
	Provider            git.ProviderClient
	Client              git.Client
	Writer              filewriter.FileWriter
	RepositoryDirectory string
}

type GitToolsOpt func(opts *GitTools)

func Build(ctx context.Context, cluster *v1alpha1.Cluster, fluxConfig *v1alpha1.FluxConfig, writer filewriter.FileWriter, opts ...GitToolsOpt) (*GitTools, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildGitClient(ctx context.Context, auth transport.AuthMethod, repoUrl string, repo string) *gitclient.GitClient {
	_ = "STUB: not implemented"
	return nil
}

func buildGithubProvider(ctx context.Context, githubToken string, config *v1alpha1.GithubProviderConfig) (git.ProviderClient, error) {
	_ = "STUB: not implemented"
	return *new(git.ProviderClient), nil
}

func newRepositoryWriter(writer filewriter.FileWriter, repository string) (filewriter.FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(filewriter.FileWriter), nil
}

func WithRepositoryDirectory(repoDir string) GitToolsOpt {
	_ = "STUB: not implemented"
	return *new(GitToolsOpt)
}

func getSSHAuthFromPrivateKey(privateKeyFile string, passphrase string, user string) (gogitssh.AuthMethod, error) {
	_ = "STUB: not implemented"
	return *new(gogitssh.AuthMethod), nil
}

func getSignerFromPrivateKeyFile(privateKeyFile string, passphrase string) (ssh.Signer, error) {
	_ = "STUB: not implemented"
	return *new(ssh.Signer), nil
}
