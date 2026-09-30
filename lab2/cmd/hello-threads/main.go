// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
// Command hello-threads demonstrates starting and joining goroutines.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine one")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine two")
	}()

	fmt.Println("Launched from main")
	wg.Wait()
}
