package ast

import (
	"github.com/vatsalgp/monkey/token"
)

// Eg: "6"
type IntegerLiteral struct {
	Token *token.Token // 6
	Value int64        // 6
}

func (intLit *IntegerLiteral) FirstTokenLiteral() string {
	return intLit.Token.Literal
}

func (intLit *IntegerLiteral) expressionNode() {}

func (intLit *IntegerLiteral) String() string {
	return intLit.Token.Literal
}
