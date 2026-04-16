package ast

import (
	"fmt"

	"github.com/vatsalgp/monkey/token"
)

// Eg: "let x = 5 ;"
type LetStatement struct {
	Token *token.Token // let
	Name  *Identifier  // x
	Value Expression   // 5
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) String() string {
	return fmt.Sprintf("%s %s = %s;", ls.TokenLiteral(), ls.Name.TokenLiteral(), ls.Value.String())
}
