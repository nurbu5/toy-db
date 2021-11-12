// package interpreter implements a basic database interpreter which parses and runs commands. It writes output to the
// io.Writer that is given when it is instantiated.
package interpreter

import (
	"fmt"
	"io"
)

const prompt = "db > "

type Interpreter struct {
	terminated bool
	writer     io.Writer
	db         database
	statement  *statement
}

type database interface {
	Insert([]string) error
	Select([]string) error
}

// New creates a new Interpreter with the given writer
func New(w io.Writer, db database) Interpreter {
	return Interpreter{terminated: false, writer: w, db: db}
}

// Terminated returns whether or not the interpreter has terminated
func (i *Interpreter) Terminated() bool {
	return i.terminated
}

// OutputPrompt writes the value in "prompt" to the interpreter's writer
func (i *Interpreter) OutputPrompt() {
	fmt.Fprint(i.writer, prompt)
}

// RunCommand takes a command string as input and runs it
func (i *Interpreter) RunCommand(c string) {
	if isMetaCommand(c) {
		i.handleMetaCommand(c)
		return
	}
	i.prepareStatement(c)
	i.executeStatement()
}
