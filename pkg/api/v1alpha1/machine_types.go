package v1alpha1

type OSFamily string

const (
	Ubuntu       OSFamily = "ubuntu"
	Bottlerocket OSFamily = "bottlerocket"
	RedHat       OSFamily = "redhat"
)

// UserConfiguration defines the configuration of the user to be added to the VM.
type UserConfiguration struct {
	Name              string   `json:"name"`
	SshAuthorizedKeys []string `json:"sshAuthorizedKeys"`
}

func defaultMachineConfigUsers(defaultUsername string, users []UserConfiguration) []UserConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func validateMachineConfigUsers(machineConfigName string, machineConfigKind string, users []UserConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}
