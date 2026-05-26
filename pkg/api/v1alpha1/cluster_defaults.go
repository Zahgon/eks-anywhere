package v1alpha1

var clusterDefaults = []func(*Cluster) error{
	setRegistryMirrorConfigDefaults,
	setWorkerNodeGroupDefaults,
	setCNIConfigDefault,
	setEtcdEncryptionConfigDefaults,
}

func setClusterDefaults(cluster *Cluster) error { _ = "STUB: not implemented"; return nil }

func setRegistryMirrorConfigDefaults(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func setWorkerNodeGroupDefaults(cluster *Cluster) error { _ = "STUB: not implemented"; return nil }

func setCNIConfigDefault(cluster *Cluster) error { _ = "STUB: not implemented"; return nil }
