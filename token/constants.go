package token

var (
	// SPECIAL
	ILLEGAL     = Spec{Type: "ILLEGAL", Expression: ""}
	END_OF_FILE = Spec{Type: "END_OF_FILE", Expression: `^$`}

	// KEYWORDS
	LET      = Spec{Type: "LET", Expression: `^let$`}
	FUNCTION = Spec{Type: "FUNCTION", Expression: `^fn$`}
	RETURN   = Spec{Type: "RETURN", Expression: `^return$`}

	// IDENTIFIERS
	IDENTIFIER = Spec{Type: "IDENTIFIER", Expression: `^[a-zA-Z_]+\w*$`}

	// LITERALS
	INTEGER_LITERAL = Spec{Type: "INTEGER_LITERAL", Expression: `^\d+$`}
	STRING_LITERAL  = Spec{Type: "STRING_LITERAL", Expression: `^\".*\"$`}
	FLOAT_LITERAL   = Spec{Type: "FLOAT_LITERAL", Expression: `^\d+\.\d+$`}

	// BOOLEANS
	TRUE  = Spec{Type: "TRUE", Expression: `^true$`}
	FALSE = Spec{Type: "FALSE", Expression: `^false$`}

	// OPERATORS
	PLUS     = Spec{Type: "PLUS", Expression: `^\+$`}
	MINUS    = Spec{Type: "MINUS", Expression: `^-$`}
	MULTIPLY = Spec{Type: "MULTIPLY", Expression: `^\*$`}
	DIVIDE   = Spec{Type: "DIVIDE", Expression: `^/$`}
	ASSIGN   = Spec{Type: "ASSIGN", Expression: `^=$`}

	// COMPARISON OPERATORS
	IS_EQUAL                 = Spec{Type: "IS_EQUAL", Expression: `^==$`}
	IS_NOT_EQUAL             = Spec{Type: "IS_NOT_EQUAL", Expression: `^!=$`}
	IS_LESS_THAN             = Spec{Type: "IS_LESS_THAN", Expression: `^<$`}
	IS_GREATER_THAN          = Spec{Type: "IS_GREATER_THAN", Expression: `^>$`}
	IS_LESS_THAN_OR_EQUAL    = Spec{Type: "IS_LESS_THAN_OR_EQUAL", Expression: `^<=$`}
	IS_GREATER_THAN_OR_EQUAL = Spec{Type: "IS_GREATER_THAN_OR_EQUAL", Expression: `^>=$`}

	// DELIMITERS
	COMMA     = Spec{Type: "COMMA", Expression: `^,$`}
	SEMICOLON = Spec{Type: "SEMICOLON", Expression: `^;$`}

	// SYMBOLS
	LEFT_PAREN  = Spec{Type: "LEFT_PAREN", Expression: `^\($`}
	RIGHT_PAREN = Spec{Type: "RIGHT_PAREN", Expression: `^\)$`}
	LEFT_BRACE  = Spec{Type: "LEFT_BRACE", Expression: `^\{$`}
	RIGHT_BRACE = Spec{Type: "RIGHT_BRACE", Expression: `^\}$`}

	// Array of TokenExpressions in priority order
	SpecsByPriority = []Spec{
		END_OF_FILE,
		LET,
		FUNCTION,
		RETURN,
		PLUS,
		MINUS,
		MULTIPLY,
		DIVIDE,
		ASSIGN,
		IS_EQUAL,
		IS_NOT_EQUAL,
		IS_LESS_THAN,
		IS_GREATER_THAN,
		IS_LESS_THAN_OR_EQUAL,
		IS_GREATER_THAN_OR_EQUAL,
		COMMA,
		SEMICOLON,
		LEFT_PAREN,
		RIGHT_PAREN,
		LEFT_BRACE,
		RIGHT_BRACE,
		TRUE,
		FALSE,
		INTEGER_LITERAL,
		FLOAT_LITERAL,
		STRING_LITERAL,
		IDENTIFIER,
		ILLEGAL,
	}
)
