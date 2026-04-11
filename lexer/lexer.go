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

func (lex *Lexer) readIdentifierOrLiteral() {
	if isAlphabet(lex.currChar) || isUnderscore(lex.currChar) {
		// Identifier or Keyword
		for isAlphabet(lex.currChar) || isUnderscore(lex.currChar) || isDigit(lex.currChar) {
			lex.readChar()
		}
	} else if isDigit(lex.currChar) || isPeriod(lex.currChar) {
		// Integer Literal or Float Literal
		for isDigit(lex.currChar) || isPeriod(lex.currChar) {
			lex.readChar()
		}
	}
}

func (lex *Lexer) createToken(start int) token.Token {
	return token.New(lex.input[start:lex.currPos])
}

func (lex *Lexer) NextToken() token.Token {
	lex.eatWhiteSpace()

	start := lex.currPos
	charTok := token.New(string(lex.currChar))

	if isComparitorStart(charTok.Type) && isEqual(lex.peekChar()) {
		lex.readChar() // For peeked char
		lex.readChar()
		return lex.createToken(start)
	}

	if isIdentifierOrLiteralStart(charTok.Type) {
		lex.readIdentifierOrLiteral()
		return lex.createToken(start)
	}

	lex.readChar()
	return charTok
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
