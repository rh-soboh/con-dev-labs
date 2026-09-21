// Barrier example for Concurrent Development Lab 2.
// Author: Shadrach Oboh
// Licence: MIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// Barrier blocks callers until the configured number of goroutines arrive.
type Barrier struct {
	mu      sync.Mutex
	count   int
	waiters int
	ready   chan struct{}
}

// NewBarrier creates a one-use barrier for count goroutines.
func NewBarrier(count int) *Barrier {
	if count <= 0 {
		panic("barrier count must be positive")
	}

	return &Barrier{
		count: count,
		ready: make(chan struct{}),
	}
}

// Wait records an arriving goroutine and blocks until all goroutines arrive.
func (b *Barrier) Wait() {
	b.mu.Lock()
	b.waiters++
	if b.waiters == b.count {
		close(b.ready)
	}
	ready := b.ready
	b.mu.Unlock()

	<-ready
}

func main() {
	const workers = 5
	barrier := NewBarrier(workers)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(workers-id) * 100 * time.Millisecond)
			fmt.Printf("Part A %d\n", id)
			barrier.Wait()
			fmt.Printf("Part B %d\n", id)
		}(i)
	}

	wg.Wait()
}
