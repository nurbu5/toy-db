package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nurbu5/toydb/interpreter"
	"github.com/nurbu5/toydb/repl"
)

type mockDB struct {
	contents string
}

func (m *mockDB) Insert(args []string) error {
	fmt.Printf("Ran INSERT with the following arguments: %v\n", args)
	return nil
}

func (m *mockDB) Select(args []string) error {
	fmt.Printf("Ran SELECT with the following arguments: %v\n", args)
	return nil
}

func main() {
	w := os.Stdout
	r := bufio.NewReader(os.Stdin)
	db := new(mockDB)
	i := interpreter.New(w, db)
	repl.Start(&i, r)
}
