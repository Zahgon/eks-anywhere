package reconciler

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/awsiamauth"
	anywhereCluster "github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
	"github.com/aws/eks-anywhere/pkg/crypto"
)

// RemoteClientRegistry defines methods for remote cluster controller clients.
type RemoteClientRegistry interface {
	GetClient(ctx context.Context, cluster client.ObjectKey) (client.Client, error)
}

// Reconciler allows to reconcile AWSIamConfig.
type Reconciler struct {
	certgen              crypto.CertificateGenerator
	templateBuilder      *awsiamauth.TemplateBuilder
	generateUUID         UUIDGenerator
	client               client.Client
	remoteClientRegistry RemoteClientRegistry
}

// UUIDGenerator generates a new UUID.
type UUIDGenerator func() uuid.UUID

// New returns a new Reconciler.
func New(certgen crypto.CertificateGenerator, generateUUID UUIDGenerator, client client.Client, remoteClientRegistry RemoteClientRegistry) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

// EnsureCASecret ensures the AWS IAM Authenticator secret is present.
// It uses a controller.Result to indicate when requeues are needed.
func (r *Reconciler) EnsureCASecret(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func (r *Reconciler) createCASecret(ctx context.Context, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile takes the AWS IAM Authenticator installation to the desired state defined in AWSIAMConfig.
// It uses a controller.Result to indicate when requeues are needed.
// Intended to be used in a kubernetes controller.
func (r *Reconciler) Reconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// CheckControlPlaneReady was not meant to be used here.
// It was intended as a phase in cluster reconciliation.
// TODO (pokearu): Break down the function to better reuse it.

// If configmap is not found, this is a first time install of aws-iam-authenticator on the cluster.
// We use a newly generated UUID.
// The configmap clusterID and kubeconfig token need to match. Hence the kubeconfig secret is created for first install.

func (r *Reconciler) applyIAMAuthManifest(ctx context.Context, client client.Client, clusterSpec *anywhereCluster.Spec, clusterID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) createKubeconfigSecret(ctx context.Context, clusterSpec *anywhereCluster.Spec, cluster *anywherev1.Cluster, clusterID uuid.UUID) error {
	_ = "STUB: not implemented"
	return nil
}

// ReconcileWorkloadClusterDelete deletes AWS IAM authenticator resources from the workload cluster
// This method uses the template-based approach to generate the manifest and then delete the resources.
func (r *Reconciler) ReconcileWorkloadClusterDelete(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster, awsIamConfig *anywherev1.AWSIamConfig) error {
	_ = "STUB: not implemented"
	// Build cluster spec with the provided AWSIamConfig
	return nil
}

// Override the AWSIamConfig in the spec with the provided one

// Get workload cluster client

// Generate the manifest using the template builder
// Use uuid.Nil since we don't need cluster ID for deletion

// Delete the resources using the serverside delete functions

// ReconcileDelete deletes any AWS Iam authenticator specific resources leftover on the eks-a cluster.
func (r *Reconciler) ReconcileDelete(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
