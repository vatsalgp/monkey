package parser

import (
	"github.com/vatsalgp/monkey/ast"
	"github.com/vatsalgp/monkey/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currTokType() {
	case token.LET.Type:
		return p.parseLetStmt()
	case token.RETURN.Type:
		return p.parseReturnStmt()
	default:
		p.logError("Expected valid Let or Return Statement")
		return nil
	}
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
	p.parseAssignToken()

	// (y)
	expr := p.parseExpression()
	letStmt.Value = expr

	return letStmt
}

func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
	returnStmt := &ast.ReturnStatement{}

	// return
	returnToken := p.parseReturnToken()
	returnStmt.Token = returnToken

	// (y)
	expr := p.parseExpression()
	returnStmt.ReturnValue = expr

	return returnStmt
}
