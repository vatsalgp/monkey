package parser

import (
	"github.com/vatsalgp/monkey/ast"
	"github.com/vatsalgp/monkey/token"
)

func (p *Parser) parseExpression() ast.Expression {
	// TODO: Parse Literal

	if p.currTokType() == token.IDENTIFIER.Type && p.peekTokType() == token.SEMICOLON.Type {
		iden := p.parseIdentifier()

		p.parseSemiColonToken()

		return iden
	}

	for p.currTokType() != token.SEMICOLON.Type {
		p.advanceToken()
	}

	p.parseSemiColonToken()

	return nil
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	if !p.expectTokType(token.IDENTIFIER.Type) {
		return nil
	}
	iden := &ast.Identifier{Token: p.currTok, Value: p.currTok.Literal}
	p.advanceToken()
	return iden
}
