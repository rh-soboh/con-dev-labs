// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	atomicWorkers    = 10
	atomicIncrements = 1000
)

func main() {
	var total atomic.Int64
	var wg sync.WaitGroup
	wg.Add(atomicWorkers)

	for worker := 0; worker < atomicWorkers; worker++ {
		go func() {
			defer wg.Done()
			for increment := 0; increment < atomicIncrements; increment++ {
				total.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final atomic value: %d (expected %d)\n", total.Load(), atomicWorkers*atomicIncrements)
}
