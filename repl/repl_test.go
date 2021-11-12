package repl

import (
	"fmt"
	"testing"
)

type mockInterpreter struct {
	terminated bool
	contents   string
}

type mockReader struct {
	commands []string
}

func (m *mockInterpreter) Terminated() bool {
	return m.terminated
}

func (m *mockInterpreter) OutputPrompt() {
	m.contents += "Interpreter > "
}

func (m *mockInterpreter) RunCommand(c string) {
	if c == "exit" {
		m.terminated = true
	} else {
		m.contents += fmt.Sprintf("Unknown command: '%s'.\n", c)
	}
}

func (m *mockReader) ReadString(b byte) (string, error) {
	c := m.commands[0]
	m.commands = m.commands[1:]
	return c, nil
}

func TestStart(t *testing.T) {
	t.Run("exit the repl", func(t *testing.T) {
		i := new(mockInterpreter)
		r := &mockReader{
			commands: []string{"exit\n"},
		}

		Start(i, r)
		got := i.contents
		want := "Interpreter > "

		if got != want {
			t.Errorf("got: %s, want %s", got, want)
		}
	})

	t.Run("use an unknown command", func(t *testing.T) {
		i := new(mockInterpreter)
		r := &mockReader{
			commands: []string{"blah", "exit\n"},
		}

		Start(i, r)
		got := i.contents
		want := "Interpreter > "
		want += "Unknown command: 'blah'.\n"
		want += "Interpreter > "

		if got != want {
			t.Errorf("got: %s, want %s", got, want)
		}
	})
}
