package parser

import (
	"github.com/vatsalgp/monkey/ast"
	"github.com/vatsalgp/monkey/token"
)

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currTokType() != token.END_OF_FILE.Type {
		stmt := p.parseStatement()
		if stmt == nil {
			// End parsing since we don't know how much to advance
			break
		}
		program.Statements = append(program.Statements, stmt)
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currTokType() {
	case token.LET.Type:
		return p.parseLetStmt()
	default:
		p.logError("Expected Valid Statement")
		return nil
	}
}

func (p *Parser) parseLetToken() *token.Token {
	if !p.expectTokType(token.LET.Type) {
		return nil
	}
	tok := p.currTok
	p.advanceToken()
	return tok
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	if !p.expectTokType(token.IDENTIFIER.Type) {
		return nil
	}
	iden := &ast.Identifier{Token: p.currTok, Value: p.currTok.Literal}
	p.advanceToken()
	return iden
}

func (p *Parser) parseAssign() bool {
	if !p.expectTokType(token.ASSIGN.Type) {
		return false
	}
	p.advanceToken()
	return true
}

func (p *Parser) parseSemi() bool {
	if !p.expectTokType(token.SEMICOLON.Type) {
		return false
	}
	p.advanceToken()
	return true
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	letStmt := &ast.LetStatement{}

	// let
	letToken := p.parseLetToken()
	letStmt.Token = letToken

	// x
	iden := p.parseIdentifier()
	letStmt.Name = iden

	// =
	p.parseAssign()

	// (y)
	expr := p.parseExpression()
	letStmt.Value = expr

	return letStmt
}

func (p *Parser) parseExpression() ast.Expression {
	// TODO: Parse Literal

	if p.currTokType() == token.IDENTIFIER.Type && p.peekTokType() == token.SEMICOLON.Type {
		iden := p.parseIdentifier()

		p.parseSemi()

		return iden
	}

	for p.currTokType() != token.SEMICOLON.Type {
		p.advanceToken()
	}

	p.parseSemi()

	return nil
}
