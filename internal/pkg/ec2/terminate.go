package ec2

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

const defaultInstanceIDBlockSize = 500

// TerminateEc2Instances terminates EC2 instances by calling the AWS API.
func TerminateEc2Instances(session *session.Session, instanceIDs []*string) error {
	_ = "STUB: not implemented"
	return nil
}

func makeChunks[T any](elements []T, chunkSize int) [][]T { _ = "STUB: not implemented"; return nil }
