package test

import (
	"github.com/go-logr/logr"
)

func NewNullLogger() logr.Logger { _ = "STUB: not implemented"; return *new(logr.Logger) }
