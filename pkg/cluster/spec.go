package cluster

import (
	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	eksav1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type Spec struct {
	*Config
	Bundles           *v1alpha1.Bundles
	OIDCConfig        *eksav1alpha1.OIDCConfig
	AWSIamConfig      *eksav1alpha1.AWSIamConfig
	ManagementCluster *types.Cluster // TODO(g-gaston): cleanup, this doesn't belong here
	EKSARelease       *v1alpha1.EKSARelease
	VersionsBundles   map[eksav1alpha1.KubernetesVersion]*VersionsBundle
}

func (s *Spec) DeepCopy() *Spec { _ = "STUB: not implemented"; return nil }

type VersionsBundle struct {
	*v1alpha1.VersionsBundle
	KubeDistro *KubeDistro
}

func deepCopyVersionsBundles(v map[eksav1alpha1.KubernetesVersion]*VersionsBundle) map[eksav1alpha1.KubernetesVersion]*VersionsBundle {
	_ = "STUB: not implemented"
	return nil
}

// EKSD represents an eks-d release.
type EKSD struct {
	// Channel is the minor Kubernetes version for the eks-d release (eg. "1.23", "1.24", etc.)
	Channel string
	// Number is the monotonically increasing number that distinguishes the different eks-d releases
	// for the same Kubernetes minor version (channel).
	Number int
}

func (k *KubeDistro) deepCopy() *KubeDistro { _ = "STUB: not implemented"; return nil }

type KubeDistro struct {
	EKSD            EKSD
	Kubernetes      VersionedRepository
	CoreDNS         VersionedRepository
	Etcd            VersionedRepository
	Pause           v1alpha1.Image
	EtcdImage       v1alpha1.Image
	EtcdVersion     string
	EtcdURL         string
	AwsIamAuthImage v1alpha1.Image
	KubeProxy       v1alpha1.Image
}

type VersionedRepository struct {
	Repository, Tag string
}

// NewSpec builds a new [Spec].
func NewSpec(config *Config, bundles *v1alpha1.Bundles, eksdReleases []eksdv1alpha1.Release, eksaRelease *v1alpha1.EKSARelease) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get first aws iam config if it exists
// Config supports multiple configs because Cluster references a slice
// But we validate that only one of each type is referenced

// Get first oidc config if it exists

func getAllVersionsBundles(cluster *eksav1alpha1.Cluster, bundles *v1alpha1.Bundles, eksdReleases []eksdv1alpha1.Release) (map[eksav1alpha1.KubernetesVersion]*VersionsBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getVersionBundles(version eksav1alpha1.KubernetesVersion, b *v1alpha1.Bundles, eksdRelease *eksdv1alpha1.Release) (*VersionsBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VersionsBundle returns a VersionsBundle if one exists for the provided kubernetes version and nil otherwise.
func (s *Spec) VersionsBundle(version eksav1alpha1.KubernetesVersion) *VersionsBundle {
	_ = "STUB: not implemented"
	return nil
}

// RootVersionsBundle returns a VersionsBundle for the Cluster objects root Kubernetes versions.
func (s *Spec) RootVersionsBundle() *VersionsBundle { _ = "STUB: not implemented"; return nil }

// WorkerNodeGroupVersionsBundle returns a VersionsBundle for the Worker Node's kubernetes version.
func (s *Spec) WorkerNodeGroupVersionsBundle(w eksav1alpha1.WorkerNodeGroupConfiguration) *VersionsBundle {
	_ = "STUB: not implemented"
	return nil
}

func buildKubeDistro(eksd *eksdv1alpha1.Release) (*KubeDistro, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get archive uri for amd64

func kubeDistroRepository(image *eksdv1alpha1.AssetImage) (repo, tag string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (vb *VersionsBundle) Ovas() []v1alpha1.Archive { _ = "STUB: not implemented"; return nil }
