package parser

import (
	"github.com/vatsalgp/monkey/ast"
	"github.com/vatsalgp/monkey/lexer"
	"github.com/vatsalgp/monkey/token"
)

type Parser struct {
	lex     *lexer.Lexer
	currTok *token.Token
	peekTok *token.Token
}

func (p *Parser) AdvanceToken() {
	p.currTok = p.peekTok
	p.peekTok = p.lex.NextToken()
}

// TODO: Handle errors: Empty token or incorrect token

func (p *Parser) ParseProgram() *ast.Program {
	var program = &ast.Program{}

	for {
		if p.currTok == nil || p.currTok.Type == token.END_OF_FILE.Type {
			break
		}

		if p.currTok.Type == token.LET.Type {
			var letStmt = p.parseLetStmt()
			program.Statements = append(program.Statements, letStmt)
		}
	}

	return program
}

func (p *Parser) parseLetToken() *token.Token {
	letToken := &token.Token{Type: token.LET.Type, Literal: "let"}
	p.AdvanceToken()
	return letToken
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	tok := &token.Token{Type: token.IDENTIFIER.Type, Literal: p.currTok.Literal}
	iden := &ast.Identifier{Token: tok, Value: p.currTok.Literal}
	p.AdvanceToken()
	return iden
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	var letStmt = &ast.LetStatement{}

	// let
	letStmt.Token = p.parseLetToken()

	// x
	letStmt.Name = p.parseIdentifier()

	// =
	p.AdvanceToken()

	// (y)
	letStmt.Value = p.parseExpression()

	return letStmt
}

func (p *Parser) parseExpression() ast.Expression {
	// TODO: Parse Literal

	if p.currTok.Type == token.IDENTIFIER.Type && p.peekTok.Type == token.SEMICOLON.Type {
		iden := p.parseIdentifier()

		// ;
		p.AdvanceToken()

		return iden
	}

	for {
		if p.currTok.Type == token.SEMICOLON.Type {
			break
		}
		p.AdvanceToken()
	}

	// ;
	p.AdvanceToken()

	return nil
}
