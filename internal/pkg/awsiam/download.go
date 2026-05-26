package awsiam

import (
	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
)

const awsIamClientBinary = "aws-iam-authenticator"

func DownloadAwsIamAuthClient(eksdRelease *eksdv1alpha1.Release) error {
	_ = "STUB: not implemented"
	return nil
}

func getKernelName() string { _ = "STUB: not implemented"; return "" }

func getAwsIamAuthClientUri(eksdRelease *eksdv1alpha1.Release, kernalName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
