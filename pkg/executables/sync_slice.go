package executables

import "sync"

type syncSlice struct {
	internal []string
	sync.RWMutex
}

func newSyncSlice() *syncSlice { _ = "STUB: not implemented"; return nil }

func (s *syncSlice) append(v ...string) { _ = "STUB: not implemented"; return }

func (s *syncSlice) iterate() <-chan string { _ = "STUB: not implemented"; return nil }
