package docker

import (
	"context"
	"sync"
)

type ConcurrentImageProcessor struct {
	maxRoutines int
}

func NewConcurrentImageProcessor(maxRoutines int) *ConcurrentImageProcessor {
	_ = "STUB: not implemented"
	return nil
}

type ImageProcessor func(ctx context.Context, image string) error

func (c *ConcurrentImageProcessor) Process(ctx context.Context, images []string, process ImageProcessor) error {
	_ = "STUB: not implemented"
	return nil
}

// This is not necessary because if we get here all workers are done, regardless if all the jobs
// were run or not, so canceling the context is not necessary since there is nothing else using it
// This is just to avoid a lint warning for possible context leeking

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }

func max(a, b int) int { _ = "STUB: not implemented"; return 0 }

type job struct {
	ctx   context.Context
	image string
}

type feeder struct {
	jobs            chan<- job
	images          []string
	done            chan<- struct{}
	workersWaiGroup *sync.WaitGroup
}

func (f *feeder) feed(ctx context.Context) { _ = "STUB: not implemented"; return }

type worker struct {
	jobs        <-chan job
	process     ImageProcessor
	waitGroup   *sync.WaitGroup
	errorReturn chan<- error
}

func (w *worker) start() { _ = "STUB: not implemented"; return }
