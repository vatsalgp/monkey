package lexer

import (
	"testing"

	"github.com/vatsalgp/monkey/token"
)

func TestNextToken(test *testing.T) {
	input := `=+(){},;`

	tokens := []token.Token{
		{Type: token.ASSIGN.Type, Literal: "="},
		{Type: token.PLUS.Type, Literal: "+"},
		{Type: token.LEFT_PAREN.Type, Literal: "("},
		{Type: token.RIGHT_PAREN.Type, Literal: ")"},
		{Type: token.LEFT_BRACE.Type, Literal: "{"},
		{Type: token.RIGHT_BRACE.Type, Literal: "}"},
		{Type: token.COMMA.Type, Literal: ","},
		{Type: token.SEMICOLON.Type, Literal: ";"},
		{Type: token.END_OF_FILE.Type, Literal: ""},
	}

	lex := New(input)

	for i, tok := range tokens {
		testTok := lex.NextToken()

		if testTok.Type != tok.Type {
			test.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tok.Type, testTok.Type)
		}

		if testTok.Literal != tok.Literal {
			test.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tok.Literal, testTok.Literal)
		}
	}
}
