package v1alpha1

const (
	bottlerocketBootconfig = `kernel {}`

	cloudInit = `datasource:
  Ec2:
    metadata_urls: [%s]
    strict_id: false
manage_etc_hosts: localhost
warnings:
  dsid_missing_source: off
`

	// HookOS embeds container images from the bundle.
	// The container images are tagged as below.
	actionImage2Disk = "127.0.0.1/embedded/image2disk"
	actionWriteFile  = "127.0.0.1/embedded/writefile"
	actionReboot     = "127.0.0.1/embedded/reboot"
)

// DefaultActions constructs a set of default actions for the given osFamily.
func DefaultActions(clusterSpec *Cluster, osImageOverride, tinkerbellLocalIP, tinkerbellLBIP string, osFamily OSFamily) []ActionOpt {
	_ = "STUB: not implemented"
	// The metadata string will have two URLs:
	// 1. one that will be used initially for bootstrap and will point to tootles running on kind.
	// 2. one that will be used when the workload cluster is up and will point to tootles running on
	//    the workload cluster.
	// Port 7172 is the tootles (metadata service) port in the mono-repo tinkerbell chart.
	return nil
}

// During workflow reconciliation when the Tinkerbell template is rendered, the Workflow
// Controller injects a subset of data from the Hardware resource. This lets us use Go template
// language to render the disks enabling mix'n'match disk types for templates that represent
// the same kind of machine such as control plane nodes.
//
// The devicePath disk index and the storagePartitionPath disk index should match.

// Order matters. This action needs to append to an existing user-data.toml file so
// must be after withBottlerocketUserDataAction().

func withStreamImageAction(disk, imageURL string, additionalEnvVar map[string]string) ActionOpt {
	_ = "STUB: not implemented"
	return *new(ActionOpt)
}

func withNetplanAction(disk string, osFamily OSFamily) ActionOpt {
	_ = "STUB: not implemented"
	return *new(ActionOpt)
}

// Bottlerocket needs to write onto the 12th partition as opposed to 2nd for non-Bottlerocket OS

// For other OS families (Ubuntu, etc.), use netplan configuration

// Instead of using DHCP to get network information, use the hardware object directly

// Check if VLAN ID exists in the hardware object and create a VLAN-tagged interface if it does

// Use VLAN template if VLANID is present in the hardware object

func withDisableCloudInitNetworkCapabilities(disk string) ActionOpt {
	_ = "STUB: not implemented"
	return *new(ActionOpt)
}

func withTinkCloudInitAction(disk, metadataURLs string) ActionOpt {
	_ = "STUB: not implemented"
	return *new(ActionOpt)
}

func withDsCloudInitAction(disk string) ActionOpt {
	_ = "STUB: not implemented"
	return *new(ActionOpt)
}

func withRebootAction() ActionOpt { _ = "STUB: not implemented"; return *new(ActionOpt) }

func withBottlerocketBootconfigAction(disk string) ActionOpt {
	_ = "STUB: not implemented"
	return *new(ActionOpt)
}

func withBottlerocketUserDataAction(disk, metadataURLs string) ActionOpt {
	_ = "STUB: not implemented"
	return *new(ActionOpt)
}

func withNetworkManagerAction(disk string) ActionOpt {
	_ = "STUB: not implemented"
	return *new(ActionOpt)
}

// NetworkManager configuration template for RedHat (no VLAN)

// VLAN connection template for RedHat (single file approach)

// Use VLAN template if VLANID is present in the hardware object
