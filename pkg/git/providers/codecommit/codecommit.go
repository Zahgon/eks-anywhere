package codecommit

const (
	// DefaultSSHAuthUser for client auth.
	DefaultSSHAuthUser = "git"
	codeCommitSubHost  = "git-codecommit"
	awsSubHost         = "amazonaws.com"
)

// IsCodeCommitURL check if repo url is code commit url and returns user from url.
func IsCodeCommitURL(repoURL string) (string, error) { _ = "STUB: not implemented"; return "", nil }
