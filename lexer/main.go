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
	nextChar byte
}

func (lex *Lexer) readChar() {
	if lex.nextPos >= len(lex.input) {
		lex.currChar = 0
	} else {
		lex.currChar = lex.input[lex.nextPos]
	}

	lex.currPos = lex.nextPos
	lex.nextPos++

	if lex.nextPos >= len(lex.input) {
		lex.nextChar = 0
	} else {
		lex.nextChar = lex.input[lex.nextPos]
	}
}

func (lex *Lexer) NextToken() token.Token {
	for isWhiteSpace(lex.currChar) {
		lex.readChar()
	}

	tok := token.NewB(lex.currChar)
	start := lex.currPos

	if tok.Type == token.ASSIGN.Type || tok.Type == token.IS_LESS_THAN.Type || tok.Type == token.IS_GREATER_THAN.Type || tok.Type == token.NOT.Type {
		if isEqual(lex.nextChar) {
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

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func CreateLexerString(line string) string {
	lex := New(line)
	tokens := []token.Token{}

	for {
		tok := lex.NextToken()
		if tok.Type == token.END_OF_FILE.Type {
			break
		}
		tokens = append(tokens, tok)
	}

	results := make([]string, len(tokens))

	for i, word := range tokens {
		results[i] = word.String()
	}

	return strings.Join(results, "\n")
}
