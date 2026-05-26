package aflag

import (
	"net/http"
)

// Header Value.
type Header http.Header

// NewHeader returns a new Header pointer.
func NewHeader(h *http.Header) *Header { _ = "STUB: not implemented"; return nil }

// String returns the string representation of the Header.
func (h *Header) String() string { _ = "STUB: not implemented"; return "" }

// Set sets the value of the Header.
// Format: "a=1;2,b=2;4;5".
func (h *Header) Set(val string) error { _ = "STUB: not implemented"; return nil }

// Type returns the flag type.
func (h *Header) Type() string { _ = "STUB: not implemented"; return "" }
