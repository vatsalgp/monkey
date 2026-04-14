package parser

import (
	"fmt"

	"github.com/vatsalgp/monkey/ast"
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

func (p *Parser) expectCurrTokType(tokType token.Type) bool {
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

// TODO: Handle errors: Empty token or incorrect token

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currTokType() != token.END_OF_FILE.Type {
		stmt := p.parseStatement()
		if stmt == nil {
			// p.logError("Expected Statement"); // Needed?
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
	if !p.expectCurrTokType(token.LET.Type) {
		return nil
	}
	tok := p.currTok
	p.advanceToken()
	return tok
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	if !p.expectCurrTokType(token.IDENTIFIER.Type) {
		return nil
	}
	iden := &ast.Identifier{Token: p.currTok, Value: p.currTok.Literal}
	p.advanceToken()
	return iden
}

func (p *Parser) parseAssign() bool {
	if !p.expectCurrTokType(token.ASSIGN.Type) {
		return false
	}
	p.advanceToken()
	return true
}

func (p *Parser) parseSemi() bool {
	if !p.expectCurrTokType(token.SEMICOLON.Type) {
		return false
	}
	p.advanceToken()
	return true
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	letStmt := &ast.LetStatement{}

	// let
	letToken := p.parseLetToken()
	if letToken == nil {
		return nil
	}
	letStmt.Token = letToken

	// x
	iden := p.parseIdentifier()
	if iden == nil {
		return nil
	}
	letStmt.Name = iden

	// =
	isAssign := p.parseAssign()
	if !isAssign {
		return nil
	}

	// (y)
	expr := p.parseExpression()
	// TODO: Do Expr error handling when expr code is done
	// if expr == nil {
	// 	return nil
	// }
	letStmt.Value = expr

	return letStmt
}

func (p *Parser) parseExpression() ast.Expression {
	// TODO: Parse Literal

	if p.currTokType() == token.IDENTIFIER.Type && p.peekTokType() == token.SEMICOLON.Type {
		iden := p.parseIdentifier()
		if iden == nil {
			return nil
		}

		// ;
		isSemi := p.parseSemi()
		if !isSemi {
			return nil
		}

		return iden
	}

	for p.currTokType() != token.SEMICOLON.Type {
		p.advanceToken()
	}

	// ;
	isSemi := p.parseSemi()
	if !isSemi {
		return nil
	}

	return nil
}
