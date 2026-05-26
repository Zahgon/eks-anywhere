package providerproxy

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	"github.com/aws/eks-anywhere-test-tool/pkg/filewriter"
)

type requestWriter struct {
	filewriter.FileWriter
}

func newRequestWriter(folderPath string) (*requestWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *requestWriter) writeRequest(logs []*cloudwatchlogs.OutputLogEvent) error {
	_ = "STUB: not implemented"
	return nil
}
