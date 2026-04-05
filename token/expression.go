package token

type Token struct {
	Type       TokenType
	Literal    string
	Expression string
	Create     func(word string) Token
}

type TokenType string

var (
	// SPECIAL
	ILLEGAL TokenType = "ILLEGAL"

	IllegalExpr = Token{
		Type:       ILLEGAL,
		Expression: "",
		Create: func(word string) Token {
			return Token{Type: ILLEGAL, Literal: string(ILLEGAL)}
		},
	}

	END_OF_FILE TokenType = "END_OF_FILE"

	EndOfFileExpr = Token{
		Type:       END_OF_FILE,
		Expression: `^$`,
		Create: func(word string) Token {
			return Token{Type: END_OF_FILE, Literal: string(END_OF_FILE)}
		},
	}

	// KEYWORDS
	LET TokenType = "LET"

	LetExpr = Token{
		Type:       LET,
		Expression: `^let$`,
		Create: func(word string) Token {
			return Token{Type: LET, Literal: string(LET)}
		},
	}

	FUNCTION TokenType = "FUNCTION"

	FunctionExpr = Token{
		Type:       FUNCTION,
		Expression: `^fn$`,
		Create: func(word string) Token {
			return Token{Type: FUNCTION, Literal: string(FUNCTION)}
		},
	}

	RETURN TokenType = "RETURN"

	ReturnExpr = Token{
		Type:       RETURN,
		Expression: `^return$`,
		Create: func(word string) Token {
			return Token{Type: RETURN, Literal: string(RETURN)}
		},
	}

	// IDENTIFIERS
	IDENTIFIER TokenType = "IDENTIFIER"

	IdentifierExpr = Token{
		Type:       IDENTIFIER,
		Expression: `^[a-zA-Z_]+\w*$`,
		Create: func(word string) Token {
			return Token{Type: IDENTIFIER, Literal: word}
		},
	}

	// LITERALS
	INTEGER_LITERAL TokenType = "INTEGER_LITERAL"

	IntegerLiteralExpr = Token{
		Type:       INTEGER_LITERAL,
		Expression: `^\d+$`,
		Create: func(word string) Token {
			return Token{Type: INTEGER_LITERAL, Literal: word}
		},
	}

	STRING_LITERAL TokenType = "STRING_LITERAL"

	StringLiteralExpr = Token{
		Type:       STRING_LITERAL,
		Expression: `^\".*\"$`,
		Create: func(word string) Token {
			return Token{Type: STRING_LITERAL, Literal: word}
		},
	}

	FLOAT_LITERAL TokenType = "FLOAT_LITERAL"

	FloatLiteralExpr = Token{
		Type:       FLOAT_LITERAL,
		Expression: `^\d+\.\d+$`,
		Create: func(word string) Token {
			return Token{Type: FLOAT_LITERAL, Literal: word}
		},
	}

	BOOLEAN_LITERAL TokenType = "BOOLEAN_LITERAL"

	TrueExpr = Token{
		Type:       BOOLEAN_LITERAL,
		Expression: `^true$`,
		Create: func(word string) Token {
			return Token{Type: BOOLEAN_LITERAL, Literal: "true"}
		},
	}

	FalseExpr = Token{
		Type:       BOOLEAN_LITERAL,
		Expression: `^false$`,
		Create: func(word string) Token {
			return Token{Type: BOOLEAN_LITERAL, Literal: "false"}
		},
	}

	// OPERATORS
	PLUS TokenType = "PLUS"

	PlusExpr = Token{
		Type:       PLUS,
		Expression: `^\+$`,
		Create: func(word string) Token {
			return Token{Type: PLUS, Literal: string(PLUS)}
		},
	}

	MINUS TokenType = "MINUS"

	MinusExpr = Token{
		Type:       MINUS,
		Expression: `^-$`,
		Create: func(word string) Token {
			return Token{Type: MINUS, Literal: string(MINUS)}
		},
	}

	MULTIPLY TokenType = "MULTIPLY"

	MultiplyExpr = Token{
		Type:       MULTIPLY,
		Expression: `^\*$`,
		Create: func(word string) Token {
			return Token{Type: MULTIPLY, Literal: string(MULTIPLY)}
		},
	}

	DIVIDE TokenType = "DIVIDE"

	DivideExpr = Token{
		Type:       DIVIDE,
		Expression: `^/$`,
		Create: func(word string) Token {
			return Token{Type: DIVIDE, Literal: string(DIVIDE)}
		},
	}

	ASSIGN TokenType = "ASSIGN"

	AssignExpr = Token{
		Type:       ASSIGN,
		Expression: `^=$`,
		Create: func(word string) Token {
			return Token{Type: ASSIGN, Literal: string(ASSIGN)}
		},
	}

	// COMPARISON OPERATORS
	IS_EQUAL TokenType = "IS_EQUAL"

	IsEqualExpr = Token{
		Type:       IS_EQUAL,
		Expression: `^==$`,
		Create: func(word string) Token {
			return Token{Type: IS_EQUAL, Literal: string(IS_EQUAL)}
		},
	}

	IS_NOT_EQUAL TokenType = "IS_NOT_EQUAL"

	IsNotEqualExpr = Token{
		Type:       IS_NOT_EQUAL,
		Expression: `^!=$`,
		Create: func(word string) Token {
			return Token{Type: IS_NOT_EQUAL, Literal: string(IS_NOT_EQUAL)}
		},
	}

	IS_LESS_THAN TokenType = "IS_LESS_THAN"

	IsLessThanExpr = Token{
		Type:       IS_LESS_THAN,
		Expression: `^<$`,
		Create: func(word string) Token {
			return Token{Type: IS_LESS_THAN, Literal: string(IS_LESS_THAN)}
		},
	}

	IS_GREATER_THAN TokenType = "IS_GREATER_THAN"

	IsGreaterThanExpr = Token{
		Type:       IS_GREATER_THAN,
		Expression: `^>$`,
		Create: func(word string) Token {
			return Token{Type: IS_GREATER_THAN, Literal: string(IS_GREATER_THAN)}
		},
	}

	IS_LESS_THAN_OR_EQUAL TokenType = "IS_LESS_THAN_OR_EQUAL"

	IsLessThanOrEqualExpr = Token{
		Type:       IS_LESS_THAN_OR_EQUAL,
		Expression: `^<=$`,
		Create: func(word string) Token {
			return Token{Type: IS_LESS_THAN_OR_EQUAL, Literal: string(IS_LESS_THAN_OR_EQUAL)}
		},
	}

	IS_GREATER_THAN_OR_EQUAL TokenType = "IS_GREATER_THAN_OR_EQUAL"

	IsGreaterThanOrEqualExpr = Token{
		Type:       IS_GREATER_THAN_OR_EQUAL,
		Expression: `^>=$`,
		Create: func(word string) Token {
			return Token{Type: IS_GREATER_THAN_OR_EQUAL, Literal: string(IS_GREATER_THAN_OR_EQUAL)}
		},
	}

	// DELIMITERS
	COMMA TokenType = "COMMA"

	CommaExpr = Token{
		Type:       COMMA,
		Expression: `^,$`,
		Create: func(word string) Token {
			return Token{Type: COMMA, Literal: string(COMMA)}
		},
	}

	SEMICOLON TokenType = "SEMICOLON"

	SemicolonExpr = Token{
		Type:       SEMICOLON,
		Expression: `^;$`,
		Create: func(word string) Token {
			return Token{Type: SEMICOLON, Literal: string(SEMICOLON)}
		},
	}

	// SYMBOLS
	LEFT_PAREN TokenType = "LEFT_PAREN"

	LeftParenExpr = Token{
		Type:       LEFT_PAREN,
		Expression: `^\($`,
		Create: func(word string) Token {
			return Token{Type: LEFT_PAREN, Literal: string(LEFT_PAREN)}
		},
	}

	RIGHT_PAREN TokenType = "RIGHT_PAREN"

	RightParenExpr = Token{
		Type:       RIGHT_PAREN,
		Expression: `^\)$`,
		Create: func(word string) Token {
			return Token{Type: RIGHT_PAREN, Literal: string(RIGHT_PAREN)}
		},
	}

	LEFT_BRACE TokenType = "LEFT_BRACE"

	LeftBraceExpr = Token{
		Type:       LEFT_BRACE,
		Expression: `^\{$`,
		Create: func(word string) Token {
			return Token{Type: LEFT_BRACE, Literal: string(LEFT_BRACE)}
		},
	}

	RIGHT_BRACE TokenType = "RIGHT_BRACE"

	RightBraceExpr = Token{
		Type:       RIGHT_BRACE,
		Expression: `^\}$`,
		Create: func(word string) Token {
			return Token{Type: RIGHT_BRACE, Literal: string(RIGHT_BRACE)}
		},
	}
)
