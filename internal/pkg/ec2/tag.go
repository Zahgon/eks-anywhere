package ec2

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func TagInstance(session *session.Session, instanceId, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}
