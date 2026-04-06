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

func (lex *Lexer) NextToken() token.Token {
	tok := token.NewB(lex.currChar)
	lex.readChar()
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
