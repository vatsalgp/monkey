package parser

import (
	"fmt"

	"github.com/vatsalgp/monkey/ast"
	"github.com/vatsalgp/monkey/token"
)

func (p *Parser) parseExpression(precedence token.Precedence) ast.Expression {
	// TODO: Parse all kinds of literals

	prefixParseFn := p.prefixParseFns[p.currTokType()]

	if prefixParseFn == nil {
		if (p.currTokType() == token.TRUE.Type || p.currTokType() == token.FALSE.Type) && p.peekTokType() == token.SEMICOLON.Type {
			boolLit := p.parseBooleanLiteral()

			p.parseSemiColonToken()

			return boolLit
		}

		for p.currTokType() != token.SEMICOLON.Type {
			p.advanceToken()
		}

		p.parseSemiColonToken()

		return nil
	}

	leftExpr := prefixParseFn()

	if p.currTok.Type == token.SEMICOLON.Type {
		p.advanceToken()
	}

	return leftExpr

}

func (p *Parser) parseIdentifier() ast.Expression {
	if !p.expectTokType(token.IDENTIFIER.Type) {
		return nil
	}
	iden := &ast.Identifier{Token: p.currTok, Value: p.currTok.Literal}
	p.advanceToken()
	return iden
}

func (p *Parser) parseTrue() *ast.BooleanLiteral {
	if !p.expectTokType(token.TRUE.Type) {
		return nil
	}
	boolLit := &ast.BooleanLiteral{Token: p.currTok, Value: p.currTok.Literal}
	p.advanceToken()
	return boolLit
}

func (p *Parser) parseFalse() *ast.BooleanLiteral {
	if !p.expectTokType(token.FALSE.Type) {
		return nil
	}
	boolLit := &ast.BooleanLiteral{Token: p.currTok, Value: p.currTok.Literal}
	p.advanceToken()
	return boolLit
}

func (p *Parser) parseBooleanLiteral() *ast.BooleanLiteral {
	if p.currTokType() == token.TRUE.Type {
		return p.parseTrue()
	}
	if p.currTokType() == token.FALSE.Type {
		return p.parseFalse()
	}
	p.logError(fmt.Sprintf("expected token to be a boolean literal, got %s instead", p.currTokType()))
	return nil
}
