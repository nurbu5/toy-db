package interpreter

import (
	"reflect"
	"testing"
)

type mockWriter struct {
	contents string
}

type mockDB struct {
	contents string
}

func (m *mockWriter) Write(p []byte) (int, error) {
	m.contents += string(p)
	return 0, nil
}

func (m *mockDB) Insert(args []string) error {
	return nil
}

func (m *mockDB) Select(args []string) error {
	return nil
}

func TestRunCommand(t *testing.T) {
	t.Run("the '.exit' command", func(t *testing.T) {
		t.Parallel()
		w := &mockWriter{}
		db := &mockDB{}
		in := New(w, db)

		in.RunCommand(".exit")
		got := in.Terminated()
		want := true

		if got != want {
			t.Errorf("Expected the interpreter to be terminated. got: %t, want: %t", got, want)
		}
	})

	t.Run("a metacommand that doesn't exist", func(t *testing.T) {
		t.Parallel()
		w := &mockWriter{}
		db := &mockDB{}
		in := New(w, db)

		in.RunCommand(".blahblahblah")
		got := w.contents
		want := "Unrecognized command '.blahblahblah'.\n"

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("an insert statement", func(t *testing.T) {
		t.Parallel()
		w := &mockWriter{}
		db := &mockDB{}
		in := New(w, db)

		in.RunCommand("INSERT 1 gopher test@foo.com")
		got := in.statement
		want := &statement{kind: STATEMENT_INSERT, args: []string{"1", "gopher", "test@foo.com"}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("an select statement", func(t *testing.T) {
		t.Parallel()
		w := &mockWriter{}
		db := &mockDB{}
		in := New(w, db)

		in.RunCommand("SELECT *")
		got := in.statement
		want := &statement{kind: STATEMENT_SELECT, args: []string{"*"}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("a command that doesn't exist", func(t *testing.T) {
		t.Parallel()
		w := &mockWriter{}
		db := &mockDB{}
		in := New(w, db)

		in.RunCommand("chicken soup")
		got := w.contents
		want := "Unrecognized keyword at start of 'chicken soup'.\n"

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}
