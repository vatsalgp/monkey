package ast

// Building Block of Monkey Language
type Node interface {
	FirstTokenLiteral() string
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
