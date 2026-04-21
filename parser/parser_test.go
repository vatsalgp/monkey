package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/vatsalgp/monkey/ast"
	"github.com/vatsalgp/monkey/lexer"
)

func TestIncorrectStatement(t *testing.T) {
	input := `
		a = 6;
	`

	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()
	errors := p.Errors()

	fmt.Printf("parser has %d errors\n", len(errors))

	for _, msg := range errors {
		fmt.Printf("parser error: %q\n", msg)
	}
}

func TestCorrectLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := (program.Statements)[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestIncorrectLetStatements(t *testing.T) {
	input := `
		let x 5;
		let = 10;
		let 838383;
	`

	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()
	errors := p.Errors()

	fmt.Printf("parser has %d errors\n", len(errors))

	for _, msg := range errors {
		fmt.Printf("parser error: %q\n", msg)
	}
}
func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.FirstTokenLiteral() != "return" {
			t.Errorf("returnStmt.FirstTokenLiteral not 'return', got %q",
				returnStmt.FirstTokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.FirstTokenLiteral() != "let" {
		t.Errorf("s.FirstTokenLiteral not 'let'. got=%q", s.FirstTokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.FirstTokenLiteral() != name {
		t.Errorf("letStmt.Name.FirstTokenLiteral() not '%s'. got=%s",
			name, letStmt.Name.FirstTokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.FirstTokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.FirstTokenLiteral())
	}
}

func TestExpression(t *testing.T) {
	input := `
		let x = false;
		let y = x;
		return y;
	`

	expectation := `
		let x = false;
		let y = x;
		return y;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	output := program.String()

	trimmedOutput := strings.TrimSpace(output)
	trimmedExpectation := strings.ReplaceAll(strings.TrimSpace(expectation), "\t", "")

	if trimmedOutput != trimmedExpectation {
		t.Errorf("program.String() wrong.\ngot=%q\nexpected=%q", trimmedOutput, trimmedExpectation)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.FirstTokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.FirstTokenLiteral())
	}
}
