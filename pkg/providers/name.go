package providers

const (
	EtcdNodeNameSuffix         = "etcd"
	ControlPlaneNodeNameSuffix = "cp"
)

func GetControlPlaneNodeName(clusterName string) string { _ = "STUB: not implemented"; return "" }

func GetEtcdNodeName(clusterName string) string { _ = "STUB: not implemented"; return "" }
