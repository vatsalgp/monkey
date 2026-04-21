package parser

import (
	"github.com/vatsalgp/monkey/lexer"
	"github.com/vatsalgp/monkey/token"
)

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex, errors: []string{}}

	p.advanceToken() // Set peekTok
	p.advanceToken() // Set currTok

	p.prefixParseFns = make(map[token.Type]prefixParseFn)
	p.registerPrefix(token.IDENTIFIER.Type, p.parseIdentifier)
	p.registerPrefix(token.INTEGER_LITERAL.Type, p.parseIntegerLiteral)

	return p
}
