package vsphere

type OVFDeployOptions struct {
	Name             string           `json:"Name"`
	PowerOn          bool             `json:"PowerOn"`
	DiskProvisioning string           `json:"DiskProvisioning"`
	WaitForIP        bool             `json:"WaitForIP"`
	NetworkMappings  []NetworkMapping `json:"NetworkMapping"`
	Annotation       string           `json:"Annotation"`
	PropertyMapping  []OVFProperty    `json:"PropertyMapping"`
	InjectOvfEnv     bool             `json:"InjectOvfEnv"`
}

type OVFProperty struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type NetworkMapping struct {
	Name    string `json:"Name"`
	Network string `json:"Network"`
}

func DeployTemplate(envMap map[string]string, library, templateName, vmName, deployFolder, datacenter, datastore, resourcePool string, opts OVFDeployOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// deploy template

func TagVirtualMachine(envMap map[string]string, vmPath, tag string) error {
	_ = "STUB: not implemented"
	return nil
}
