package ast

import (
	"github.com/vatsalgp/monkey/token"
)

// Eg: "x+5;"
type ExpressionStatement struct {
	Token      *token.Token // the first token of the expression
	Expression Expression   // x+5
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) statementNode() {}
