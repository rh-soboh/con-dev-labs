// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"fmt"
	"sync"
	"time"
)

// channelBarrier is a one-use barrier backed by an unbuffered release channel.
type channelBarrier struct {
	mu      sync.Mutex
	total   int
	arrived int
	release chan struct{}
}

func newChannelBarrier(total int) *channelBarrier {
	if total <= 0 {
		panic("barrier total must be positive")
	}
	return &channelBarrier{
		total:   total,
		release: make(chan struct{}),
	}
}

func (b *channelBarrier) wait() {
	b.mu.Lock()
	b.arrived++
	if b.arrived == b.total {
		close(b.release)
	}
	b.mu.Unlock()

	<-b.release
}

func main() {
	const workers = 10
	barrier := newChannelBarrier(workers)
	var wg sync.WaitGroup
	wg.Add(workers)

	for id := 0; id < workers; id++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(workers-id) * 20 * time.Millisecond)
			fmt.Println("Part A", id)
			barrier.wait()
			fmt.Println("Part B", id)
		}(id)
	}

	wg.Wait()
}
