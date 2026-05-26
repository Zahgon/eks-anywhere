package validations

func ValidateClusterNameArg(args []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func FileExists(filename string) bool { _ = "STUB: not implemented"; return false }

func FileExistsAndIsNotEmpty(filename string) bool { _ = "STUB: not implemented"; return false }

// ValidateClusterNameFromCommandAndConfig validates if cluster name provided in command matches with cluster name in config file.
func ValidateClusterNameFromCommandAndConfig(args []string, clusterNameConfig string) error {
	_ = "STUB: not implemented"
	return nil
}
