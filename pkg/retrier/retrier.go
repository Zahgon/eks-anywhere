package retrier

import (
	"time"
)

type Retrier struct {
	retryPolicy   RetryPolicy
	timeout       time.Duration
	backoffFactor *float32
}

type (
	// RetryPolicy allows to customize the retrying logic. The boolean retry indicates if a new retry
	// should be performed and the wait duration indicates the wait time before the next retry.
	RetryPolicy func(totalRetries int, err error) (retry bool, wait time.Duration)
	RetrierOpt  func(*Retrier)
)

// New creates a new retrier with a global timeout (max time allowed for the whole execution)
// The default retry policy is to always retry with no wait time in between retries.
func New(timeout time.Duration, opts ...RetrierOpt) *Retrier { _ = "STUB: not implemented"; return nil }

// NewWithMaxRetries creates a new retrier with no global timeout and a max retries policy.
func NewWithMaxRetries(maxRetries int, backOffPeriod time.Duration) *Retrier {
	_ = "STUB: not implemented"
	// this value is roughly 292 years, so in practice there is no timeout
	return nil
}

// NewWithNoTimeout creates a new retrier with no global timeout and infinite retries.
func NewWithNoTimeout() *Retrier { _ = "STUB: not implemented"; return nil }

// WithMaxRetries sets a retry policy that will retry up to maxRetries times
// with a wait time between retries of backOffPeriod.
func WithMaxRetries(maxRetries int, backOffPeriod time.Duration) RetrierOpt {
	_ = "STUB: not implemented"
	return *new(RetrierOpt)
}

func WithBackoffFactor(factor float32) RetrierOpt {
	_ = "STUB: not implemented"
	return *new(RetrierOpt)
}

func WithRetryPolicy(policy RetryPolicy) RetrierOpt {
	_ = "STUB: not implemented"
	return *new(RetrierOpt)
}

// Retry runs the fn function until it either successful completes (not error),
// the set timeout reached or the retry policy aborts the execution.
func (r *Retrier) Retry(fn func() error) error {
	_ = "STUB: not implemented"
	// While it seems aberrant to call a method with a nil receiver, several unit tests actually do.  With a previous
	// version of this module (which didn't attempt to dereference the receiver until after the wrapped function failed)
	// these passed.  Changes below, to log the receiver struct's key params changed that breaking the unit tests.
	// The below conditional block restores the original behavior, enabling these tests to again pass.
	return nil
}

// If there's not enough time left for the policy-proposed wait, there's no value in waiting that duration
// before quitting at the bottom of the loop.  Just do it now.

// Retry runs fn with a MaxRetriesPolicy.
func Retry(maxRetries int, backOffPeriod time.Duration, fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

// BackOffPolicy retries until top level timeout is reached, waiting a
// backoff period in between retries.
func BackOffPolicy(backoff time.Duration) RetryPolicy {
	_ = "STUB: not implemented"
	return *new(RetryPolicy)
}

func zeroWaitPolicy(_ int, _ error) (retry bool, wait time.Duration) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration)
}

func maxRetriesPolicy(maxRetries int, backOffPeriod time.Duration) RetryPolicy {
	_ = "STUB: not implemented"
	return *new(RetryPolicy)
}
