package features

import (
	"sync"
)

var globalFeatures = newFeatures()

type features struct {
	cache     *mutexMap
	gates     map[string]string
	initGates sync.Once
}

func newFeatures() *features { _ = "STUB: not implemented"; return nil }

func (f *features) feedGates(featureGates []string) { _ = "STUB: not implemented"; return }

func (f *features) isActiveForEnvVar(envVar string) func() bool {
	_ = "STUB: not implemented"
	return nil
}

func (f *features) isActiveForEnvVarOrGate(envVar, gateName string) func() bool {
	_ = "STUB: not implemented"
	return nil
}

func (f *features) clearCache() { _ = "STUB: not implemented"; return }
