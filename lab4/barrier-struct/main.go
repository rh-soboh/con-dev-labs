// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"fmt"
	"sync"
	"time"
)

// barrier is a one-use barrier built from a mutex and an unbuffered channel.
type barrier struct {
	theChan chan bool
	theLock sync.Mutex
	total   int
	count   int
}

func newBarrier(total int) *barrier {
	if total <= 0 {
		panic("barrier total must be positive")
	}

	return &barrier{
		theChan: make(chan bool),
		total:   total,
	}
}

func (b *barrier) wait() {
	b.theLock.Lock()
	b.count++
	if b.count == b.total {
		b.theLock.Unlock()
		for i := 0; i < b.total-1; i++ {
			<-b.theChan
		}
		return
	}
	b.theLock.Unlock()
	b.theChan <- true
}

func main() {
	const workers = 5
	barrier := newBarrier(workers)
	var wg sync.WaitGroup
	wg.Add(workers)

	for id := 0; id < workers; id++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration((workers-id)%workers) * 20 * time.Millisecond)
			fmt.Println("Part A", id)
			barrier.wait()
			fmt.Println("Part B", id)
		}(id)
	}

	wg.Wait()
}
