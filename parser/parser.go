package parser

import (
	"fmt"

	"github.com/vatsalgp/monkey/lexer"
	"github.com/vatsalgp/monkey/token"
)

type Parser struct {
	lex     *lexer.Lexer
	currTok *token.Token
	peekTok *token.Token
	errors  []string
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
	p.logTypeError(tokType)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) logError(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) logTypeError(tok token.Type) {
	p.logError(fmt.Sprintf("expected token to be %s, got %s instead", tok, p.currTokType()))
}
