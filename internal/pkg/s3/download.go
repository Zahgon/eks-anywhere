package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws/session"
)

func Download(session *session.Session, key, bucket string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func download(session *session.Session, key, bucket string, w io.WriterAt) error {
	_ = "STUB: not implemented"
	return nil
}

func DownloadToDisk(session *session.Session, key, bucket, dst string) error {
	_ = "STUB: not implemented"
	return nil
}
