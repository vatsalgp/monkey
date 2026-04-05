package token

import (
	"fmt"
	"regexp"
)

type Token struct {
	Type    TokenType
	Literal TokenLiteral
}

func (token Token) String() string {
	return fmt.Sprintf("Token(%s, %s)", token.Type, token.Literal)
}

func CreateToken(s string) Token {
	if s == "" {
		return Token{
			Type:    END_OF_FILE,
			Literal: TokenLiteral(END_OF_FILE),
		}
	}

	if s == "let" {
		return Token{
			Type:    LET,
			Literal: TokenLiteral(LET),
		}
	}

	if s == "fn" {
		return Token{
			Type:    FUNCTION,
			Literal: TokenLiteral(FUNCTION),
		}
	}

	if s == "return" {
		return Token{
			Type:    RETURN,
			Literal: TokenLiteral(RETURN),
		}
	}

	if s == "+" {
		return Token{
			Type:    PLUS,
			Literal: TokenLiteral(PLUS),
		}
	}

	if s == "-" {
		return Token{
			Type:    MINUS,
			Literal: TokenLiteral(MINUS),
		}
	}

	if s == "*" {
		return Token{
			Type:    MULTIPLY,
			Literal: TokenLiteral(MULTIPLY),
		}
	}

	if s == "/" {
		return Token{
			Type:    DIVIDE,
			Literal: TokenLiteral(DIVIDE),
		}
	}

	if s == "=" {
		return Token{
			Type:    ASSIGN,
			Literal: TokenLiteral(ASSIGN),
		}
	}

	if s == "==" {
		return Token{
			Type:    IS_EQUAL,
			Literal: TokenLiteral(IS_EQUAL),
		}
	}

	if s == "!=" {
		return Token{
			Type:    IS_NOT_EQUAL,
			Literal: TokenLiteral(IS_NOT_EQUAL),
		}
	}

	if s == "<" {
		return Token{
			Type:    IS_LESS_THAN,
			Literal: TokenLiteral(IS_LESS_THAN),
		}
	}

	if s == ">" {
		return Token{
			Type:    IS_GREATER_THAN,
			Literal: TokenLiteral(IS_GREATER_THAN),
		}
	}

	if s == "<=" {
		return Token{
			Type:    IS_LESS_THAN_OR_EQUAL,
			Literal: TokenLiteral(IS_LESS_THAN_OR_EQUAL),
		}
	}

	if s == ">=" {
		return Token{
			Type:    IS_GREATER_THAN_OR_EQUAL,
			Literal: TokenLiteral(IS_GREATER_THAN_OR_EQUAL),
		}
	}

	if s == "," {
		return Token{
			Type:    COMMA,
			Literal: TokenLiteral(COMMA),
		}
	}

	if s == ";" {
		return Token{
			Type:    SEMICOLON,
			Literal: TokenLiteral(SEMICOLON),
		}
	}

	if s == "(" {
		return Token{
			Type:    LEFT_PAREN,
			Literal: TokenLiteral(LEFT_PAREN),
		}
	}

	if s == ")" {
		return Token{
			Type:    RIGHT_PAREN,
			Literal: TokenLiteral(RIGHT_PAREN),
		}
	}

	if s == "{" {
		return Token{
			Type:    LEFT_BRACE,
			Literal: TokenLiteral(LEFT_BRACE),
		}
	}

	if s == "}" {
		return Token{
			Type:    RIGHT_BRACE,
			Literal: TokenLiteral(RIGHT_BRACE),
		}
	}

	if s == "true" {
		return Token{
			Type:    BOOLEAN_LITERAL,
			Literal: TokenLiteral("true"),
		}
	}

	if s == "false" {
		return Token{
			Type:    BOOLEAN_LITERAL,
			Literal: TokenLiteral("false"),
		}
	}

	isInteger, _ := regexp.MatchString(`^\d+$`, s)
	if isInteger {
		return Token{
			Type:    INTEGER_LITERAL,
			Literal: TokenLiteral(s),
		}
	}

	isFloat, _ := regexp.MatchString(`^\d+\.\d+$`, s)
	if isFloat {
		return Token{
			Type:    FLOAT_LITERAL,
			Literal: TokenLiteral(s),
		}
	}

	isString, _ := regexp.MatchString(`^\".*\"$`, s)
	if isString {
		return Token{
			Type:    STRING_LITERAL,
			Literal: TokenLiteral(s),
		}
	}

	isIdentifier, _ := regexp.MatchString(`^[a-zA-Z_]+\w*$`, s)
	if isIdentifier {
		return Token{
			Type:    IDENTIFIER,
			Literal: TokenLiteral(s),
		}
	}

	return Token{
		Type:    ILLEGAL,
		Literal: TokenLiteral(ILLEGAL),
	}
}
