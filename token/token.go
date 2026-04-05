package token

func CreateIllegalToken() Token {
	return Token{Type: ILLEGAL, Literal: TokenLiteral(ILLEGAL)}
}

func CreateEndOfFileToken() Token {
	return Token{Type: END_OF_FILE, Literal: TokenLiteral(END_OF_FILE)}
}

func CreateLetToken() Token {
	return Token{Type: LET, Literal: TokenLiteral(LET)}
}

func CreateFunctionToken() Token {
	return Token{Type: FUNCTION, Literal: TokenLiteral(FUNCTION)}
}

func CreateReturnToken() Token {
	return Token{Type: RETURN, Literal: TokenLiteral(RETURN)}
}

func CreatePlusToken() Token {
	return Token{Type: PLUS, Literal: TokenLiteral(PLUS)}
}

func CreateMinusToken() Token {
	return Token{Type: MINUS, Literal: TokenLiteral(MINUS)}
}

func CreateMultiplyToken() Token {
	return Token{Type: MULTIPLY, Literal: TokenLiteral(MULTIPLY)}
}

func CreateDivideToken() Token {
	return Token{Type: DIVIDE, Literal: TokenLiteral(DIVIDE)}
}

func CreateAssignToken() Token {
	return Token{Type: ASSIGN, Literal: TokenLiteral(ASSIGN)}
}

func CreateIsEqualToken() Token {
	return Token{Type: IS_EQUAL, Literal: TokenLiteral(IS_EQUAL)}
}

func CreateIsNotEqualToken() Token {
	return Token{Type: IS_NOT_EQUAL, Literal: TokenLiteral(IS_NOT_EQUAL)}
}

func CreateIsLessThanToken() Token {
	return Token{Type: IS_LESS_THAN, Literal: TokenLiteral(IS_LESS_THAN)}
}

func CreateIsGreaterThanToken() Token {
	return Token{Type: IS_GREATER_THAN, Literal: TokenLiteral(IS_GREATER_THAN)}
}

func CreateIsLessThanOrEqualToken() Token {
	return Token{Type: IS_LESS_THAN_OR_EQUAL, Literal: TokenLiteral(IS_LESS_THAN_OR_EQUAL)}
}

func CreateIsGreaterThanOrEqualToken() Token {
	return Token{Type: IS_GREATER_THAN_OR_EQUAL, Literal: TokenLiteral(IS_GREATER_THAN_OR_EQUAL)}
}

func CreateCommaToken() Token {
	return Token{Type: COMMA, Literal: TokenLiteral(COMMA)}
}

func CreateSemicolonToken() Token {
	return Token{Type: SEMICOLON, Literal: TokenLiteral(SEMICOLON)}
}

func CreateLeftParenToken() Token {
	return Token{Type: LEFT_PAREN, Literal: TokenLiteral(LEFT_PAREN)}
}

func CreateRightParenToken() Token {
	return Token{Type: RIGHT_PAREN, Literal: TokenLiteral(RIGHT_PAREN)}
}

func CreateLeftBraceToken() Token {
	return Token{Type: LEFT_BRACE, Literal: TokenLiteral(LEFT_BRACE)}
}

func CreateRightBraceToken() Token {
	return Token{Type: RIGHT_BRACE, Literal: TokenLiteral(RIGHT_BRACE)}
}

func CreateTrueToken() Token {
	return Token{Type: BOOLEAN_LITERAL, Literal: TokenLiteral("true")}
}

func CreateFalseToken() Token {
	return Token{Type: BOOLEAN_LITERAL, Literal: TokenLiteral("false")}
}

func CreateIdentifierToken(value string) Token {
	return Token{Type: IDENTIFIER, Literal: TokenLiteral(value)}
}

func CreateIntegerLiteralToken(value string) Token {
	return Token{Type: INTEGER_LITERAL, Literal: TokenLiteral(value)}
}

func CreateFloatLiteralToken(value string) Token {
	return Token{Type: FLOAT_LITERAL, Literal: TokenLiteral(value)}
}

func CreateStringLiteralToken(value string) Token {
	return Token{Type: STRING_LITERAL, Literal: TokenLiteral(value)}
}
