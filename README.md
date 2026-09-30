# Concurrent Development Labs

This repository contains solutions for the Concurrent Development laboratory
sessions. Each laboratory is kept in its own directory with its own source
code, documentation, and tests.

## Labs

### [Lab 1](lab1/)

Basic Go program used to create the first laboratory.

### [Lab 2](lab2/)

Go implementations of semaphore, rendezvous, and mutual exclusion
exercises. See the [Lab 2 README](lab2/README.md) for installation and usage
instructions.

### [Lab 3](lab3/)

Go implementations of the Barrier and Rendezvous synchronization exercises.
See the [Lab 3 README](lab3/README.md) for installation and usage
instructions.

### [Lab 4](lab4/)

Two barrier implementations based on the supplied barrier templates. See the
[Lab 4 README](lab4/README.md) for installation and usage instructions.

### [Lab 5](lab5/)

A deadlock-safe Dining Philosophers implementation based on the supplied C++
and Go demonstrations. See the [Lab 5 README](lab5/README.md) for installation
and usage instructions.

## Repository structure

```text
con-dev-labs/
├── lab1/
├── lab2/
├── lab3/
├── lab4/
├── lab5/
├── go.mod
└── README.md
```

Future lab sessions should be added as `lab6`, `lab7`, and so on. Each lab
should include its own README describing its requirements, how to run it, and
the files it contains.

## Licence

This repository is licensed under the [MIT License](LICENSE). Individual labs
also contain copies of the license so they remain self-contained.
