package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func ObjectPresent(session *session.Session, key, bucket string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
