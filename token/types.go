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
