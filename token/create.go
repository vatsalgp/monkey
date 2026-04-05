package token

import (
	"regexp"
)

func CreateToken(word string) Token {
	if isMatch, _ := regexp.MatchString(EndOfFileExpr.Expression, word); isMatch {
		return EndOfFileExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(LetExpr.Expression, word); isMatch {
		return LetExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(FunctionExpr.Expression, word); isMatch {
		return FunctionExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(ReturnExpr.Expression, word); isMatch {
		return ReturnExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(PlusExpr.Expression, word); isMatch {
		return PlusExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(MinusExpr.Expression, word); isMatch {
		return MinusExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(MultiplyExpr.Expression, word); isMatch {
		return MultiplyExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(DivideExpr.Expression, word); isMatch {
		return DivideExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(AssignExpr.Expression, word); isMatch {
		return AssignExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(IsEqualExpr.Expression, word); isMatch {
		return IsEqualExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(IsNotEqualExpr.Expression, word); isMatch {
		return IsNotEqualExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(IsLessThanExpr.Expression, word); isMatch {
		return IsLessThanExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(IsGreaterThanExpr.Expression, word); isMatch {
		return IsGreaterThanExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(IsLessThanOrEqualExpr.Expression, word); isMatch {
		return IsLessThanOrEqualExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(IsGreaterThanOrEqualExpr.Expression, word); isMatch {
		return IsGreaterThanOrEqualExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(CommaExpr.Expression, word); isMatch {
		return CommaExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(SemicolonExpr.Expression, word); isMatch {
		return SemicolonExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(LeftParenExpr.Expression, word); isMatch {
		return LeftParenExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(RightParenExpr.Expression, word); isMatch {
		return RightParenExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(LeftBraceExpr.Expression, word); isMatch {
		return LeftBraceExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(RightBraceExpr.Expression, word); isMatch {
		return RightBraceExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(TrueExpr.Expression, word); isMatch {
		return TrueExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(FalseExpr.Expression, word); isMatch {
		return FalseExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(IntegerLiteralExpr.Expression, word); isMatch {
		return IntegerLiteralExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(FloatLiteralExpr.Expression, word); isMatch {
		return FloatLiteralExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(StringLiteralExpr.Expression, word); isMatch {
		return StringLiteralExpr.Create(word)
	}

	if isMatch, _ := regexp.MatchString(IdentifierExpr.Expression, word); isMatch {
		return IdentifierExpr.Create(word)
	}

	return IllegalExpr.Create(word)
}
