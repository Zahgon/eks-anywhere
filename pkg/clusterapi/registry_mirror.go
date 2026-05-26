package clusterapi

import (
	_ "embed"

	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/registrymirror"
)

//go:embed config/containerd_config_append.toml
var containerdConfig string

//go:embed config/hosts.toml
var hostsTemplate string

// SetRegistryMirrorInKubeadmControlPlaneForBottlerocket sets up registry mirror configuration in kubeadmControlPlane for bottlerocket.
func SetRegistryMirrorInKubeadmControlPlaneForBottlerocket(kcp *controlplanev1beta2.KubeadmControlPlane, mirrorConfig *v1alpha1.RegistryMirrorConfiguration) {
	_ = "STUB: not implemented"
	return
}

// SetRegistryMirrorInKubeadmControlPlaneForUbuntu sets up registry mirror configuration in kubeadmControlPlane for ubuntu.
func SetRegistryMirrorInKubeadmControlPlaneForUbuntu(kcp *controlplanev1beta2.KubeadmControlPlane, mirrorConfig *v1alpha1.RegistryMirrorConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

// SetRegistryMirrorInKubeadmConfigTemplateForBottlerocket sets up registry mirror configuration in kubeadmConfigTemplate for bottlerocket.
func SetRegistryMirrorInKubeadmConfigTemplateForBottlerocket(kct *bootstrapv1beta2.KubeadmConfigTemplate, mirrorConfig *v1alpha1.RegistryMirrorConfiguration) {
	_ = "STUB: not implemented"
	return
}

// SetRegistryMirrorInKubeadmConfigTemplateForUbuntu sets up registry mirror configuration in kubeadmConfigTemplate for ubuntu.
func SetRegistryMirrorInKubeadmConfigTemplateForUbuntu(kct *bootstrapv1beta2.KubeadmConfigTemplate, mirrorConfig *v1alpha1.RegistryMirrorConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

// setRegistryMirrorInEtcdCluster sets up registry mirror configuration in etcdadmCluster.
func setRegistryMirrorInEtcdCluster(etcd *etcdv1.EtcdadmCluster, mirrorConfig *v1alpha1.RegistryMirrorConfiguration) {
	_ = "STUB: not implemented"
	return
}

func registryMirror(mirrorConfig *v1alpha1.RegistryMirrorConfiguration) bootstrapv1beta2.RegistryMirrorConfiguration {
	_ = "STUB: not implemented"
	return *new(bootstrapv1beta2.RegistryMirrorConfiguration)
}

type values map[string]interface{}

func registryMirrorConfigContent(registryMirror *registrymirror.RegistryMirror) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func hostsFileContent(registryMirror *registrymirror.RegistryMirror, server, host string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func registryMirrorConfig(registryMirrorConfig *v1alpha1.RegistryMirrorConfiguration) (files []bootstrapv1beta2.File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Main config file

// CA certificate if present

// Mirror base hosts.toml

// Individual registry hosts.toml files

func addRegistryMirrorInKubeadmConfigSpecFiles(kcs *bootstrapv1beta2.KubeadmConfigSpec, mirrorConfig *v1alpha1.RegistryMirrorConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}
