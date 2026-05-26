package providerproxy

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	"github.com/aws/eks-anywhere-test-tool/pkg/cloudwatch"
	"github.com/aws/eks-anywhere-test-tool/pkg/codebuild"
)

type FetchSessionOpts func(options *fetchSessionsConfig) (err error)

func WithCodebuildBuild(buildId string) FetchSessionOpts {
	_ = "STUB: not implemented"
	return *new(FetchSessionOpts)
}

func WithCodebuildProject(project string) FetchSessionOpts {
	_ = "STUB: not implemented"
	return *new(FetchSessionOpts)
}

type fetchSessionsConfig struct {
	buildId string
	project string
}

type (
	requestFilter   func(logs []*cloudwatchlogs.OutputLogEvent) (filteredLogs []*cloudwatchlogs.OutputLogEvent, err error)
	requestConsumer func(logs []*cloudwatchlogs.OutputLogEvent) error
)

type ProxyFetcherOpt func(*proxyLogFetcher)

func WithLogStdout() ProxyFetcherOpt { _ = "STUB: not implemented"; return *new(ProxyFetcherOpt) }

type proxyLogFetcher struct {
	buildAccountCwClient        *cloudwatch.Cloudwatch
	testAccountCwClient         *cloudwatch.Cloudwatch
	buildAccountCodebuildClient *codebuild.Codebuild
	writer                      *requestWriter
	filterRequests              requestFilter
	processRequests             requestConsumer
}

func New(buildAccountCwClient *cloudwatch.Cloudwatch, testAccountCwClient *cloudwatch.Cloudwatch, buildAccountCodebuildClient *codebuild.Codebuild, opts ...ProxyFetcherOpt) *proxyLogFetcher {
	_ = "STUB: not implemented"
	return nil
}

func (l *proxyLogFetcher) FetchProviderProxyLogs(opts ...FetchSessionOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *proxyLogFetcher) FetchProviderProxyLogsForbuild(project string, buildId string) ([]*cloudwatchlogs.OutputLogEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *proxyLogFetcher) ensureWriter(folderPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func noFilter(logs []*cloudwatchlogs.OutputLogEvent) (outputLogs []*cloudwatchlogs.OutputLogEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
