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

func (tok Token) String() string {
	return "Token(" + string(tok.Type) + ", '" + string(tok.Literal) + "')"
}

func (spec Spec) String() string {
	return "Spec(" + string(spec.Type) + ", '" + string(spec.Expression) + "')"
}
