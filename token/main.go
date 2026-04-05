package token

import (
	"regexp"
)

func CreateToken(word string) Token {
	for _, spec := range SpecsByPriority {
		if isMatch, _ := regexp.MatchString(string(spec.Expression), word); isMatch {
			return Token{Type: spec.Type, Literal: word}
		}
	}

	return Token{Type: ILLEGAL.Type, Literal: word}
}

func CreateTokenString(word string) string {
	return CreateToken(word).String()
}
