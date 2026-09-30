// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
// Command mutual-exclusion demonstrates protecting a shared counter.
package main

import (
	"fmt"
	"sync"

	"labone/semaphore"
)

const (
	workers    = 100
	increments = 1000
)

func main() {
	lock := semaphore.New(1)
	shared := 0
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				lock.Wait()
				shared++
				lock.Signal()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final shared value: %d (expected %d)\n", shared, workers*increments)
}
