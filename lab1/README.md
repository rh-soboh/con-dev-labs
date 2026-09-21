# Go Concurrency Lab One

This project is a Go implementation of the semaphore, rendezvous, and mutual
exclusion examples supplied in the C++ lab template.

## Author and licence

Author: Shadrach Oboh  
Licence: [MIT License](LICENSE)

## Requirements

- Go 1.21 or newer
- Doxygen (optional, for generated API documentation)

## Run

From the project root:

```sh
go test ./...
go run ./cmd/hello-threads
go run ./cmd/rendezvous
go run ./cmd/mutual-exclusion
go test -race ./...
```

The mutual exclusion example should always print `100000`, because the
semaphore permits only one goroutine at a time to update the shared counter.

## Generate documentation

Go source comments use GoDoc conventions and can be viewed with:

```sh
go doc ./semaphore
```

Doxygen can also generate HTML documentation from the source comments:

```sh
doxygen Doxyfile
```

The generated documentation is written to `docs/html`.

## Files

- `semaphore/semaphore.go`: reusable counting semaphore implementation.
- `semaphore/semaphore_test.go`: semaphore behavior tests.
- `cmd/hello-threads`: basic goroutine and wait-group example.
- `cmd/rendezvous`: two-way semaphore rendezvous.
- `cmd/mutual-exclusion`: semaphore-protected shared counter.
- `Doxyfile`: optional Doxygen configuration.
- `LICENSE`: MIT licence.

## To do

- Add further semaphore exercises from later lab sessions.
