package e2e

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var svc *cloudwatch.CloudWatch

const integrationTestCloudWatchNamespaceOverrideEnvVar = "INTEGRATION_TEST_CLOUDWATCH_NAMESPACE_OVERRIDE"

func init() {
	if s, err := session.NewSession(); err == nil {
		svc = cloudwatch.New(s)
	} else {
		fmt.Println("Cannot create CloudWatch service", err)
	}
}

func putInstanceTestResultMetrics(r instanceTestsResults) { _ = "STUB: not implemented"; return }

// Note 0 metrics are emitted for the purpose of aggregation. For example, when the succeededCount metrics are [0, 1, 0, 1], we can calculate the success rate as 2 / 4 = 50%. However, when 0 are excluded, the metrics becomes [1, 1], and you would not be able to calculate the success rate from that series.

// TODO: publish time metrics

func getProviderName(testRe string) string { _ = "STUB: not implemented"; return "" }

func putMetric(data *cloudwatch.MetricDatum, metricName string, value int) {
	_ = "STUB: not implemented"
	return
}
