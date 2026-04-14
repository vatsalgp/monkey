package parser

import "github.com/vatsalgp/monkey/token"

func (p *Parser) parseLetToken() *token.Token {
	if !p.expectTokType(token.LET.Type) {
		return nil
	}
	tok := p.currTok
	p.advanceToken()
	return tok
}

func (p *Parser) parseReturnToken() *token.Token {
	if !p.expectTokType(token.RETURN.Type) {
		return nil
	}
	tok := p.currTok
	p.advanceToken()
	return tok
}

func (p *Parser) parseAssignToken() bool {
	if !p.expectTokType(token.ASSIGN.Type) {
		return false
	}
	p.advanceToken()
	return true
}

func (p *Parser) parseSemiColonToken() bool {
	if !p.expectTokType(token.SEMICOLON.Type) {
		return false
	}
	p.advanceToken()
	return true
}
