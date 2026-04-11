package lexer

func isPeriod(ch byte) bool {
	return ch == '.'
}

func isUnderscore(ch byte) bool {
	return ch == '_'
}

func isEqual(ch byte) bool {
	return ch == '='
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlphabet(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r'
}
