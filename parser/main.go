package parser

import (
	"github.com/vatsalgp/monkey/lexer"
)

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}

	p.advanceToken() // Set peekTok
	p.advanceToken() // Set currTok

	return p
}
