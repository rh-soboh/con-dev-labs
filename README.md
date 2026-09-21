# Concurrent Development Labs

This repository contains solutions for the Concurrent Development laboratory
sessions. Each laboratory is kept in its own directory with its own source
code, documentation, tests, and licence information.

## Labs

### [Lab 1](lab1/)

Go concurrency examples covering:

- Basic goroutines and `sync.WaitGroup`
- Semaphore-based rendezvous
- Semaphore-based mutual exclusion
- A reusable counting semaphore implementation
- Unit tests and race-detector testing
- GoDoc and generated Doxygen documentation

See the [Lab 1 README](lab1/README.md) for installation, usage, and
documentation instructions.

## Repository structure

```text
con-dev-labs/
├── lab1/
│   ├── cmd/
│   ├── semaphore/
│   ├── docs/
│   ├── Doxyfile
│   ├── go.mod
│   ├── LICENSE
│   └── README.md
└── README.md
```

Future lab sessions should be added as `lab2`, `lab3`, and so on. Each lab
should include its own README describing its requirements, how to run it, and
the files it contains.

## Licence

Individual labs contain their own licence files. Refer to the licence inside
each lab directory for the terms that apply to that lab.
