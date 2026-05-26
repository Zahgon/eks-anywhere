package ssm

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type ActivationInfo struct {
	ActivationCode string
	ActivationID   string
}

func CreateActivation(session *session.Session, instanceName, role string) (*ActivationInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteActivation(session *session.Session, activationId string) (*ssm.DeleteActivationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
