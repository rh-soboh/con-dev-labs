// Package main demonstrates producer-consumer with goroutines and a bounded
// channel buffer.
package main

import (
	"fmt"
	"sync"
)

const (
	producerCount     = 2
	consumerCount     = 2
	eventsPerProducer = 5
	bufferSize        = 3
)

// Event is a unit of work created by a producer and handled by a consumer.
type Event struct {
	ID         int
	ProducerID int
}

// SafeBuffer is a bounded, concurrency-safe buffer. Putting into a full
// buffer waits for a consumer; getting from an empty buffer waits for a
// producer.
type SafeBuffer struct {
	events chan Event
}

func newSafeBuffer(size int) *SafeBuffer {
	return &SafeBuffer{events: make(chan Event, size)}
}

func (b *SafeBuffer) put(event Event) {
	b.events <- event
}

func (b *SafeBuffer) get() (Event, bool) {
	event, ok := <-b.events
	return event, ok
}

func producer(id int, buffer *SafeBuffer, done *sync.WaitGroup) {
	defer done.Done()

	for eventID := 1; eventID <= eventsPerProducer; eventID++ {
		event := Event{ID: eventID, ProducerID: id}
		fmt.Printf("Producer %d produced event %d\n", id, eventID)
		buffer.put(event)
	}
}

func consumer(id int, buffer *SafeBuffer, done *sync.WaitGroup) {
	defer done.Done()

	for {
		event, ok := buffer.get()
		if !ok {
			return
		}
		fmt.Printf("Consumer %d consumed event %d from producer %d\n",
			id, event.ID, event.ProducerID)
	}
}

func main() {
	buffer := newSafeBuffer(bufferSize)
	var producers sync.WaitGroup
	var consumers sync.WaitGroup

	consumers.Add(consumerCount)
	for id := 1; id <= consumerCount; id++ {
		go consumer(id, buffer, &consumers)
	}

	producers.Add(producerCount)
	for id := 1; id <= producerCount; id++ {
		go producer(id, buffer, &producers)
	}

	producers.Wait()
	close(buffer.events)
	consumers.Wait()

	fmt.Printf("Finished: %d events produced and consumed\n",
		producerCount*eventsPerProducer)
}
