package lexer

import (
	"github.com/vatsalgp/monkey/token"
)

func isComparitorStart(typ token.Type) bool {
	return typ == token.ASSIGN.Type || typ == token.IS_LESS_THAN.Type || typ == token.IS_GREATER_THAN.Type || typ == token.NOT.Type
}

func isIdentifierOrLiteralStart(typ token.Type) bool {
	return typ == token.INTEGER_LITERAL.Type || typ == token.IDENTIFIER.Type || typ == token.ILLEGAL.Type
}
