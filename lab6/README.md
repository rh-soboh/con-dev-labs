# Go Concurrency Lab Six

This lab demonstrates the producer-consumer problem from the supplied C++
example. Producer goroutines create events and put them into a bounded,
concurrency-safe buffer. Consumer goroutines take events from the buffer and
process them. A buffered Go channel provides the buffer's synchronization and
capacity limit.

## Author and licence

Author: Shadrach Oboh  
Licence: [MIT License](LICENSE)

## Requirements

- Go 1.21 or newer
- Doxygen (optional, for generated API documentation)

## Run

From the repository root:

```sh
go run ./lab6/producer-consumer
```

The program starts two producers and two consumers. Each producer creates five
events, and the buffer can hold three events at a time. The program finishes
after all ten events have been consumed.

## Generate documentation

From the `lab6` directory:

```sh
doxygen Doxyfile
```

The generated HTML is written to `lab6/docs/html`.

## Files

- `producer-consumer/main.go`: event type, bounded buffer, producers,
  consumers, and program entry point.
- `Doxyfile`: optional Doxygen configuration.
- `LICENSE`: MIT licence.

## To do

- Nothing currently planned.
