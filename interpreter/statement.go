package interpreter

import (
	"fmt"
	"strings"
)

type StatementKind int

type statement struct {
	kind StatementKind
	args []string
}

const (
	STATEMENT_INSERT StatementKind = iota
	STATEMENT_SELECT
)

func (s *statement) setArgs(c string) {
	args := strings.Split(c, " ")
	if len(args) > 1 {
		s.args = args[1:]
	}
}

func (i *Interpreter) prepareStatement(c string) {
	if len(c) >= 6 {
		if c[:6] == "INSERT" {
			s := &statement{kind: STATEMENT_INSERT}
			s.setArgs(c)
			i.statement = s
			return
		}
		if c[:6] == "SELECT" {
			s := &statement{kind: STATEMENT_SELECT}
			s.setArgs(c)
			i.statement = s
			return
		}
	}
	fmt.Fprintf(i.writer, "Unrecognized keyword at start of '%s'.\n", c)
}

func (i *Interpreter) executeStatement() {
	if i.statement == nil {
		return
	}

	var err error
	switch i.statement.kind {
	case STATEMENT_INSERT:
		err = i.db.Insert(i.statement.args)
	case STATEMENT_SELECT:
		err = i.db.Select(i.statement.args)
	}

	if err != nil {
		fmt.Fprintf(i.writer, "Error: '%s'", err)
	}
}
