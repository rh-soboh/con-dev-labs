// This code was written with help from Terry Huynh (C00300806) and Isabel Rafter (C00303465).
// I also helped Terry Huynh and Isabel Rafter with their code.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	philosopherCount = 5
	mealRounds       = 3
)

type table struct {
	forks []chan struct{}
	room  chan struct{}
}

func newTable(count int) *table {
	if count < 2 {
		panic("at least two philosophers are required")
	}

	result := &table{
		forks: make([]chan struct{}, count),
		room:  make(chan struct{}, count-1),
	}
	for index := range result.forks {
		result.forks[index] = make(chan struct{}, 1)
		result.forks[index] <- struct{}{}
	}
	return result
}

func (t *table) getForks(index int) {
	t.room <- struct{}{}
	<-t.forks[index]
	<-t.forks[(index+1)%len(t.forks)]
}

func (t *table) putForks(index int) {
	t.forks[(index+1)%len(t.forks)] <- struct{}{}
	t.forks[index] <- struct{}{}
	<-t.room
}

func philosopher(index, rounds int, t *table, completed *atomic.Int32) {
	for round := 0; round < rounds; round++ {
		time.Sleep(time.Duration((index+round)%3) * 10 * time.Millisecond)
		fmt.Printf("Philosopher %d is thinking\n", index)
		t.getForks(index)
		fmt.Printf("Philosopher %d is eating\n", index)
		time.Sleep(5 * time.Millisecond)
		t.putForks(index)
		completed.Add(1)
	}
}

func main() {
	table := newTable(philosopherCount)
	var completed atomic.Int32
	var wg sync.WaitGroup
	wg.Add(philosopherCount)

	for index := 0; index < philosopherCount; index++ {
		go func(index int) {
			defer wg.Done()
			philosopher(index, mealRounds, table, &completed)
		}(index)
	}

	wg.Wait()
	fmt.Printf("Completed meals: %d\n", completed.Load())
}
