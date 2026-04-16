package ast

import (
	"testing"

	"github.com/vatsalgp/monkey/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: &token.Token{Type: token.LET.Type, Literal: "let"},
				Name: &Identifier{
					Token: &token.Token{Type: token.IDENTIFIER.Type, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: &token.Token{Type: token.IDENTIFIER.Type, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;\n" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
