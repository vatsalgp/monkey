package token

type Token struct {
	Type    TokenType
	Literal TokenLiteral
}

func (token Token) String() string {
	return "Token(" + string(token.Type) + ", '" + string(token.Literal) + "')"
}

func CreateTokenString(word string) string {
	return CreateToken(word).String()
}
