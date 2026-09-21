// This code was written with help from Terry Huynh (C00300806).
package semaphore

import (
	"testing"
	"time"
)

func TestWaitBlocksUntilSignal(t *testing.T) {
	s := New(0)
	acquired := make(chan struct{})

	go func() {
		s.Wait()
		close(acquired)
	}()

	select {
	case <-acquired:
		t.Fatal("Wait returned before Signal")
	case <-time.After(20 * time.Millisecond):
	}

	s.Signal()
	select {
	case <-acquired:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return after Signal")
	}
}

func TestTryWaitTimesOut(t *testing.T) {
	s := New(0)
	if s.TryWait(20 * time.Millisecond) {
		t.Fatal("TryWait acquired a permit that was unavailable")
	}
}

func TestTryWaitAcquiresPermit(t *testing.T) {
	s := New(1)
	if !s.TryWait(time.Second) {
		t.Fatal("TryWait did not acquire an available permit")
	}
}
