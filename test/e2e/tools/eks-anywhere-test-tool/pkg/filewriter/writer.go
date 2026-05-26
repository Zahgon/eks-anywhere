package filewriter

type writer struct {
	dir string
}

func NewWriter(dir string) FileWriter { _ = "STUB: not implemented"; return *new(FileWriter) }

func (t *writer) Write(fileName string, content []byte, f ...FileOptionsFunc) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Default file options. -->> temporary file with default permissions

func (w *writer) WithDir(dir string) (FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(FileWriter), nil
}

func (t *writer) Dir() string {
	_ = "STUB: not implemented"

	// This method writes the e2e test artifacts from S3 to files in a directory named after the e2e test name.
	return ""
}

func (t *writer) WriteTestArtifactsS3ToFile(key string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *writer) CleanUp() { _ = "STUB: not implemented"; return }

func (t *writer) CleanUpTemp() { _ = "STUB: not implemented"; return }
