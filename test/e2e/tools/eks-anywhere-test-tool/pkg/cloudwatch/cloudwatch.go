package cloudwatch

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	"github.com/aws/eks-anywhere-test-tool/pkg/awsprofiles"
)

type Cloudwatch struct {
	session *session.Session
	svc     *cloudwatchlogs.CloudWatchLogs
}

func New(account awsprofiles.EksAccount) (*Cloudwatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cloudwatch) GetLogs(logGroupName string, logStreamName string) ([]*cloudwatchlogs.OutputLogEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cloudwatch) GetLogsInTimeframe(logGroupName string, logStreamName string, startTime int64, endTime int64) ([]*cloudwatchlogs.OutputLogEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cloudwatch) getLogs(logGroupName string, logStreamName string, startTime *int64, endTime *int64) ([]*cloudwatchlogs.OutputLogEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Cloudwatch) getLogSegment(logGroupName string, logStreamName string, startTime *int64, endTime *int64, nextToken *string) (*cloudwatchlogs.GetLogEventsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isInvalidParameterError(err error) bool { _ = "STUB: not implemented"; return false }
