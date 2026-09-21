// This code was written with help from Terry Huynh (C00300806).
// Rendezvous example for Concurrent Development Lab 2.
// Author: Shadrach Oboh
// Licence: MIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// worker performs Part A, waits for the other worker, and then performs Part B.
func worker(name string, arrived chan<- struct{}, release <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("%s: Part A\n", name)
	arrived <- struct{}{}
	<-release
	fmt.Printf("%s: Part B\n", name)
}

func main() {
	firstArrived := make(chan struct{})
	secondArrived := make(chan struct{})
	firstRelease := make(chan struct{})
	secondRelease := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go worker("First", firstArrived, firstRelease, &wg)
	go worker("Second", secondArrived, secondRelease, &wg)

	<-firstArrived
	<-secondArrived
	close(firstRelease)
	close(secondRelease)

	wg.Wait()
}
