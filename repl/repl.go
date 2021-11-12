// Package repl provides a method to start a repl
package repl

import (
	"log"
	"strings"
)

// stringReader is the interface that wraps the ReadString method.
type stringReader interface {
	ReadString(byte) (string, error)
}

type interpreter interface {
	Terminated() bool
	OutputPrompt()
	RunCommand(string)
}

// Start starts the repl. It writes output to the given io.Writer and reads input from the given
// stringReader.
func Start(i interpreter, r stringReader) {
	for {
		i.OutputPrompt()
		in := readInput(r)

		i.RunCommand(in)
		if i.Terminated() {
			return
		}
	}
}

func readInput(r stringReader) string {
	s, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(s)
}
