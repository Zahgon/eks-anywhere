package clustermanager

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
)

// KubernetesRetrierClient wraps around a KubernetesClient, offering retry functionality for some operations.
type KubernetesRetrierClient struct {
	KubernetesClient
	retrier *retrier.Retrier
}

// NewRetrierClient constructs a new RetrierClient.
func NewRetrierClient(client KubernetesClient, retrier *retrier.Retrier) *KubernetesRetrierClient {
	_ = "STUB: not implemented"
	return nil
}

// ApplyKubeSpecFromBytes creates/updates the objects defined in a yaml manifest against the api server following a client side apply mechanism.
func (c *KubernetesRetrierClient) ApplyKubeSpecFromBytes(ctx context.Context, cluster *types.Cluster, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Apply creates/updates an object against the api server following a client side apply mechanism.
func (c *KubernetesRetrierClient) Apply(ctx context.Context, kubeconfigPath string, obj runtime.Object, opts ...kubernetes.KubectlApplyOption) error {
	_ = "STUB: not implemented"
	return nil
}

// PauseCAPICluster adds a `spec.Paused: true` to the CAPI cluster resource. This will cause all
// downstream CAPI + provider controllers to skip reconciling on the paused cluster's objects.
func (c *KubernetesRetrierClient) PauseCAPICluster(ctx context.Context, cluster, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

// ResumeCAPICluster removes the `spec.Paused` on the CAPI cluster resource. This will cause all
// downstream CAPI + provider controllers to resume reconciling on the paused cluster's objects.
func (c *KubernetesRetrierClient) ResumeCAPICluster(ctx context.Context, cluster, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyKubeSpecFromBytesForce creates/updates the objects defined in a yaml manifest against the api server following a client side apply mechanism.
// It forces the operation, so if api validation failed, it will delete and re-create the object.
func (c *KubernetesRetrierClient) ApplyKubeSpecFromBytesForce(ctx context.Context, cluster *types.Cluster, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyKubeSpecFromBytesWithNamespace creates/updates the objects defined in a yaml manifest against the api server following a client side apply mechanism.
// It applies all objects in the given namespace.
func (c *KubernetesRetrierClient) ApplyKubeSpecFromBytesWithNamespace(ctx context.Context, cluster *types.Cluster, data []byte, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateAnnotationInNamespace adds/updates an annotation for the given kubernetes resource.
func (c *KubernetesRetrierClient) UpdateAnnotationInNamespace(ctx context.Context, resourceType, objectName string, annotations map[string]string, cluster *types.Cluster, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveAnnotationInNamespace deletes an annotation for the given kubernetes resource if present.
func (c *KubernetesRetrierClient) RemoveAnnotationInNamespace(ctx context.Context, resourceType, objectName, key string, cluster *types.Cluster, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// ListObjects reads all Objects of a particular resource type in a namespace.
func (c *KubernetesRetrierClient) ListObjects(ctx context.Context, resourceType, namespace, kubeconfig string, list kubernetes.ObjectList) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteGitOpsConfig deletes a GitOpsConfigObject from the cluster.
func (c *KubernetesRetrierClient) DeleteGitOpsConfig(ctx context.Context, cluster *types.Cluster, name string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteEKSACluster deletes an EKSA Cluster object from the cluster.
func (c *KubernetesRetrierClient) DeleteEKSACluster(ctx context.Context, cluster *types.Cluster, name string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAWSIamConfig deletes an AWSIamConfig object from the cluster.
func (c *KubernetesRetrierClient) DeleteAWSIamConfig(ctx context.Context, cluster *types.Cluster, name string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteOIDCConfig deletes a OIDCConfig object from the cluster.
func (c *KubernetesRetrierClient) DeleteOIDCConfig(ctx context.Context, cluster *types.Cluster, name string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteCluster deletes a CAPI Cluster from the cluster.
func (c *KubernetesRetrierClient) DeleteCluster(ctx context.Context, cluster, clusterToDelete *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
