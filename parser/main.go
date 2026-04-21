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
	p.registerPrefix(token.NOT.Type, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS.Type, p.parsePrefixExpression)
	p.registerPrefix(token.IDENTIFIER.Type, p.parseIdentifier)
	p.registerPrefix(token.TRUE.Type, p.parseTrue)
	p.registerPrefix(token.FALSE.Type, p.parseFalse)
	p.registerPrefix(token.INTEGER_LITERAL.Type, p.parseIntegerLiteral)

	return p
}
