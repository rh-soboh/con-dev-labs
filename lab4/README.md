# Go Concurrency Lab Four

This lab contains two working barrier implementations based on the supplied
barrier templates.

## Run

From the repository root:

```sh
go test ./lab4/...
go run ./lab4/barrier2
go run ./lab4/barrier-struct
go test -race ./lab4/...
```

`barrier2` demonstrates a one-use barrier using a mutex and a channel.
`barrier-struct` demonstrates the same barrier idea stored in a dedicated
barrier type, using a mutex and an unbuffered channel.

## Generate documentation

From the `lab4` directory:

```sh
doxygen Doxyfile
```

The generated HTML is written to `lab4/docs/html`.
