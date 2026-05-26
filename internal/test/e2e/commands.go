package e2e

type copyCommand string

func newCopyCommand() copyCommand { _ = "STUB: not implemented"; return *new(copyCommand) }

func (c copyCommand) String() string { _ = "STUB: not implemented"; return "" }

func (c copyCommand) addOption(opt string) copyCommand {
	_ = "STUB: not implemented"
	return *new(copyCommand)
}

func (c copyCommand) from(p ...string) copyCommand {
	_ = "STUB: not implemented"
	return *new(copyCommand)
}

func (c copyCommand) to(p ...string) copyCommand {
	_ = "STUB: not implemented"
	return *new(copyCommand)
}

func (c copyCommand) recursive() copyCommand { _ = "STUB: not implemented"; return *new(copyCommand) }

func (c copyCommand) exclude(v string) copyCommand {
	_ = "STUB: not implemented"
	return *new(copyCommand)
}

func (c copyCommand) include(v string) copyCommand {
	_ = "STUB: not implemented"
	return *new(copyCommand)
}
