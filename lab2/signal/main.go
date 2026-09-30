// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arrived := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("First - Part A")
		arrived <- struct{}{}
		fmt.Println("First - Part B")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Second - Part A")
		<-arrived
		fmt.Println("Second - Part B")
	}()

	wg.Wait()
}
