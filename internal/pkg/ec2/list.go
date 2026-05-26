package ec2

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func ListInstances(session *session.Session, key string, value string, maxAge float64) ([]*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
