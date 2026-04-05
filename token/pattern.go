package token

type TokenPattern string

const (
	EndOfFileExpression            TokenPattern = `^$`
	LetExpression                  TokenPattern = `^let$`
	FunctionExpression             TokenPattern = `^fn$`
	ReturnExpression               TokenPattern = `^return$`
	PlusExpression                 TokenPattern = `^\+$`
	MinusExpression                TokenPattern = `^-$`
	MultiplyExpression             TokenPattern = `^\*$`
	DivideExpression               TokenPattern = `^/$`
	AssignExpression               TokenPattern = `^=$`
	IsEqualExpression              TokenPattern = `^==$`
	IsNotEqualExpression           TokenPattern = `^!=$`
	IsLessThanExpression           TokenPattern = `^<$`
	IsGreaterThanExpression        TokenPattern = `^>$`
	IsLessThanOrEqualExpression    TokenPattern = `^<=$`
	IsGreaterThanOrEqualExpression TokenPattern = `^>=$`
	CommaExpression                TokenPattern = `^,$`
	SemicolonExpression            TokenPattern = `^;$`
	LeftParenExpression            TokenPattern = `^\($`
	RightParenExpression           TokenPattern = `^\)$`
	LeftBraceExpression            TokenPattern = `^\{$`
	RightBraceExpression           TokenPattern = `^\}$`
	TrueExpression                 TokenPattern = `^true$`
	FalseExpression                TokenPattern = `^false$`
	IntegerLiteralExpression       TokenPattern = `^\d+$`
	FloatLiteralExpression         TokenPattern = `^\d+\.\d+$`
	StringLiteralExpression        TokenPattern = `^\".*\"$`
	IdentifierExpression           TokenPattern = `^[a-zA-Z_]+\w*$`
)
