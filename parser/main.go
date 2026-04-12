package parser

import (
	"github.com/vatsalgp/monkey/lexer"
)

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}

	p.AdvanceToken() // Set peekTok
	p.AdvanceToken() // Set currTok

	return p
}
