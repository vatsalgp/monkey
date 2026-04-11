package lexer

import (
	"strings"

	"github.com/vatsalgp/monkey/token"
)

type Lexer struct {
	input    string
	currPos  int
	nextPos  int
	currChar byte
}

func (lex *Lexer) readChar() {
	if lex.nextPos >= len(lex.input) {
		lex.currChar = 0
	} else {
		lex.currChar = lex.input[lex.nextPos]
	}
	lex.currPos = lex.nextPos
	lex.nextPos++
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlphabet(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isPeriod(ch byte) bool {
	return ch == '.'
}

func isUnderscore(ch byte) bool {
	return ch == '_'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r'
}

func (lex *Lexer) NextToken() token.Token {
	for isWhiteSpace(lex.currChar) {
		lex.readChar()
	}

	tok := token.NewB(lex.currChar)
	if tok.Type == token.INTEGER_LITERAL.Type || tok.Type == token.IDENTIFIER.Type || tok.Type == token.ILLEGAL.Type {
		start := lex.currPos

		if isAlphabet(lex.currChar) || isUnderscore(lex.currChar) {
			// Identifier or Keyword
			for isAlphabet(lex.currChar) || isUnderscore(lex.currChar) || isDigit(lex.currChar) {
				lex.readChar()
			}
		} else if isDigit(lex.currChar) || isPeriod(lex.currChar) {
			// Integer or Float Literal
			for isDigit(lex.currChar) || isPeriod(lex.currChar) {
				lex.readChar()
			}
		}

		end := lex.currPos
		val := lex.input[start:end]
		return token.New(val)
	} else {
		lex.readChar()
	}
	return tok
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func CreateLexerString(line string) string {
	words := strings.Fields(line)

	tokens := make([]string, len(words))

	for i, word := range words {
		tokens[i] = token.NewString(word)
	}

	return strings.Join(tokens, "\n")
}
