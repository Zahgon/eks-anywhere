package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/eks-anywhere-test-tool/pkg/awsprofiles"
	"github.com/aws/eks-anywhere/pkg/retrier"
)

type S3 struct {
	session *session.Session
	svc     *s3.S3
	retrier *retrier.Retrier
}

func New(account awsprofiles.EksAccount) (*S3, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *S3) ListObjects(bucket string, prefix string) (listedObjects []*s3.Object, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3) GetObject(bucket string, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getObjectRetirer() *retrier.Retrier { _ = "STUB: not implemented"; return nil }

func isThrottledError(err error) bool { _ = "STUB: not implemented"; return false }
