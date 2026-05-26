package common

import (
	auditv1 "k8s.io/apiserver/pkg/apis/audit/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// GetAuditPolicy returns the audit policy either v1 or v1beta1 depending on kube version.
func GetAuditPolicy(kubeVersion v1alpha1.KubernetesVersion) (string, error) {
	_ = "STUB: not implemented"
	// appending the ".0" as the patch version to have a valid semver string and use those semvers for comparison
	return "", nil
}

// AuditPolicyV1Yaml returns the byte array for yaml created with v1 api version for audit policy.
func AuditPolicyV1Yaml() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// AuditPolicyV1 returns the v1 audit policy.
func AuditPolicyV1() *auditv1.Policy { _ = "STUB: not implemented"; return nil }
