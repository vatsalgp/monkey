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

// TODO: Handle errors: Empty token or incorrect token

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currTokType() != token.END_OF_FILE.Type {
		if p.currTokType() == token.LET.Type {
			letStmt := p.parseLetStmt()
			program.Statements = append(program.Statements, letStmt)
		}
	}

	return program
}

func (p *Parser) parseLetToken() *token.Token {
	letToken := &token.Token{Type: token.LET.Type, Literal: "let"}
	p.advanceToken()
	return letToken
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	tok := &token.Token{Type: token.IDENTIFIER.Type, Literal: p.currTok.Literal}
	iden := &ast.Identifier{Token: tok, Value: p.currTok.Literal}
	p.advanceToken()
	return iden
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	letStmt := &ast.LetStatement{}

	// let
	letStmt.Token = p.parseLetToken()

	// x
	letStmt.Name = p.parseIdentifier()

	// =
	p.advanceToken()

	// (y)
	letStmt.Value = p.parseExpression()

	return letStmt
}

func (p *Parser) parseExpression() ast.Expression {
	// TODO: Parse Literal

	if p.currTokType() == token.IDENTIFIER.Type && p.peekTokType() == token.SEMICOLON.Type {
		iden := p.parseIdentifier()

		// ;
		p.advanceToken()

		return iden
	}

	for {
		if p.currTokType() == token.SEMICOLON.Type {
			break
		}
		p.advanceToken()
	}

	// ;
	p.advanceToken()

	return nil
}
