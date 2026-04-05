package lexer

import (
	"strings"

	"github.com/vatsalgp/monkey/token"
)

func CreateLexerString(line string) string {
	words := strings.Fields(line)

	tokens := make([]string, len(words))

	for i, word := range words {
		tokens[i] = token.CreateTokenString(word)
	}

	return strings.Join(tokens, "\n")
}
