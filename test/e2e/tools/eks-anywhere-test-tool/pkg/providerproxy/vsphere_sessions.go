package providerproxy

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func VsphereSessionsFilter(logs []*cloudwatchlogs.OutputLogEvent) (outputLogs []*cloudwatchlogs.OutputLogEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
