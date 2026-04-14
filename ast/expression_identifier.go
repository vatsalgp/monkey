package ast

import (
	"github.com/vatsalgp/monkey/token"
)

// Eg: "x"
type Identifier struct {
	Token *token.Token // x
	Value string       // 0
}

func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Literal
}

func (ident *Identifier) expressionNode() {}
