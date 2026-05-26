package test

import (
	"github.com/golang/mock/gomock"
)

type ofType struct{ t string }

func OfType(t string) gomock.Matcher { _ = "STUB: not implemented"; return *new(gomock.Matcher) }

func (o *ofType) Matches(x interface{}) bool { _ = "STUB: not implemented"; return false }

func (o *ofType) String() string { _ = "STUB: not implemented"; return "" }

// AContext returns a gomock matchers that evaluates if the receive value can
// fullfills the context.Context interface.
func AContext() gomock.Matcher { _ = "STUB: not implemented"; return *new(gomock.Matcher) }

type bytesMatchFile struct{ file string }

// MatchFile returns a gomock matcher that compares []byte input to the content
// of the given file.
func MatchFile(file string) gomock.Matcher { _ = "STUB: not implemented"; return *new(gomock.Matcher) }

func (o *bytesMatchFile) Matches(x interface{}) bool { _ = "STUB: not implemented"; return false }

func (o *bytesMatchFile) String() string { _ = "STUB: not implemented"; return "" }

func (o *bytesMatchFile) Got(got interface{}) string { _ = "STUB: not implemented"; return "" }
