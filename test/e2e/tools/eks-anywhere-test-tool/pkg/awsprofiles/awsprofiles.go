package awsprofiles

type EksAccount int64

const (
	BuildAccount EksAccount = iota
	TestAccount
)

func (s EksAccount) ProfileName() string { _ = "STUB: not implemented"; return "" }
