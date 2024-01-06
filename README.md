# Monkey Programming language

This repo is a codealong to Thorsten Ball's books: "Writing an Interpreter in Go" and "Writing a Compiler in Go". Tests can be run by cding into the respective folders and running `go test`.

### Usage
```
$ go run main.go
```

Will start the REPL for the Monkey language. It works a lot like Ruby does and follows the same kind of semantics, except we use `let` to define variables.

Scoping rules are also enforced for functions (no leaking of function local vars).