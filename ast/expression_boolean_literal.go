package ast

import (
	"github.com/vatsalgp/monkey/token"
)

// Eg: "true"
type BooleanLiteral struct {
	Token *token.Token // true
	Value bool         // true
}

func (boolLit *BooleanLiteral) FirstTokenLiteral() string {
	return boolLit.Token.Literal
}

func (boolLit *BooleanLiteral) expressionNode() {}

func (boolLit *BooleanLiteral) String() string {
	return boolLit.Token.Literal
}
