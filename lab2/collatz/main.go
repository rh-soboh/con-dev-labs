// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"fmt"
	"sync"
)

const (
	collatzInputs = 64
	collatzWorkers = 4
)

func collatzSteps(value int) int {
	if value <= 0 {
		panic("collatz input must be positive")
	}

	steps := 0
	for value > 1 {
		if value%2 == 0 {
			value /= 2
		} else {
			value = 3*value + 1
		}
		steps++
	}
	return steps
}

func main() {
	results := make([]int, collatzInputs)
	tokens := make(chan struct{}, collatzWorkers)
	var wg sync.WaitGroup

	for input := 1; input <= collatzInputs; input++ {
		index := input - 1
		tokens <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-tokens }()
			results[index] = collatzSteps(index + 1)
		}()
	}

	wg.Wait()
	fmt.Println(results)
}
