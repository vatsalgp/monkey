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
				let five_=5;
				let _TEN1 = 10;

				let x = true;
				let y = false;

				let z = 5.7;

				let add = fn(x, y) {
					return x + y;
				};

				let result = add(five_, _TEN1);
			`,
			tokens: []token.Token{
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "five_"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.INTEGER_LITERAL.Type, Literal: "5"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "_TEN1"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.INTEGER_LITERAL.Type, Literal: "10"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "x"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.TRUE.Type, Literal: "true"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "y"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.FALSE.Type, Literal: "false"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.LET.Type, Literal: "let"},
				{Type: token.IDENTIFIER.Type, Literal: "z"},
				{Type: token.ASSIGN.Type, Literal: "="},
				{Type: token.FLOAT_LITERAL.Type, Literal: "5.7"},
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
				{Type: token.RETURN.Type, Literal: "return"},
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
				{Type: token.IDENTIFIER.Type, Literal: "five_"},
				{Type: token.COMMA.Type, Literal: ","},
				{Type: token.IDENTIFIER.Type, Literal: "_TEN1"},
				{Type: token.RIGHT_PAREN.Type, Literal: ")"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.END_OF_FILE.Type, Literal: ""},
			},
		},
		{
			input: `
				!-/*5;
				5 < 10 > 5;

				if (5 < 10) {
					return true;
				} else {
					return false;
				}

				10 == 10;
				10 != 9;
			`,
			tokens: []token.Token{
				{Type: token.NOT.Type, Literal: "!"},
				{Type: token.MINUS.Type, Literal: "-"},
				{Type: token.DIVIDE.Type, Literal: "/"},
				{Type: token.MULTIPLY.Type, Literal: "*"},
				{Type: token.INTEGER_LITERAL.Type, Literal: "5"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.INTEGER_LITERAL.Type, Literal: "5"},
				{Type: token.IS_LESS_THAN.Type, Literal: "<"},
				{Type: token.INTEGER_LITERAL.Type, Literal: "10"},
				{Type: token.IS_GREATER_THAN.Type, Literal: ">"},
				{Type: token.INTEGER_LITERAL.Type, Literal: "5"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.IF.Type, Literal: "if"},
				{Type: token.LEFT_PAREN.Type, Literal: "("},
				{Type: token.INTEGER_LITERAL.Type, Literal: "5"},
				{Type: token.IS_LESS_THAN.Type, Literal: "<"},
				{Type: token.INTEGER_LITERAL.Type, Literal: "10"},
				{Type: token.RIGHT_PAREN.Type, Literal: ")"},
				{Type: token.LEFT_BRACE.Type, Literal: "{"},
				{Type: token.RETURN.Type, Literal: "return"},
				{Type: token.TRUE.Type, Literal: "true"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.RIGHT_BRACE.Type, Literal: "}"},
				{Type: token.ELSE.Type, Literal: "else"},
				{Type: token.LEFT_BRACE.Type, Literal: "{"},
				{Type: token.RETURN.Type, Literal: "return"},
				{Type: token.FALSE.Type, Literal: "false"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.RIGHT_BRACE.Type, Literal: "}"},
				{Type: token.INTEGER_LITERAL.Type, Literal: "10"},
				{Type: token.IS_EQUAL.Type, Literal: "=="},
				{Type: token.INTEGER_LITERAL.Type, Literal: "10"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.INTEGER_LITERAL.Type, Literal: "10"},
				{Type: token.IS_NOT_EQUAL.Type, Literal: "!="},
				{Type: token.INTEGER_LITERAL.Type, Literal: "9"},
				{Type: token.SEMICOLON.Type, Literal: ";"},
				{Type: token.END_OF_FILE.Type, Literal: ""},
			},
		},
	}
)

func TestNextToken(test *testing.T) {
	for tn, tt := range TESTS {
		lex := New(tt.input)

		for i, tok := range tt.tokens {
			testTok := lex.NextToken()

			if testTok.Type != tok.Type || testTok.Literal != tok.Literal {
				test.Fatalf("\ntests[%d][%d]\nexpected=%s\ngot=%s", tn+1, i+1, tok, testTok)
			}
		}
	}
}
