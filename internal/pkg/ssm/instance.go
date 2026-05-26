package ssm

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func GetInstanceByActivationId(session *session.Session, id string) (*ssm.InstanceInformation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeregisterInstance(session *session.Session, id string) (*ssm.DeregisterManagedInstanceOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
