package executables

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	tinkv1alpha1 "github.com/tinkerbell/tink/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/version"
	cloudstackv1 "sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta3"
	vspherev1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"
	addons "sigs.k8s.io/cluster-api/api/addons/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	packagesv1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	rufiov1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/rufio"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/rufiounreleased"
	"github.com/aws/eks-anywhere/pkg/types"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	kubectlPath        = "kubectl"
	timeoutPrecision   = 2
	minimumWaitTimeout = 0.01 // Smallest express-able timeout value given the precision

	networkFaultBaseRetryTime = 10 * time.Second
	networkFaultBackoffFactor = 1.5

	lastAppliedAnnotation = "kubectl.kubernetes.io/last-applied-configuration"
)

var (
	capiClustersResourceType             = fmt.Sprintf("clusters.%s", clusterv1beta2.GroupVersion.Group)
	capiProvidersResourceType            = fmt.Sprintf("providers.clusterctl.%s", clusterv1beta2.GroupVersion.Group)
	capiMachinesType                     = fmt.Sprintf("machines.%s", clusterv1beta2.GroupVersion.Group)
	capiMachineDeploymentsType           = fmt.Sprintf("machinedeployments.%s", clusterv1beta2.GroupVersion.Group)
	capiMachineSetsType                  = fmt.Sprintf("machinesets.%s", clusterv1beta2.GroupVersion.Group)
	eksaClusterResourceType              = fmt.Sprintf("clusters.%s", v1alpha1.GroupVersion.Group)
	eksaVSphereDatacenterResourceType    = fmt.Sprintf("vspheredatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaVSphereMachineResourceType       = fmt.Sprintf("vspheremachineconfigs.%s", v1alpha1.GroupVersion.Group)
	vsphereMachineTemplatesType          = fmt.Sprintf("vspheremachinetemplates.infrastructure.%s", clusterv1beta2.GroupVersion.Group)
	eksaTinkerbellDatacenterResourceType = fmt.Sprintf("tinkerbelldatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaTinkerbellMachineResourceType    = fmt.Sprintf("tinkerbellmachineconfigs.%s", v1alpha1.GroupVersion.Group)
	TinkerbellHardwareResourceType       = fmt.Sprintf("hardware.%s", tinkv1alpha1.GroupVersion.Group)
	rufioMachineResourceType             = fmt.Sprintf("machines.%s", rufiov1alpha1.GroupVersion.Group)
	eksaCloudStackDatacenterResourceType = fmt.Sprintf("cloudstackdatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaCloudStackMachineResourceType    = fmt.Sprintf("cloudstackmachineconfigs.%s", v1alpha1.GroupVersion.Group)
	cloudstackMachineTemplatesType       = fmt.Sprintf("cloudstackmachinetemplates.infrastructure.%s", clusterv1beta2.GroupVersion.Group)
	eksaNutanixDatacenterResourceType    = fmt.Sprintf("nutanixdatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaNutanixMachineResourceType       = fmt.Sprintf("nutanixmachineconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaAwsResourceType                  = fmt.Sprintf("awsdatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaGitOpsResourceType               = fmt.Sprintf("gitopsconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaFluxConfigResourceType           = fmt.Sprintf("fluxconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaOIDCResourceType                 = fmt.Sprintf("oidcconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaAwsIamResourceType               = fmt.Sprintf("awsiamconfigs.%s", v1alpha1.GroupVersion.Group)
	etcdadmClustersResourceType          = fmt.Sprintf("etcdadmclusters.%s", etcdv1.GroupVersion.Group)
	bundlesResourceType                  = fmt.Sprintf("bundles.%s", releasev1alpha1.GroupVersion.Group)
	clusterResourceSetResourceType       = fmt.Sprintf("clusterresourcesets.%s", addons.GroupVersion.Group)
	kubeadmControlPlaneResourceType      = fmt.Sprintf("kubeadmcontrolplanes.controlplane.%s", clusterv1beta2.GroupVersion.Group)
	eksdReleaseType                      = fmt.Sprintf("releases.%s", eksdv1alpha1.GroupVersion.Group)
	eksaPackagesType                     = fmt.Sprintf("packages.%s", packagesv1.GroupVersion.Group)
	eksaPackagesBundleControllerType     = fmt.Sprintf("packagebundlecontroller.%s", packagesv1.GroupVersion.Group)
	eksaPackageBundlesType               = fmt.Sprintf("packagebundles.%s", packagesv1.GroupVersion.Group)
	kubectlConnectionRefusedRegex        = regexp.MustCompile("The connection to the server .* was refused")
	kubectlConnectionTimeoutRegex        = regexp.MustCompile("Unable to connect to the server.*timeout.*")
)

type Kubectl struct {
	Executable
	// networkFaultBackoffFactor drives the exponential backoff wait
	// for transient network failures during retry operations.
	networkFaultBackoffFactor float64

	// networkFaultBaseRetryTime drives the base time wait for the
	// exponential backoff for transient network failures during retry operations.
	networkFaultBaseRetryTime time.Duration
}

// KubectlConfigOpt configures Kubectl on construction.
type KubectlConfigOpt func(*Kubectl)

// NewKubectl builds a new Kubectl.
func NewKubectl(executable Executable, opts ...KubectlConfigOpt) *Kubectl {
	_ = "STUB: not implemented"
	return nil
}

// WithKubectlNetworkFaultBaseRetryTime configures the base time wait for the
// exponential backoff for transient network failures during retry operations.
func WithKubectlNetworkFaultBaseRetryTime(wait time.Duration) KubectlConfigOpt {
	_ = "STUB: not implemented"
	return *new(KubectlConfigOpt)
}

// WithNetworkFaultBackoffFactor configures the exponential backoff wait
// for transient network failures during retry operations.
func WithNetworkFaultBackoffFactor(factor float64) KubectlConfigOpt {
	_ = "STUB: not implemented"
	return *new(KubectlConfigOpt)
}

type capiMachinesResponse struct {
	Items []clusterv1beta2.Machine
}

// GetCAPIMachines returns all the CAPI machines for the provided clusterName.
func (k *Kubectl) GetCAPIMachines(ctx context.Context, cluster *types.Cluster, clusterName string) ([]clusterv1beta2.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) SearchCloudStackMachineConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.CloudStackMachineConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) SearchCloudStackDatacenterConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.CloudStackDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaCloudStackMachineConfig(ctx context.Context, cloudstackMachineConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.CloudStackMachineConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) DeleteEksaCloudStackDatacenterConfig(ctx context.Context, cloudstackDatacenterConfigName string, kubeconfigFile string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) GetEksaCloudStackDatacenterConfig(ctx context.Context, cloudstackDatacenterConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.CloudStackDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) DeleteEksaCloudStackMachineConfig(ctx context.Context, cloudstackMachineConfigName string, kubeconfigFile string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

type VersionResponse struct {
	ClientVersion version.Info `json:"clientVersion"`
	ServerVersion version.Info `json:"serverVersion"`
}

func (k *Kubectl) GetNamespace(ctx context.Context, kubeconfig string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) CreateNamespace(ctx context.Context, kubeconfig string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) CreateNamespaceIfNotPresent(ctx context.Context, kubeconfig string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteNamespace(ctx context.Context, kubeconfig string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) LoadSecret(ctx context.Context, secretObject string, secretObjectType string, secretObjectName string, kubeConfFile string) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyManifest uses client-side logic to create/update objects defined in a yaml manifest.
func (k *Kubectl) ApplyManifest(ctx context.Context, kubeconfigPath, manifestPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ApplyKubeSpecWithNamespace(ctx context.Context, cluster *types.Cluster, spec string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ApplyKubeSpecFromBytes(ctx context.Context, cluster *types.Cluster, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ApplyKubeSpecFromBytesWithNamespace(ctx context.Context, cluster *types.Cluster, data []byte, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ApplyKubeSpecFromBytesForce(ctx context.Context, cluster *types.Cluster, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteManifest uses client-side logic to delete objects defined in a yaml manifest.
func (k *Kubectl) DeleteManifest(ctx context.Context, kubeconfigPath, manifestPath string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteKubeSpecFromBytes(ctx context.Context, cluster *types.Cluster, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) WaitForClusterReady(ctx context.Context, cluster *types.Cluster, timeout string, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) WaitForControlPlaneReady(ctx context.Context, cluster *types.Cluster, timeout string, newClusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) WaitForMachineDeploymentReady(ctx context.Context, cluster *types.Cluster, timeout string, machineDeploymentName string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForService blocks until an IP address is assigned.
//
// Until more generic status matching comes around (possibly in 1.23), poll
// the service, checking for an IP address. Would you like to know more?
// https://github.com/kubernetes/kubernetes/issues/83094
func (k *Kubectl) WaitForService(ctx context.Context, kubeconfig string, timeout string, target string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) WaitForDeployment(ctx context.Context, cluster *types.Cluster, timeout string, condition string, target string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForResourceRolledout waits for a resource (deployment, daemonset, or statefulset) to be successfully rolled out before returning.
func (k *Kubectl) WaitForResourceRolledout(ctx context.Context, cluster *types.Cluster, timeout string, target string, namespace string, resource string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForPod waits for a pod resource to reach desired condition before returning.
func (k *Kubectl) WaitForPod(ctx context.Context, cluster *types.Cluster, timeout string, condition string, target string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForRufioMachines blocks until all Rufio Machines have the desired condition.
// RufioMachines with the skip-contactability-check label set to "true" are excluded from the wait.
func (k *Kubectl) WaitForRufioMachines(ctx context.Context, cluster *types.Cluster, timeout string, condition string, namespace string) error {
	_ = "STUB: not implemented"
	// Use label selector to exclude machines with skip-contactability-check=true
	return nil
}

// WaitForJobCompleted waits for a job resource to reach desired condition before returning.
func (k *Kubectl) WaitForJobCompleted(ctx context.Context, kubeconfig, timeout string, condition string, target string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForPackagesInstalled waits for a package resource to reach installed state before returning.
func (k *Kubectl) WaitForPackagesInstalled(ctx context.Context, cluster *types.Cluster, name string, timeout string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForPodCompleted waits for a pod to be terminated with a Completed state before returning.
func (k *Kubectl) WaitForPodCompleted(ctx context.Context, cluster *types.Cluster, name string, timeout string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) Wait(ctx context.Context, kubeconfig string, timeout string, forCondition string, property string, namespace string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	// On each retry kubectl wait timeout values will have to be adjusted to only wait for the remaining timeout duration.
	// Here we establish an absolute timeout time for this based on the caller-specified timeout.
	return nil
}

// WaitJSONPathLoop will wait for a given JSONPath to reach a required state similar to wait command for objects without conditions.
// This will be deprecated in favor of WaitJSONPath after version 1.23.
func (k *Kubectl) WaitJSONPathLoop(ctx context.Context, kubeconfig string, timeout string, jsonpath, forCondition string, property string, namespace string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	// On each retry kubectl wait timeout values will have to be adjusted to only wait for the remaining timeout duration.
	//  Here we establish an absolute timeout time for this based on the caller-specified timeout.
	return nil
}

// WaitJSONPath will wait for a given JSONPath of a required state. Only compatible on K8s 1.23+.
func (k *Kubectl) WaitJSONPath(ctx context.Context, kubeconfig string, timeout string, jsonpath, forCondition string, property string, namespace string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	// On each retry kubectl wait timeout values will have to be adjusted to only wait for the remaining timeout duration.
	//  Here we establish an absolute timeout time for this based on the caller-specified timeout.
	return nil
}

func (k *Kubectl) kubectlWaitRetryPolicy(totalRetries int, err error) (retry bool, wait time.Duration) {
	_ = "STUB: not implemented"
	// Exponential backoff on network errors.  Retrier built-in backoff is linear, so implementing here.
	return false, *new(time.Duration)
}

// Retrier first calls the policy before retry #1.  We want it zero-based for exponentiation.

func (k *Kubectl) wait(ctx context.Context, kubeconfig string, timeoutTime time.Time, forCondition string, property string, namespace string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) waitJSONPath(ctx context.Context, kubeconfig, timeout string, jsonpath string, forCondition string, property string, namespace string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	return nil
}

// waitJsonPathLoop will be deprecated in favor of waitJsonPath after version 1.23.
func (k *Kubectl) waitJSONPathLoop(ctx context.Context, kubeconfig string, timeout string, jsonpath string, forCondition string, property string, namespace string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteEksaDatacenterConfig(ctx context.Context, eksaDatacenterResourceType string, eksaDatacenterConfigName string, kubeconfigFile string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteEksaMachineConfig(ctx context.Context, eksaMachineConfigResourceType string, eksaMachineConfigName string, kubeconfigFile string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteEKSACluster(ctx context.Context, managementCluster *types.Cluster, eksaClusterName, eksaClusterNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteGitOpsConfig(ctx context.Context, managementCluster *types.Cluster, gitOpsConfigName, gitOpsConfigNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteFluxConfig(ctx context.Context, managementCluster *types.Cluster, fluxConfigName, fluxConfigNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPackageBundleController will retrieve the packagebundlecontroller from eksa-packages namespace and return the object.
func (k *Kubectl) GetPackageBundleController(ctx context.Context, kubeconfigFile, clusterName string) (packagesv1.PackageBundleController, error) {
	_ = "STUB: not implemented"
	return *new(packagesv1.PackageBundleController), nil
}

// GetPackageBundleList will retrieve the packagebundle list from eksa-packages namespace and return the list.
func (k *Kubectl) GetPackageBundleList(ctx context.Context, kubeconfigFile string) ([]packagesv1.PackageBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) DeletePackageResources(ctx context.Context, managementCluster *types.Cluster, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteSecret(ctx context.Context, managementCluster *types.Cluster, secretName, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteOIDCConfig(ctx context.Context, managementCluster *types.Cluster, oidcConfigName, oidcConfigNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteAWSIamConfig(ctx context.Context, managementCluster *types.Cluster, awsIamConfigName, awsIamConfigNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteCluster(ctx context.Context, managementCluster, clusterToDelete *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ListCluster(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (k *Kubectl) GetNodes(ctx context.Context, kubeconfig string) ([]corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetControlPlaneNodes(ctx context.Context, kubeconfig string) ([]corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVsphereMachine will return list of vSphere machines.
func (k *Kubectl) GetVsphereMachine(ctx context.Context, kubeconfig string, selector string) ([]vspherev1.VSphereMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) ValidateNodes(ctx context.Context, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteOldWorkerNodeGroup(ctx context.Context, md *clusterv1beta2.MachineDeployment, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ValidateControlPlaneNodes(ctx context.Context, cluster *types.Cluster, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ValidateWorkerNodes(ctx context.Context, clusterName string, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) CountMachineDeploymentReplicasReady(ctx context.Context, clusterName string, kubeconfig string) (ready, total int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (k *Kubectl) VsphereWorkerNodesMachineTemplate(ctx context.Context, clusterName string, kubeconfig string, namespace string) (*vspherev1.VSphereMachineTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) CloudstackWorkerNodesMachineTemplate(ctx context.Context, clusterName string, kubeconfig string, namespace string) (*cloudstackv1.CloudStackMachineTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) MachineTemplateName(ctx context.Context, clusterName string, kubeconfig string, opts ...KubectlOpt) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *Kubectl) ValidatePods(ctx context.Context, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

// RunCurlPod will run Kubectl with an image (with curl installed) and the command you pass in.
func (k *Kubectl) RunCurlPod(ctx context.Context, namespace, name, kubeconfig, image string, command []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetPodNameByLabel will return the name of the first pod that matches the label.
func (k *Kubectl) GetPodNameByLabel(ctx context.Context, namespace, label, kubeconfig string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetPodIP will return the ip of the pod.
func (k *Kubectl) GetPodIP(ctx context.Context, namespace, podName, kubeconfig string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetPodLogs returns the logs of the specified container (namespace/pod/container).
func (k *Kubectl) GetPodLogs(ctx context.Context, namespace, podName, containerName, kubeconfig string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetPodLogsSince returns the logs of the specified container (namespace/pod/container) since a timestamp.
func (k *Kubectl) GetPodLogsSince(ctx context.Context, namespace, podName, containerName, kubeconfig string, since time.Time) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *Kubectl) getPodLogs(ctx context.Context, namespace, podName, containerName, kubeconfig string, sinceTime *metav1.Time, tailLines *int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *Kubectl) SaveLog(ctx context.Context, cluster *types.Cluster, deployment *types.Deployment, fileName string, writer filewriter.FileWriter) error {
	_ = "STUB: not implemented"
	return nil
}

type machinesResponse struct {
	Items []types.Machine `json:"items,omitempty"`
}

func (k *Kubectl) GetMachines(ctx context.Context, cluster *types.Cluster, clusterName string) ([]types.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type machineSetResponse struct {
	Items []clusterv1beta2.MachineSet `json:"items,omitempty"`
}

func (k *Kubectl) GetMachineSets(ctx context.Context, machineDeploymentName string, cluster *types.Cluster) ([]clusterv1beta2.MachineSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ClustersResponse struct {
	Items []types.CAPICluster `json:"items,omitempty"`
}

type GitOpsConfigResponse struct {
	Items []*v1alpha1.GitOpsConfig `json:"items,omitempty"`
}

type VSphereDatacenterConfigResponse struct {
	Items []*v1alpha1.VSphereDatacenterConfig `json:"items,omitempty"`
}

type CloudStackDatacenterConfigResponse struct {
	Items []*v1alpha1.CloudStackDatacenterConfig `json:"items,omitempty"`
}

// TinkerbellDatacenterConfigResponse contains list of TinkerbellDatacenterConfig.
type TinkerbellDatacenterConfigResponse struct {
	Items []*v1alpha1.TinkerbellDatacenterConfig `json:"items,omitempty"`
}

type NutanixDatacenterConfigResponse struct {
	Items []*v1alpha1.NutanixDatacenterConfig `json:"items,omitempty"`
}

type IdentityProviderConfigResponse struct {
	Items []*v1alpha1.Ref `json:"items,omitempty"`
}

type VSphereMachineConfigResponse struct {
	Items []*v1alpha1.VSphereMachineConfig `json:"items,omitempty"`
}

type CloudStackMachineConfigResponse struct {
	Items []*v1alpha1.CloudStackMachineConfig `json:"items,omitempty"`
}

// TinkerbellMachineConfigResponse contains list of TinkerbellMachineConfig.
type TinkerbellMachineConfigResponse struct {
	Items []*v1alpha1.TinkerbellMachineConfig `json:"items,omitempty"`
}

type NutanixMachineConfigResponse struct {
	Items []*v1alpha1.NutanixMachineConfig `json:"items,omitempty"`
}

func (k *Kubectl) ValidateClustersCRD(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ValidateEKSAClustersCRD(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) RolloutRestartDaemonSet(ctx context.Context, dsName, dsNamespace, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) SetEksaControllerEnvVar(ctx context.Context, envVar, envVarVal, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) GetClusters(ctx context.Context, cluster *types.Cluster) ([]types.CAPICluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetApiServerUrl(ctx context.Context, cluster *types.Cluster) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *Kubectl) Version(ctx context.Context, cluster *types.Cluster) (*VersionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type KubectlOpt func(*[]string)

// WithToken is a kubectl option to pass a token when making a kubectl call.
func WithToken(t string) KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

func WithServer(s string) KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

func WithCluster(c *types.Cluster) KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

func WithKubeconfig(kubeconfigFile string) KubectlOpt {
	_ = "STUB: not implemented"
	return *new(KubectlOpt)
}

func WithNamespace(n string) KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

// WithResourceName is a kubectl option to pass a resource name when making a kubectl call.
func WithResourceName(name string) KubectlOpt {
	_ = "STUB: not implemented"
	return *

	// WithAllNamespaces is a kubectl option to add all namespaces when making a kubectl call.
	new(KubectlOpt)
}

func WithAllNamespaces() KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

func WithSkipTLSVerify() KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

func WithOverwrite() KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

func WithWaitAll() KubectlOpt {
	_ = "STUB: not implemented"
	return *

	// WithSelector is a kubectl option to pass a selector when making kubectl calls.
	new(KubectlOpt)
}

func WithSelector(selector string) KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

func appendOpt(new ...string) KubectlOpt { _ = "STUB: not implemented"; return *new(KubectlOpt) }

func applyOpts(params *[]string, opts ...KubectlOpt) { _ = "STUB: not implemented"; return }

func (k *Kubectl) GetPods(ctx context.Context, opts ...KubectlOpt) ([]corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetDeployments(ctx context.Context, opts ...KubectlOpt) ([]appsv1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetSecretFromNamespace(ctx context.Context, kubeconfigFile, name, namespace string) (*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetSecret(ctx context.Context, secretObjectName string, opts ...KubectlOpt) (*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetKubeadmControlPlanes(ctx context.Context, opts ...KubectlOpt) ([]controlplanev1beta2.KubeadmControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetKubeadmControlPlane(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...KubectlOpt) (*controlplanev1beta2.KubeadmControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetMachineDeployment(ctx context.Context, workerNodeGroupName string, opts ...KubectlOpt) (*clusterv1beta2.MachineDeployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMachineDeployments retrieves all Machine Deployments.
func (k *Kubectl) GetMachineDeployments(ctx context.Context, opts ...KubectlOpt) ([]clusterv1beta2.MachineDeployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMachineDeploymentsForCluster retrieves all the Machine Deployments for a cluster with name "clusterName".
func (k *Kubectl) GetMachineDeploymentsForCluster(ctx context.Context, clusterName string, opts ...KubectlOpt) ([]clusterv1beta2.MachineDeployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) UpdateEnvironmentVariables(ctx context.Context, resourceType, resourceName string, envMap map[string]string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) UpdateEnvironmentVariablesInNamespace(ctx context.Context, resourceType, resourceName string, envMap map[string]string, cluster *types.Cluster, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) UpdateAnnotation(ctx context.Context, resourceType, objectName string, annotations map[string]string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) UpdateAnnotationInNamespace(ctx context.Context, resourceType, objectName string, annotations map[string]string, cluster *types.Cluster, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) RemoveAnnotation(ctx context.Context, resourceType, objectName string, key string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) RemoveAnnotationInNamespace(ctx context.Context, resourceType, objectName, key string, cluster *types.Cluster, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) GetEksaCluster(ctx context.Context, cluster *types.Cluster, clusterName string) (*v1alpha1.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) SearchVsphereMachineConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.VSphereMachineConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SearchTinkerbellMachineConfig returns the list of TinkerbellMachineConfig in the cluster.
func (k *Kubectl) SearchTinkerbellMachineConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.TinkerbellMachineConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) SearchIdentityProviderConfig(ctx context.Context, ipName string, kind string, kubeconfigFile string, namespace string) ([]*v1alpha1.VSphereDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) SearchVsphereDatacenterConfig(ctx context.Context, datacenterName string, kubeconfigFile string, namespace string) ([]*v1alpha1.VSphereDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SearchTinkerbellDatacenterConfig returns the list of TinkerbellDatacenterConfig in the cluster.
func (k *Kubectl) SearchTinkerbellDatacenterConfig(ctx context.Context, datacenterName string, kubeconfigFile string, namespace string) ([]*v1alpha1.TinkerbellDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaFluxConfig(ctx context.Context, gitOpsConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.FluxConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaGitOpsConfig(ctx context.Context, gitOpsConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.GitOpsConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaOIDCConfig(ctx context.Context, oidcConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.OIDCConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaAWSIamConfig(ctx context.Context, awsIamConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.AWSIamConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaTinkerbellDatacenterConfig(ctx context.Context, tinkerbellDatacenterConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.TinkerbellDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaVSphereDatacenterConfig(ctx context.Context, vsphereDatacenterConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.VSphereDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaTinkerbellMachineConfig(ctx context.Context, tinkerbellMachineConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.TinkerbellMachineConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUnprovisionedTinkerbellHardware retrieves unprovisioned Tinkerbell Hardware objects.
// Unprovisioned objects are those without any owner reference information.
func (k *Kubectl) GetUnprovisionedTinkerbellHardware(ctx context.Context, kubeconfig, namespace string) ([]tinkv1alpha1.Hardware, error) {
	_ = "STUB: not implemented"
	// Retrieve hardware resources that don't have the `v1alpha1.tinkerbell.org/ownerName` label.
	// This label is used to populate hardware when the CAPT controller acquires the Hardware
	// resource for provisioning.
	// See https://github.com/chrisdoherty4/cluster-api-provider-tinkerbell/blob/main/controllers/machine.go#L271
	return nil, nil
}

// GetProvisionedTinkerbellHardware retrieves provisioned Tinkerbell Hardware objects.
// Provisioned objects are those with owner reference information.
func (k *Kubectl) GetProvisionedTinkerbellHardware(ctx context.Context, kubeconfig, namespace string) ([]tinkv1alpha1.Hardware, error) {
	_ = "STUB: not implemented"
	// Retrieve hardware resources that have the `v1alpha1.tinkerbell.org/ownerName` label.
	// This label is used to populate hardware when the CAPT controller acquires the Hardware
	// resource for provisioning.
	return nil, nil
}

func (k *Kubectl) GetEksaVSphereMachineConfig(ctx context.Context, vsphereMachineConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.VSphereMachineConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaAWSDatacenterConfig(ctx context.Context, awsDatacenterConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.AWSDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetCurrentClusterContext(ctx context.Context, cluster *types.Cluster) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *Kubectl) GetEtcdadmCluster(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...KubectlOpt) (*etcdv1.EtcdadmCluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) ValidateNodesVersion(ctx context.Context, kubeconfig string, kubeVersion v1alpha1.KubernetesVersion) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) GetBundles(ctx context.Context, kubeconfigFile, name, namespace string) (*releasev1alpha1.Bundles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetClusterResourceSet(ctx context.Context, kubeconfigFile, name, namespace string) (*addons.ClusterResourceSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetConfigMap(ctx context.Context, kubeconfigFile, name, namespace string) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) SetDaemonSetImage(ctx context.Context, kubeconfigFile, name, namespace, container, image string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) setImage(ctx context.Context, kind, name, container, image string, opts ...KubectlOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) CheckProviderExists(ctx context.Context, kubeconfigFile, name, namespace string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type Toleration struct {
	Effect            string      `json:"effect,omitempty"`
	Key               string      `json:"key,omitempty"`
	Operator          string      `json:"operator,omitempty"`
	Value             string      `json:"value,omitempty"`
	TolerationSeconds json.Number `json:"tolerationSeconds,omitempty"`
}

func (k *Kubectl) ApplyTolerationsFromTaintsToDaemonSet(ctx context.Context, oldTaints []corev1.Taint, newTaints []corev1.Taint, dsName string, kubeconfigFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ApplyTolerationsFromTaints(ctx context.Context, oldTaints []corev1.Taint, newTaints []corev1.Taint, resource string, name string, kubeconfigFile string, namespace string, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// PauseCAPICluster adds a `spec.Paused: true` to the CAPI cluster resource. This will cause all
// downstream CAPI + provider controllers to skip reconciling on the paused cluster's objects.
func (k *Kubectl) PauseCAPICluster(ctx context.Context, cluster, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

// ResumeCAPICluster removes the `spec.Paused` on the CAPI cluster resource. This will cause all
// downstream CAPI + provider controllers to resume reconciling on the paused cluster's objects
// `spec.Paused` is set to `null` to drop the field instead of setting it to `false`.
func (k *Kubectl) ResumeCAPICluster(ctx context.Context, cluster, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

// MergePatchResource patches named resource using merge patch.
func (k *Kubectl) MergePatchResource(ctx context.Context, resource, name, patch, kubeconfig, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) KubeconfigSecretAvailable(ctx context.Context, kubeconfig string, clusterName string, namespace string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// HasResource implements KubectlRunner.
func (k *Kubectl) HasResource(ctx context.Context, resourceType string, name string, kubeconfig string, namespace string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetObject performs a GET call to the kube API server authenticating with a kubeconfig file
// and unmarshalls the response into the provided Object
// If the object is not found, it returns an error implementing apimachinery errors.APIStatus.
func (k *Kubectl) GetObject(ctx context.Context, resourceType, name, namespace, kubeconfig string, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// GetClusterObject performs a GET class like above except without namespace required.
func (k *Kubectl) GetClusterObject(ctx context.Context, resourceType, name, kubeconfig string, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ListObjects(ctx context.Context, resourceType, namespace, kubeconfig string, list kubernetes.ObjectList) error {
	_ = "STUB: not implemented"
	return nil
}

func withGetResourceName(name string) kubernetes.KubectlGetOption {
	_ = "STUB: not implemented"
	return *new(kubernetes.KubectlGetOption)
}

// withNamespaceOrDefaultForGet returns an option for a get command to use the provided namespace
// or the default namespace if an empty string is provided.
// For backwards compatibility, we us the default namespace if this method is called explicitly
// with an empty namespace since some parts of the code rely on kubectl using the default namespace
// when no namespace argument is passed.
func withNamespaceOrDefaultForGet(namespace string) kubernetes.KubectlGetOption {
	_ = "STUB: not implemented"
	return *new(kubernetes.KubectlGetOption)
}

func withClusterScope() kubernetes.KubectlGetOption {
	_ = "STUB: not implemented"
	return *new(kubernetes.KubectlGetOption)
}

// Get performs a kubectl get command.
func (k *Kubectl) Get(ctx context.Context, resourceType, kubeconfig string, obj runtime.Object, opts ...kubernetes.KubectlGetOption) error {
	_ = "STUB: not implemented"
	return nil
}

func getParams(resourceType, kubeconfig string, o *kubernetes.KubectlGetOptions) []string {
	_ = "STUB: not implemented"
	return nil
}

// Create performs a kubectl create command.
func (k *Kubectl) Create(ctx context.Context, kubeconfig string, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

const alreadyExistsErrorMessageSubString = "AlreadyExists"

func isKubectlAlreadyExistsError(err error) bool { _ = "STUB: not implemented"; return false }

const notFoundErrorMessageSubString = "NotFound"

// IsKubectlNotFoundError returns true if the kubectl call returned the NotFound error.
func IsKubectlNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

func newAlreadyExistsErrorForObj(obj runtime.Object) error { _ = "STUB: not implemented"; return nil }

func groupResourceFromObj(obj runtime.Object) schema.GroupResource {
	_ = "STUB: not implemented"
	return *new(schema.GroupResource)
}

// If this doesn't implement the client object interface,
// we don't know how to process it. This should never happen for
// any of the known types.

func resourceNameFromObj(obj runtime.Object) string { _ = "STUB: not implemented"; return "" }

// If this doesn't implement the client object interface,
// we don't know how to process it. This should never happen for
// any of the known types.

func newNotFoundErrorForTypeAndName(resourceType, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Replace performs a kubectl replace command.
func (k *Kubectl) Replace(ctx context.Context, kubeconfig string, obj runtime.Object) error {
	_ = "STUB: not implemented"
	// Even is --save-config=false is set (which is the default), kubectl replace will
	// not only respect the last-applied annotation if present in the object, but it will update
	// it with the provided state of the resource. This includes the metadata.resourceVersion. This
	// breaks future uses of kubectl apply. Since those commands' input never provide the resourceVersion,
	// kubectl will send a request trying to remove that field. That is obviously not a valid request, so
	// it gets rejected by the kube API server. To avoid this, we simply remove the annotation when passing
	// it to the replace command.
	// It's not recommended to use both imperative and "declarative" commands for the same resource. Unfortunately
	// our CLI makes extensive use of client side apply. Although not ideal, this mechanism allows us to perform
	// updates (using replace) where idempotency is necessary while maintaining the ability to continue to use apply.
	return nil
}

// removeLastAppliedAnnotation deletes the kubectl last-applied annotation
// from the object if present.
func removeLastAppliedAnnotation(obj runtime.Object) runtime.Object {
	_ = "STUB: not implemented"
	return *new(runtime.Object)
}

// If this doesn't implement the client object interface,
// we don't know how to access the annotations.
// All the objects that we pass here do implement client.Client.

// Delete performs a delete command authenticating with a kubeconfig file.
func (k *Kubectl) Delete(ctx context.Context, resourceType, kubeconfig string, opts ...kubernetes.KubectlDeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteParams(resourceType, kubeconfig string, o *kubernetes.KubectlDeleteOptions) []string {
	_ = "STUB: not implemented"
	return nil
}

// Apply creates the resource or it updates if it already exists.
func (k *Kubectl) Apply(ctx context.Context, kubeconfig string, obj runtime.Object, opts ...kubernetes.KubectlApplyOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) GetEksdRelease(ctx context.Context, name, namespace, kubeconfigFile string) (*eksdv1alpha1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetDeployment(ctx context.Context, name, namespace, kubeconfig string) (*appsv1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetDaemonSet(ctx context.Context, name, namespace, kubeconfig string) (*appsv1.DaemonSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) ExecuteCommand(ctx context.Context, opts ...string) (bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

// DeleteClusterObject performs a DELETE call like above except without namespace required.
func (k *Kubectl) DeleteClusterObject(ctx context.Context, resourceType, name, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) ExecuteFromYaml(ctx context.Context, yaml []byte, opts ...string) (bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (k *Kubectl) SearchNutanixMachineConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.NutanixMachineConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) SearchNutanixDatacenterConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.NutanixDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaNutanixDatacenterConfig(ctx context.Context, nutanixDatacenterConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.NutanixDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) GetEksaNutanixMachineConfig(ctx context.Context, nutanixMachineConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.NutanixMachineConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kubectl) DeleteEksaNutanixDatacenterConfig(ctx context.Context, nutanixDatacenterConfigName string, kubeconfigFile string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubectl) DeleteEksaNutanixMachineConfig(ctx context.Context, nutanixMachineConfigName string, kubeconfigFile string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// AllBaseboardManagements returns all the baseboard management resources in the cluster.
func (k *Kubectl) AllBaseboardManagements(ctx context.Context, kubeconfig string) ([]rufiounreleased.BaseboardManagement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllTinkerbellHardware returns all the hardware resources in the cluster.
func (k *Kubectl) AllTinkerbellHardware(ctx context.Context, kubeconfig string) ([]tinkv1alpha1.Hardware, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HasCRD checks if the given CRD exists in the cluster specified by kubeconfig.
func (k *Kubectl) HasCRD(ctx context.Context, crd, kubeconfig string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// DeleteCRD removes the given CRD from the cluster specified in kubeconfig.
func (k *Kubectl) DeleteCRD(ctx context.Context, crd, kubeconfig string) error {
	_ = "STUB: not implemented"
	return nil
}
