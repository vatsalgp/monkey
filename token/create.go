package token

import (
	"regexp"
)

func CreateToken(word string) Token {
	if isMatch, _ := regexp.MatchString(string(EndOfFileExpr.Expr), word); isMatch {
		return CreateEndOfFileToken()
	}

	if isMatch, _ := regexp.MatchString(string(LetExpr.Expr), word); isMatch {
		return CreateLetToken()
	}

	if isMatch, _ := regexp.MatchString(string(FunctionExpr.Expr), word); isMatch {
		return CreateFunctionToken()
	}

	if isMatch, _ := regexp.MatchString(string(ReturnExpr.Expr), word); isMatch {
		return CreateReturnToken()
	}

	if isMatch, _ := regexp.MatchString(string(PlusExpr.Expr), word); isMatch {
		return CreatePlusToken()
	}

	if isMatch, _ := regexp.MatchString(string(MinusExpr.Expr), word); isMatch {
		return CreateMinusToken()
	}

	if isMatch, _ := regexp.MatchString(string(MultiplyExpr.Expr), word); isMatch {
		return CreateMultiplyToken()
	}

	if isMatch, _ := regexp.MatchString(string(DivideExpr.Expr), word); isMatch {
		return CreateDivideToken()
	}

	if isMatch, _ := regexp.MatchString(string(AssignExpr.Expr), word); isMatch {
		return CreateAssignToken()
	}

	if isMatch, _ := regexp.MatchString(string(IsEqualExpr.Expr), word); isMatch {
		return CreateIsEqualToken()
	}

	if isMatch, _ := regexp.MatchString(string(IsNotEqualExpr.Expr), word); isMatch {
		return CreateIsNotEqualToken()
	}

	if isMatch, _ := regexp.MatchString(string(IsLessThanExpr.Expr), word); isMatch {
		return CreateIsLessThanToken()
	}

	if isMatch, _ := regexp.MatchString(string(IsGreaterThanExpr.Expr), word); isMatch {
		return CreateIsGreaterThanToken()
	}

	if isMatch, _ := regexp.MatchString(string(IsLessThanOrEqualExpr.Expr), word); isMatch {
		return CreateIsLessThanOrEqualToken()
	}

	if isMatch, _ := regexp.MatchString(string(IsGreaterThanOrEqualExpr.Expr), word); isMatch {
		return CreateIsGreaterThanOrEqualToken()
	}

	if isMatch, _ := regexp.MatchString(string(CommaExpr.Expr), word); isMatch {
		return CreateCommaToken()
	}

	if isMatch, _ := regexp.MatchString(string(SemicolonExpr.Expr), word); isMatch {
		return CreateSemicolonToken()
	}

	if isMatch, _ := regexp.MatchString(string(LeftParenExpr.Expr), word); isMatch {
		return CreateLeftParenToken()
	}

	if isMatch, _ := regexp.MatchString(string(RightParenExpr.Expr), word); isMatch {
		return CreateRightParenToken()
	}

	if isMatch, _ := regexp.MatchString(string(LeftBraceExpr.Expr), word); isMatch {
		return CreateLeftBraceToken()
	}

	if isMatch, _ := regexp.MatchString(string(RightBraceExpr.Expr), word); isMatch {
		return CreateRightBraceToken()
	}

	if isMatch, _ := regexp.MatchString(string(TrueExpr.Expr), word); isMatch {
		return CreateTrueToken()
	}

	if isMatch, _ := regexp.MatchString(string(FalseExpr.Expr), word); isMatch {
		return CreateFalseToken()
	}

	if isMatch, _ := regexp.MatchString(string(IntegerLiteralExpr.Expr), word); isMatch {
		return CreateIntegerLiteralToken(word)
	}

	if isMatch, _ := regexp.MatchString(string(FloatLiteralExpr.Expr), word); isMatch {
		return CreateFloatLiteralToken(word)
	}

	if isMatch, _ := regexp.MatchString(string(StringLiteralExpr.Expr), word); isMatch {
		return CreateStringLiteralToken(word)
	}

	if isMatch, _ := regexp.MatchString(string(IdentifierExpr.Expr), word); isMatch {
		return CreateIdentifierToken(word)
	}

	return CreateIllegalToken()
}
