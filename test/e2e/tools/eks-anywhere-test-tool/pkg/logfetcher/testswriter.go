package logfetcher

import (
	"bytes"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	awscodebuild "github.com/aws/aws-sdk-go/service/codebuild"

	"github.com/aws/eks-anywhere-test-tool/pkg/filewriter"
)

type testsWriter struct {
	filewriter.FileWriter
}

func newTestsWriter(folderPath string) (*testsWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *testsWriter) writeCodeBuild(build *awscodebuild.Build) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *testsWriter) writeMessages(allMessages, filteredMessages *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *testsWriter) writeTest(testName string, logs []*cloudwatchlogs.OutputLogEvent) error {
	_ = "STUB: not implemented"
	return nil
}
