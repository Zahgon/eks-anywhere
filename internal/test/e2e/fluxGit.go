package e2e

import (
	"github.com/aws/eks-anywhere/pkg/git"
)

type s3Files struct {
	key, dstPath string
	permission   int
}

type fileFromBytes struct {
	dstPath    string
	permission int
	content    []byte
}

func (f *fileFromBytes) contentString() string { _ = "STUB: not implemented"; return "" }

func (e *E2ESession) setupFluxGitEnv(testRegex string) error { _ = "STUB: not implemented"; return nil }

// add the newly generated repository to the test

func buildFluxGitFiles(envVars map[string]string) []s3Files { _ = "STUB: not implemented"; return nil }

func (e *E2ESession) decodeAndWriteFileToInstance(file fileFromBytes) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *E2ESession) downloadFileInInstance(file s3Files) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *E2ESession) setUpSshAgent(privateKeyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *E2ESession) setupGithubRepo() (*git.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Github API urls get funky if you use ":" in the repo name

// Create a new github repository for the tests to run on

// Add the newly generated public key to the newly created repository as a deploy key

// Newly generated repositories may take some time to show up in the GitHub API; retry a few times to get around this

// Generate a PEM file from the private key and write it instance at the user-provided path

func encodePrivateKey(privateKey []byte) []byte { _ = "STUB: not implemented"; return nil }

func (e *E2ESession) generateKeyPairForGitTest() (privateKeyBytes, publicKeyBytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func gitRepoSshUrl(repo, owner string) string { _ = "STUB: not implemented"; return "" }
