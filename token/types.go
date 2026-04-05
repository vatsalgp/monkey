package token

type Type string

type Expression string

type Spec struct {
	Type       Type
	Expression Expression
}

type Token struct {
	Type    Type
	Literal string
}

func (token Token) String() string {
	return "Token(" + string(token.Type) + ", '" + string(token.Literal) + "')"
}

func (spec Spec) String() string {
	return "Spec(" + string(spec.Type) + ", '" + string(spec.Expression) + "')"
}
