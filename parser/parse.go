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

func (p *Parser) Errors() []string {
	return p.errors
}
