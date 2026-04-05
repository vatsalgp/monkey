package token

import (
	"fmt"
)

type Token struct {
	Type    TokenType
	Literal TokenLiteral
}

func (token Token) String() string {
	return fmt.Sprintf("Token(%s, %s)", token.Type, token.Literal)
}
