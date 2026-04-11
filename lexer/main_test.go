package lexer

import (
	"testing"

	"github.com/vatsalgp/monkey/token"
)

type TextTokenTest struct {
	input  string
	tokens []token.Token
}

var (
	TESTS = []TextTokenTest{
		{
			input: `=+(){},;`,
			tokens: []token.Token{
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.PLUS.Type, Literal: "+"},
				{Type: token.LEFT_PAREN.Type, Literal: "("},
				{Type: token.RIGHT_PAREN.Type, Literal: ")"},
				{Type: token.LEFT_BRACE.Type, Literal: "{"},
				{Type: token.RIGHT_BRACE.Type, Literal: "}"},
				{Type: token.COMMA.Type, Literal: ","},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.END_OF_FILE.Type, Literal: ""},
			},
		},
		{
			input: `
				let five = 5;
				let ten = 10;

				let add = fn(x, y) {
					x + y;
				};

				let result = add(five, ten);
			`,
			tokens: []token.Token{
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "five"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.INTEGER_LITERAL.Type, Literal: "5"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "ten"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.INTEGER_LITERAL.Type, Literal: "10"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "add"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.FUNCTION.Type, Literal: "fn"},
				{Type: token.LEFT_PAREN.Type, Literal: "("},
				{Type: token.IDENTIFIER.Type, Literal: "x"},
				{Type: token.COMMA.Type, Literal: ","},
				{Type: token.IDENTIFIER.Type, Literal: "y"},
				{Type: token.RIGHT_PAREN.Type, Literal: ")"},
				{Type: token.LEFT_BRACE.Type, Literal: "{"},
				{Type: token.IDENTIFIER.Type, Literal: "x"},
				{Type: token.PLUS.Type, Literal: "+"},
				{Type: token.IDENTIFIER.Type, Literal: "y"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.RIGHT_BRACE.Type, Literal: "}"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "result"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.IDENTIFIER.Type, Literal: "add"},
				{Type: token.LEFT_PAREN.Type, Literal: "("},
				{Type: token.IDENTIFIER.Type, Literal: "five"},
				{Type: token.COMMA.Type, Literal: ","},
				{Type: token.IDENTIFIER.Type, Literal: "ten"},
				{Type: token.RIGHT_PAREN.Type, Literal: ")"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.END_OF_FILE.Type, Literal: ""},
			},
		},
	}
)

func TestNextToken(test *testing.T) {
	for _, tt := range TESTS {
		lex := New(tt.input)

		for i, tok := range tt.tokens {
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
}
