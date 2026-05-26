package artifacts

import (
	"github.com/aws/eks-anywhere-test-tool/pkg/cloudwatch"
	"github.com/aws/eks-anywhere-test-tool/pkg/codebuild"
	"github.com/aws/eks-anywhere-test-tool/pkg/filewriter"
	"github.com/aws/eks-anywhere-test-tool/pkg/s3"
	"github.com/aws/eks-anywhere/pkg/retrier"
)

type FetchArtifactsOpt func(options *fetchArtifactConfig) (err error)

func WithCodebuildBuild(buildId string) FetchArtifactsOpt {
	_ = "STUB: not implemented"
	return *new(FetchArtifactsOpt)
}

func WithCodebuildProject(project string) FetchArtifactsOpt {
	_ = "STUB: not implemented"
	return *new(FetchArtifactsOpt)
}

func WithAllArtifacts() FetchArtifactsOpt {
	_ = "STUB: not implemented"
	return *new(FetchArtifactsOpt)
}

type fetchArtifactConfig struct {
	buildId  string
	bucket   string
	project  string
	fetchAll bool
}

type testArtifactFetcher struct {
	testAccountS3Client         *s3.S3
	buildAccountCodebuildClient *codebuild.Codebuild
	buildAccountCwClient        *cloudwatch.Cloudwatch
	writer                      filewriter.FileWriter
	retrier                     *retrier.Retrier
}

func New(testAccountS3Client *s3.S3, buildAccountCodebuildCient *codebuild.Codebuild, writer filewriter.FileWriter, cwClient *cloudwatch.Cloudwatch) *testArtifactFetcher {
	_ = "STUB: not implemented"
	return nil
}

func (l *testArtifactFetcher) FetchArtifacts(opts ...FetchArtifactsOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func excludedKey(key string) bool { _ = "STUB: not implemented"; return false }

func fileWriterRetrier() *retrier.Retrier { _ = "STUB: not implemented"; return nil }

func isTooManyOpenFilesError(err error) bool { _ = "STUB: not implemented"; return false }
