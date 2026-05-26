package types

type ChangeDiff struct {
	ComponentReports []ComponentChangeDiff `json:"components"`
}

type ComponentChangeDiff struct {
	ComponentName string `json:"name"`
	OldVersion    string `json:"oldVersion"`
	NewVersion    string `json:"newVersion"`
}

func NewChangeDiff(componentReports ...*ComponentChangeDiff) *ChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

func (c *ChangeDiff) Append(changeDiffs ...*ChangeDiff) { _ = "STUB: not implemented"; return }

func (c *ChangeDiff) Changed() bool { _ = "STUB: not implemented"; return false }
