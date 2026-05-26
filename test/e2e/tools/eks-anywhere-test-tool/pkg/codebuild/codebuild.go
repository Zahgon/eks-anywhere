package codebuild

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"

	"github.com/aws/eks-anywhere-test-tool/pkg/awsprofiles"
)

type Codebuild struct {
	session *session.Session
	svc     *codebuild.CodeBuild
}

func New(account awsprofiles.EksAccount) (*Codebuild, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Codebuild) FetchBuildForProject(id string) (*codebuild.Build, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Codebuild) FetchLatestBuildForProject(project string) (*codebuild.Build, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find latest build that is not in progress

func (c *Codebuild) getBuildById(id string) (*codebuild.Build, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Codebuild) FetchBuildsForProject(project string) *codebuild.ListBuildsForProjectOutput {
	_ = "STUB: not implemented"
	// we're using this to get the latest build, so we don't care about pagination at the moment
	return nil
}
