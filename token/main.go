package token

import (
	"regexp"
)

func New(word string) Token {
	for _, spec := range SpecsByPriority {
		if isMatch, _ := regexp.MatchString(string(spec.Expression), word); isMatch {
			return Token{Type: spec.Type, Literal: word}
		}
	}

	return Token{Type: ILLEGAL.Type, Literal: word}
}

func NewB(byt byte) Token {
	if byt == 0 {
		return New(string(""))
	}
	return New(string(byt))
}

func NewString(word string) string {
	return New(word).String()
}
