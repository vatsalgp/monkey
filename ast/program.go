package ast

import (
	"strings"
)

// The Whole Monkey Program
type Program struct {
	Statements []Statement
}

func (prog *Program) FirstTokenLiteral() string {
	if len(prog.Statements) == 0 {
		return ""
	}
	return prog.Statements[0].String()
}

func (prog *Program) String() string {
	sb := strings.Builder{}

	for _, stat := range prog.Statements {
		sb.WriteString(stat.String())
		sb.WriteByte('\n')
	}

	return sb.String()
}

func NewProgram() *Program {
	return &Program{Statements: []Statement{}}
}
