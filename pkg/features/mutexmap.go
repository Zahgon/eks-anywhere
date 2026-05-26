package features

import "sync"

func newMutexMap() *mutexMap { _ = "STUB: not implemented"; return nil }

type mutexMap struct {
	internal map[string]bool
	sync.RWMutex
}

func (m *mutexMap) load(key string) (value bool, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (m *mutexMap) store(key string, value bool) { _ = "STUB: not implemented"; return }

func (m *mutexMap) clear() { _ = "STUB: not implemented"; return }
