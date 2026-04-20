package token

type Precedence int

const (
	_ Precedence = iota

	LOWEST
	EQUALS            // ==
	LESSER_OR_GREATER // < or >
	SUM               // +
	PRODUCT           // *
	PREFIX            // -X or !X
	CALL              // myFunction(X)
)
