package templater

import (
	"github.com/aws/eks-anywhere/pkg/filewriter"
)

type Templater struct {
	writer filewriter.FileWriter
}

func New(writer filewriter.FileWriter) *Templater { _ = "STUB: not implemented"; return nil }

func (t *Templater) WriteToFile(templateContent string, data interface{}, fileName string, f ...filewriter.FileOptionsFunc) (filePath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *Templater) WriteBytesToFile(content []byte, fileName string, f ...filewriter.FileOptionsFunc) (filePath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func Execute(templateContent string, data interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	// Apply sprig functions for easy templating.
	// See https://masterminds.github.io/sprig/ for a list of available functions.
	return nil, nil
}

func toYAML(v any) string { _ = "STUB: not implemented"; return "" }
