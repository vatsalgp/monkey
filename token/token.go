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
	return fmt.Sprintf("Token(%s, '%s')", token.Type, token.Literal)
}

func CreateToken(word string) Token {
	if word == "" {
		return Token{
			Type:    END_OF_FILE,
			Literal: TokenLiteral(END_OF_FILE),
		}
	}

	if word == "let" {
		return Token{
			Type:    LET,
			Literal: TokenLiteral(LET),
		}
	}

	if word == "fn" {
		return Token{
			Type:    FUNCTION,
			Literal: TokenLiteral(FUNCTION),
		}
	}

	if word == "return" {
		return Token{
			Type:    RETURN,
			Literal: TokenLiteral(RETURN),
		}
	}

	if word == "+" {
		return Token{
			Type:    PLUS,
			Literal: TokenLiteral(PLUS),
		}
	}

	if word == "-" {
		return Token{
			Type:    MINUS,
			Literal: TokenLiteral(MINUS),
		}
	}

	if word == "*" {
		return Token{
			Type:    MULTIPLY,
			Literal: TokenLiteral(MULTIPLY),
		}
	}

	if word == "/" {
		return Token{
			Type:    DIVIDE,
			Literal: TokenLiteral(DIVIDE),
		}
	}

	if word == "=" {
		return Token{
			Type:    ASSIGN,
			Literal: TokenLiteral(ASSIGN),
		}
	}

	if word == "==" {
		return Token{
			Type:    IS_EQUAL,
			Literal: TokenLiteral(IS_EQUAL),
		}
	}

	if word == "!=" {
		return Token{
			Type:    IS_NOT_EQUAL,
			Literal: TokenLiteral(IS_NOT_EQUAL),
		}
	}

	if word == "<" {
		return Token{
			Type:    IS_LESS_THAN,
			Literal: TokenLiteral(IS_LESS_THAN),
		}
	}

	if word == ">" {
		return Token{
			Type:    IS_GREATER_THAN,
			Literal: TokenLiteral(IS_GREATER_THAN),
		}
	}

	if word == "<=" {
		return Token{
			Type:    IS_LESS_THAN_OR_EQUAL,
			Literal: TokenLiteral(IS_LESS_THAN_OR_EQUAL),
		}
	}

	if word == ">=" {
		return Token{
			Type:    IS_GREATER_THAN_OR_EQUAL,
			Literal: TokenLiteral(IS_GREATER_THAN_OR_EQUAL),
		}
	}

	if word == "," {
		return Token{
			Type:    COMMA,
			Literal: TokenLiteral(COMMA),
		}
	}

	if word == ";" {
		return Token{
			Type:    SEMICOLON,
			Literal: TokenLiteral(SEMICOLON),
		}
	}

	if word == "(" {
		return Token{
			Type:    LEFT_PAREN,
			Literal: TokenLiteral(LEFT_PAREN),
		}
	}

	if word == ")" {
		return Token{
			Type:    RIGHT_PAREN,
			Literal: TokenLiteral(RIGHT_PAREN),
		}
	}

	if word == "{" {
		return Token{
			Type:    LEFT_BRACE,
			Literal: TokenLiteral(LEFT_BRACE),
		}
	}

	if word == "}" {
		return Token{
			Type:    RIGHT_BRACE,
			Literal: TokenLiteral(RIGHT_BRACE),
		}
	}

	if word == "true" {
		return Token{
			Type:    BOOLEAN_LITERAL,
			Literal: TokenLiteral("true"),
		}
	}

	if word == "false" {
		return Token{
			Type:    BOOLEAN_LITERAL,
			Literal: TokenLiteral("false"),
		}
	}

	isInteger, _ := regexp.MatchString(`^\d+$`, word)
	if isInteger {
		return Token{
			Type:    INTEGER_LITERAL,
			Literal: TokenLiteral(word),
		}
	}

	isFloat, _ := regexp.MatchString(`^\d+\.\d+$`, word)
	if isFloat {
		return Token{
			Type:    FLOAT_LITERAL,
			Literal: TokenLiteral(word),
		}
	}

	isString, _ := regexp.MatchString(`^\".*\"$`, word)
	if isString {
		return Token{
			Type:    STRING_LITERAL,
			Literal: TokenLiteral(word),
		}
	}

	isIdentifier, _ := regexp.MatchString(`^[a-zA-Z_]+\w*$`, word)
	if isIdentifier {
		return Token{
			Type:    IDENTIFIER,
			Literal: TokenLiteral(word),
		}
	}

	return Token{
		Type:    ILLEGAL,
		Literal: TokenLiteral(ILLEGAL),
	}
}
