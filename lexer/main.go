package lexer

import (
	"strings"
)

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func NewString(line string) string {
	sb := strings.Builder{}

	for tok := range New(line).All() {
		sb.WriteString(tok.String())
		sb.WriteByte('\n')
	}

	return sb.String()
}
