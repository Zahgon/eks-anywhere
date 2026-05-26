package curatedpackages

import (
	"io"
	"text/tabwriter"
)

// cpTabwriter is a modified tabwriter for curated packages CLI duty.
type cpTabwriter struct {
	*tabwriter.Writer
}

// newCPTabwriter instantiates a curated packages custom tabwriter.
//
// If customParams is nil, cpTabwriterDefaultParams will be used. The caller
// should call Flush just as they would with an unmodified tabwriter.Writer.
func newCPTabwriter(w io.Writer, customParams *cpTabwriterParams) *cpTabwriter {
	_ = "STUB: not implemented"
	return nil
}

// writeTable from a 2-D slice of strings, joining every string with tabs.
//
// Tab characters and newlines will be added to the end of each rank.
func (w *cpTabwriter) writeTable(lines [][]string) error { _ = "STUB: not implemented"; return nil }

// A final "\t" is added, as tabwriter is tab-terminated, not the more
// common tab-separated. See https://pkg.go.dev/text/tabwriter#Writer
// for details. There are cases where one might not want this trailing
// tab, but it hasn't come up yet, and is easily worked around when
// the time comes.

// cpTabwriterParams makes it easier to reuse common tabwriter parameters.
//
// See https://pkg.go.dev/text/tabwriter#Writer.Init for details.
type cpTabwriterParams struct {
	// minWidth is the minimal cell width including any padding
	minWidth int
	// tabWidth width of tab characters (equivalent number of spaces)
	tabWidth int
	// padding added to a cell before computing its width
	padding int
	// padChar ASCII char used for padding
	padChar byte
	// flags formatting control
	flags uint
}

// cpTabwriterDefaultParams is just a convenience when making tabwriters.
//
// Its implemented as a function to make it harder to override the defaults
// accidentally.
func cpTabwriterDefaultParams() *cpTabwriterParams { _ = "STUB: not implemented"; return nil }
