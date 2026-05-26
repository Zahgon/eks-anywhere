package testresults

import (
	"bytes"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type TestFilter func(logs []*cloudwatchlogs.OutputLogEvent) (filteredTestsLogs *bytes.Buffer, filteredTestResults []TestResult, err error)

type TestResult struct {
	InstanceId string `json:"instanceId"`
	JobId      string `json:"jobId"`
	CommandId  string `json:"commandId"`
	Tests      string `json:"tests"`
	Status     string `json:"status"`
	Error      string `json:"error"`
}

func GetFailedTests(logs []*cloudwatchlogs.OutputLogEvent) (failedTestMessages *bytes.Buffer, failedTestResults []TestResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func NewTestFilterByName(tests []string) TestFilter {
	_ = "STUB: not implemented"
	return *new(TestFilter)
}

func TestResultsJobIdMap(tests []TestResult) map[string]bool { _ = "STUB: not implemented"; return nil }

func isResultMessage(message string) bool { _ = "STUB: not implemented"; return false }
