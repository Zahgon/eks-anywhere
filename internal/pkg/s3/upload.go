package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadOpt func(*s3manager.UploadInput)

func WithPublicRead() UploadOpt { _ = "STUB: not implemented"; return *new(UploadOpt) }

func UploadFile(session *session.Session, file, key, bucket string, opts ...UploadOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func Upload(session *session.Session, body []byte, key, bucket string, opts ...UploadOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func upload(session *session.Session, body io.Reader, key, bucket string, opts ...UploadOpt) error {
	_ = "STUB: not implemented"
	return nil
}
