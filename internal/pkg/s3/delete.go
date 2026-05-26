package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func CleanUpS3Bucket(session *session.Session, bucket string, maxAge float64) error {
	_ = "STUB: not implemented"
	return nil
}
