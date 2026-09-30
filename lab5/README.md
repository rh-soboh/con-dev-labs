# Go Concurrency Lab Five

This lab implements the Dining Philosophers problem from the supplied C++
and Go demonstrations.

## Run

From the repository root:

```sh
go test ./lab5/...
go run ./lab5/philosophers
go test -race ./lab5/...
```

The philosophers use one buffered channel per fork. A room channel allows at
most four of the five philosophers to try to acquire forks at once, preventing
the circular wait that causes the template program to deadlock. The demo runs
three finite meal rounds so it can finish and be tested.

## Generate documentation

From the `lab5` directory:

```sh
doxygen Doxyfile
```

The generated HTML is written to `lab5/docs/html`.
