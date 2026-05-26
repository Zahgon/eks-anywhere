package executables

type cmkCommandArgs func(*[]string)

func newCmkCommand(command string) []string { _ = "STUB: not implemented"; return nil }

func applyCmkArgs(params *[]string, args ...cmkCommandArgs) { _ = "STUB: not implemented"; return }

func appendArgs(new ...string) cmkCommandArgs {
	_ = "STUB: not implemented"
	return *new(cmkCommandArgs)
}

func withCloudStackDomainId(domainId string) cmkCommandArgs {
	_ = "STUB: not implemented"
	return *new(cmkCommandArgs)
}

func withCloudStackAccount(account string) cmkCommandArgs {
	_ = "STUB: not implemented"
	return *new(cmkCommandArgs)
}

func withCloudStackZoneId(zoneId string) cmkCommandArgs {
	_ = "STUB: not implemented"
	return *new(cmkCommandArgs)
}

func withCloudStackId(id string) cmkCommandArgs {
	_ = "STUB: not implemented"
	return *new(cmkCommandArgs)
}

func withCloudStackName(name string) cmkCommandArgs {
	_ = "STUB: not implemented"
	return *new(cmkCommandArgs)
}

func withCloudStackKeyword(keyword string) cmkCommandArgs {
	_ = "STUB: not implemented"
	return *new(cmkCommandArgs)
}
