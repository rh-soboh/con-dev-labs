// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPhilosophersFinishWithoutDeadlock(t *testing.T) {
	const rounds = 20
	table := newTable(philosopherCount)
	var completed atomic.Int32
	var wg sync.WaitGroup
	wg.Add(philosopherCount)

	for index := 0; index < philosopherCount; index++ {
		go func(index int) {
			defer wg.Done()
			philosopher(index, rounds, table, &completed)
		}(index)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("philosophers did not finish; possible deadlock")
	}

	expected := int32(philosopherCount * rounds)
	if completed.Load() != expected {
		t.Fatalf("completed %d meals, expected %d", completed.Load(), expected)
	}
}
