package lexer

import (
	"iter"

	"github.com/vatsalgp/monkey/token"
)

type Lexer struct {
	input    string
	currPos  int
	nextPos  int
	currChar byte
}

func (lex *Lexer) peekChar() byte {
	if lex.nextPos >= len(lex.input) {
		return 0
	}

	return lex.input[lex.nextPos]
}

func (lex *Lexer) readChar() {
	lex.currChar = lex.peekChar()
	lex.currPos = lex.nextPos
	lex.nextPos++
}

func (lex *Lexer) eatWhiteSpace() {
	for isWhiteSpace(lex.currChar) {
		lex.readChar()
	}
}

func (lex *Lexer) NextToken() token.Token {
	lex.eatWhiteSpace()

	tok := token.New(string(lex.currChar))
	start := lex.currPos

	if tok.Type == token.ASSIGN.Type || tok.Type == token.IS_LESS_THAN.Type || tok.Type == token.IS_GREATER_THAN.Type || tok.Type == token.NOT.Type {
		if isEqual(lex.peekChar()) {
			lex.readChar()
			lex.readChar()
			end := lex.currPos
			val := lex.input[start:end]
			return token.New(val)
		}
	}

	if tok.Type == token.INTEGER_LITERAL.Type || tok.Type == token.IDENTIFIER.Type || tok.Type == token.ILLEGAL.Type {
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

// Yields tokens until END_OF_FILE.
func (lex *Lexer) All() iter.Seq[token.Token] {
	return func(yield func(token.Token) bool) {
		for {
			tok := lex.NextToken()

			if tok.Type == token.END_OF_FILE.Type {
				return
			}

			if !yield(tok) {
				return
			}
		}
	}
}
