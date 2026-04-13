package ast

import (
	"strings"
)

// The Whole Monkey Program
type Program struct {
	Statements []Statement
}

func (prog *Program) TokenLiteral() string {
	sb := strings.Builder{}

	for _, stat := range prog.Statements {
		sb.WriteString(stat.TokenLiteral())
		sb.WriteByte('\n')
	}

	return sb.String()
}
