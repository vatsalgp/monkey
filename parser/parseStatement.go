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
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	letStmt := &ast.LetStatement{}

	// let
	letToken := p.parseLetToken()
	letStmt.Token = letToken

	// x
	iden, ok := p.parseIdentifier().(*ast.Identifier)
	if !ok {
		return nil
	}
	letStmt.Name = iden

	// =
	p.parseAssignToken()

	// (y)
	expr := p.parseExpression(token.LOWEST)
	letStmt.Value = expr

	return letStmt
}

func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
	returnStmt := &ast.ReturnStatement{}

	// return
	returnToken := p.parseReturnToken()
	returnStmt.Token = returnToken

	// (y)
	expr := p.parseExpression(token.LOWEST)
	returnStmt.ReturnValue = expr

	return returnStmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	exprStmt := &ast.ExpressionStatement{}

	exprStmt.Token = p.currTok

	expr := p.parseExpression(token.LOWEST)
	exprStmt.Expression = expr

	return exprStmt
}
