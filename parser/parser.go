package parser

import (
	"fmt"

	"github.com/vatsalgp/monkey/ast"
	"github.com/vatsalgp/monkey/lexer"
	"github.com/vatsalgp/monkey/token"
)

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

type Parser struct {
	lex     *lexer.Lexer
	currTok *token.Token
	peekTok *token.Token
	errors  []string

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

func (p *Parser) registerPrefix(tt token.Type, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}

func (p *Parser) registerInfix(tt token.Type, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}

func (p *Parser) advanceToken() {
	p.currTok = p.peekTok
	p.peekTok = p.lex.NextToken()
}

func (p *Parser) currTokType() token.Type {
	if p.currTok == nil {
		return token.END_OF_FILE.Type
	}
	return p.currTok.Type
}

func (p *Parser) peekTokType() token.Type {
	if p.peekTok == nil {
		return token.END_OF_FILE.Type
	}
	return p.peekTok.Type
}

func (p *Parser) expectTokType(tokType token.Type) bool {
	if p.currTokType() == tokType {
		return true
	}
	p.logError(fmt.Sprintf("expected token to be %s, got %s instead", tokType, p.currTokType()))
	return false
}

func (p *Parser) logError(msg string) {
	p.errors = append(p.errors, msg)
}
