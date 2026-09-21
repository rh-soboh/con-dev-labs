// This code was written with help from Terry Huynh (C00300806).
// Command rendezvous demonstrates a two-way semaphore rendezvous.
package main

import (
	"fmt"
	"sync"
	"time"

	"labone/semaphore"
)

func main() {
	aArrived := semaphore.New(0)
	bArrived := semaphore.New(0)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("A1: task A has arrived")
		aArrived.Signal()
		bArrived.Wait()
		fmt.Println("A2: task A continues after task B")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("B1: task B has arrived")
		bArrived.Signal()
		aArrived.Wait()
		fmt.Println("B2: task B continues after task A")
	}()

	wg.Wait()
}
