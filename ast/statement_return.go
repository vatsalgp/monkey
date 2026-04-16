package ast

import (
	"fmt"

	"github.com/vatsalgp/monkey/token"
)

// Eg: "return 5+5;"
type ReturnStatement struct {
	Token       *token.Token // return
	ReturnValue Expression   // 5+5
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) String() string {
	return fmt.Sprintf("%s %s;", rs.TokenLiteral(), rs.ReturnValue.String())
}
