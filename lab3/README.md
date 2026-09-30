# Go Concurrency Lab Three

This lab implements the Barrier and Rendezvous exercises from the
Concurrent Development laboratory.

## Author and licence

Author: Shadrach Oboh
Licence: [MIT License](LICENSE)

## Requirements

- Go 1.21 or newer
- Run the commands from the repository root, where `go.mod` is located.

## Run

From the repository root:

```sh
go run ./lab3/barrier
go run ./lab3/rendezvous
```

The barrier program prints `Part A` for every worker before any `Part B`
message is printed. The rendezvous program waits for both workers to reach
their rendezvous before allowing either worker to continue.

## Generate documentation

Install Doxygen if it is not already available, then run this command from
the `lab3` directory:

```sh
doxygen Doxyfile
```

The generated HTML documentation is written to `lab3/docs/html`.

## Files

- `barrier/main.go`: one-use barrier implementation and demonstration.
- `rendezvous/main.go`: two-party rendezvous implementation and demonstration.
- `Doxyfile`: Doxygen configuration.
- `LICENSE`: MIT licence.

## To do

- Add automated tests for the barrier and rendezvous synchronization guarantees.
