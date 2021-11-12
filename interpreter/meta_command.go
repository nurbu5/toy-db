package interpreter

import "fmt"

// handleMetaCommand reads a metacommand string and determines which how to execute that metacommand
func (i *Interpreter) handleMetaCommand(c string) {
	switch c {
	case ".exit":
		i.terminated = true
	default:
		fmt.Fprintf(i.writer, "Unrecognized command '%s'.\n", c)
	}
}

// isMetaCommand determines whether or not a string is a metacommand
func isMetaCommand(c string) bool {
	return len(c) > 0 && c[0] == '.'
}
