// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"sync"
	"testing"
)

func TestBarrierWaitsForEveryWorker(t *testing.T) {
	const workers = 4
	barrier := newBarrier(workers)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			barrier.wait()
		}()
	}

	wg.Wait()
}
