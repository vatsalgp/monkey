package statements

import (
	"github.com/vatsalgp/monkey/ast"
	"github.com/vatsalgp/monkey/ast/expressions"
	"github.com/vatsalgp/monkey/token"
)

// Eg: "let x = 5 ;"
type LetStatement struct {
	Token *token.Token            // let
	Name  *expressions.Identifier // x
	Value *ast.Expression         // 5
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {}
