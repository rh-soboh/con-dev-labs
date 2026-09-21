// Package semaphore provides a counting semaphore for coordinating goroutines.
//
// A Semaphore deliberately exposes only Wait and Signal. Its current count is
// private, so callers cannot depend on implementation details.
package semaphore

import (
	"sync"
	"time"
)

// Semaphore is a counting semaphore.
//
// A call to Wait consumes one permit and blocks when no permit is available.
// A call to Signal adds one permit and wakes one waiting goroutine, if any.
type Semaphore struct {
	mu      sync.Mutex
	permits int
	waiters int
	changed *sync.Cond
}

// New creates a semaphore with initial permits.
func New(initial int) *Semaphore {
	s := &Semaphore{permits: initial}
	s.changed = sync.NewCond(&s.mu)
	return s
}

// Wait acquires one permit, blocking until one is available.
func (s *Semaphore) Wait() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for s.permits == 0 {
		s.waiters++
		s.changed.Wait()
		s.waiters--
	}
	s.permits--
}

// TryWait waits for at most timeout to acquire one permit.
//
// It returns true when a permit was acquired and false when the timeout
// elapsed. A non-positive timeout performs an immediate, non-blocking try.
func (s *Semaphore) TryWait(timeout time.Duration) bool {
	s.mu.Lock()

	if s.permits > 0 {
		s.permits--
		s.mu.Unlock()
		return true
	}
	if timeout <= 0 {
		s.mu.Unlock()
		return false
	}

	timedOut := false
	go func() {
		time.Sleep(timeout)
		s.mu.Lock()
		if !timedOut {
			timedOut = true
			s.changed.Broadcast()
		}
		s.mu.Unlock()
	}()

	s.waiters++
	for s.permits == 0 && !timedOut {
		s.changed.Wait()
	}
	s.waiters--
	acquired := !timedOut
	if acquired {
		s.permits--
	}
	timedOut = true
	s.mu.Unlock()
	return acquired
}

// Signal releases one permit and wakes one waiting goroutine.
func (s *Semaphore) Signal() {
	s.mu.Lock()
	s.permits++
	if s.waiters > 0 {
		s.changed.Signal()
	}
	s.mu.Unlock()
}
