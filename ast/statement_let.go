package ast

import (
	"github.com/vatsalgp/monkey/token"
)

// Eg: "let x = 5 ;"
type LetStatement struct {
	Token *token.Token // let
	Name  *Identifier  // x
	Value Expression   // 5
}

func (ls LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls LetStatement) statementNode() {}
