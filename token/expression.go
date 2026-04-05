package token

type TokenExpression struct {
	Type TokenType
	Expr TokenPattern
}

var (
	EndOfFileExpr = TokenExpression{Type: END_OF_FILE, Expr: EndOfFileExpression}

	LetExpr      = TokenExpression{Type: LET, Expr: LetExpression}
	FunctionExpr = TokenExpression{Type: FUNCTION, Expr: FunctionExpression}
	ReturnExpr   = TokenExpression{Type: RETURN, Expr: ReturnExpression}

	PlusExpr     = TokenExpression{Type: PLUS, Expr: PlusExpression}
	MinusExpr    = TokenExpression{Type: MINUS, Expr: MinusExpression}
	MultiplyExpr = TokenExpression{Type: MULTIPLY, Expr: MultiplyExpression}
	DivideExpr   = TokenExpression{Type: DIVIDE, Expr: DivideExpression}
	AssignExpr   = TokenExpression{Type: ASSIGN, Expr: AssignExpression}

	IsEqualExpr              = TokenExpression{Type: IS_EQUAL, Expr: IsEqualExpression}
	IsNotEqualExpr           = TokenExpression{Type: IS_NOT_EQUAL, Expr: IsNotEqualExpression}
	IsLessThanExpr           = TokenExpression{Type: IS_LESS_THAN, Expr: IsLessThanExpression}
	IsGreaterThanExpr        = TokenExpression{Type: IS_GREATER_THAN, Expr: IsGreaterThanExpression}
	IsLessThanOrEqualExpr    = TokenExpression{Type: IS_LESS_THAN_OR_EQUAL, Expr: IsLessThanOrEqualExpression}
	IsGreaterThanOrEqualExpr = TokenExpression{Type: IS_GREATER_THAN_OR_EQUAL, Expr: IsGreaterThanOrEqualExpression}

	CommaExpr     = TokenExpression{Type: COMMA, Expr: CommaExpression}
	SemicolonExpr = TokenExpression{Type: SEMICOLON, Expr: SemicolonExpression}

	LeftParenExpr  = TokenExpression{Type: LEFT_PAREN, Expr: LeftParenExpression}
	RightParenExpr = TokenExpression{Type: RIGHT_PAREN, Expr: RightParenExpression}
	LeftBraceExpr  = TokenExpression{Type: LEFT_BRACE, Expr: LeftBraceExpression}
	RightBraceExpr = TokenExpression{Type: RIGHT_BRACE, Expr: RightBraceExpression}

	TrueExpr  = TokenExpression{Type: BOOLEAN_LITERAL, Expr: TrueExpression}
	FalseExpr = TokenExpression{Type: BOOLEAN_LITERAL, Expr: FalseExpression}

	IntegerLiteralExpr = TokenExpression{Type: INTEGER_LITERAL, Expr: IntegerLiteralExpression}
	FloatLiteralExpr   = TokenExpression{Type: FLOAT_LITERAL, Expr: FloatLiteralExpression}
	StringLiteralExpr  = TokenExpression{Type: STRING_LITERAL, Expr: StringLiteralExpression}
	IdentifierExpr     = TokenExpression{Type: IDENTIFIER, Expr: IdentifierExpression}
)
