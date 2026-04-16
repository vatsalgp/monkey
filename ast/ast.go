package ast

// Building Block of Monkey Language
type Node interface {
	TokenLiteral() string // returns the literal value of the first token
	String() string
}

// Node which produces value
type Expression interface {
	Node
	expressionNode() // for debugging purposes
}

// Node which does not produce value
type Statement interface {
	Node
	statementNode() // for debugging purposes
}
