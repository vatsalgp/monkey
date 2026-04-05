package token

func (token Token) String() string {
	return "Token(" + string(token.Type) + ", '" + string(token.Literal) + "')"
}

func CreateTokenString(word string) string {
	return CreateToken(word).String()
}
