package snow

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	snowv1 "github.com/aws/eks-anywhere/pkg/providers/snow/api/v1beta1"
)

// BaseControlPlane represents a CAPI Snow control plane.
type BaseControlPlane = clusterapi.ControlPlane[*snowv1.AWSSnowCluster, *snowv1.AWSSnowMachineTemplate]

// ControlPlane holds the Snow specific objects for a CAPI snow control plane.
type ControlPlane struct {
	BaseControlPlane
	Secret       *corev1.Secret
	CAPASIPPools CAPASIPPools
}

// Objects returns the control plane objects associated with the snow cluster.
func (c ControlPlane) Objects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// ControlPlaneSpec builds a snow ControlPlane definition based on an eks-a cluster spec.
func ControlPlaneSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, clusterSpec *cluster.Spec) (*ControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// credentialsSecret generates the credentials secret(s) used for provisioning a snow cluster.
// - eks-a credentials secret: user managed secret referred from snowdatacenterconfig identityRef
// - snow credentials secret: eks-a creates, updates and deletes in eksa-system namespace. this secret is fully managed by eks-a. User shall treat it as a "read-only" object.
func capasCredentialsSecret(clusterSpec *cluster.Spec) (*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we reconcile the snow credentials secret to be in sync with the eks-a credentials secret user manages.
// notice for cli upgrade, we handle the eks-a credentials secret update in a separate step - under provider.UpdateSecrets
// which runs before the actual cluster upgrade.
// for controller secret, the user is responsible for making sure the eks-a credentials secret is created and up to date.
