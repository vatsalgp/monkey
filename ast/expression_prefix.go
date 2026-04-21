package ast

import (
	"fmt"

	"github.com/vatsalgp/monkey/token"
)

type PrefixExpression struct {
	Token    *token.Token
	Operator *token.Token
	Right    Expression
}

func (pe *PrefixExpression) FirstTokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", pe.Operator.Literal, pe.Right.String())
}
