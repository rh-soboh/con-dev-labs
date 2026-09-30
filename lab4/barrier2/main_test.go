// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestChannelBarrierWaitsForEveryWorker(t *testing.T) {
	const workers = 5
	barrier := newChannelBarrier(workers)
	var arrived atomic.Int32
	var passed atomic.Int32
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			arrived.Add(1)
			barrier.wait()
			passed.Add(1)
		}()
	}

	wg.Wait()
	if arrived.Load() != workers || passed.Load() != workers {
		t.Fatalf("barrier released %d of %d workers", passed.Load(), workers)
	}
}
