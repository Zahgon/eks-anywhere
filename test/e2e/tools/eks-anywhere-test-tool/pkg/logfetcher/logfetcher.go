package logfetcher

import (
	"bytes"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	awscodebuild "github.com/aws/aws-sdk-go/service/codebuild"

	"github.com/aws/eks-anywhere-test-tool/pkg/cloudwatch"
	"github.com/aws/eks-anywhere-test-tool/pkg/codebuild"
	"github.com/aws/eks-anywhere-test-tool/pkg/testresults"
)

type FetchLogsOpt func(options *fetchLogsConfig) (err error)

func WithCodebuildBuild(buildId string) FetchLogsOpt {
	_ = "STUB: not implemented"
	return *new(FetchLogsOpt)
}

func WithCodebuildProject(project string) FetchLogsOpt {
	_ = "STUB: not implemented"
	return *new(FetchLogsOpt)
}

type fetchLogsConfig struct {
	buildId string
	project string
}

var ssmCommandExecutionLogStreamTemplate = "%s/%s/aws-runShellScript/%s"

type (
	codebuildConsumer func(*awscodebuild.Build) error
	messagesConsumer  func(allMessages, filteredMessages *bytes.Buffer) error
	testConsumer      func(testName string, logs []*cloudwatchlogs.OutputLogEvent) error
)

type LogFetcherOpt func(*testLogFetcher)

func WithTestFilterByName(tests []string) LogFetcherOpt {
	_ = "STUB: not implemented"
	return *new(LogFetcherOpt)
}

func WithLogStdout() LogFetcherOpt { _ = "STUB: not implemented"; return *new(LogFetcherOpt) }

type testLogFetcher struct {
	buildAccountCwClient        *cloudwatch.Cloudwatch
	testAccountCwClient         *cloudwatch.Cloudwatch
	buildAccountCodebuildClient *codebuild.Codebuild
	writer                      *testsWriter
	filterTests                 testresults.TestFilter
	processCodebuild            codebuildConsumer
	processMessages             messagesConsumer
	processTest                 testConsumer
}

func New(buildAccountCwClient *cloudwatch.Cloudwatch, testAccountCwClient *cloudwatch.Cloudwatch, buildAccountCodebuildClient *codebuild.Codebuild, opts ...LogFetcherOpt) *testLogFetcher {
	_ = "STUB: not implemented"
	return nil
}

func (l *testLogFetcher) FetchLogs(opts ...FetchLogsOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *testLogFetcher) GetBuildProjectLogs(project string, buildId string) ([]testresults.TestResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *testLogFetcher) FetchTestLogs(tests []testresults.TestResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *testLogFetcher) ensureWriter(folderPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func allMessages(logs []*cloudwatchlogs.OutputLogEvent) *bytes.Buffer {
	_ = "STUB: not implemented"
	return nil
}
