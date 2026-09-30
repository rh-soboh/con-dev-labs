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
