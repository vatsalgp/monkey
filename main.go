package main

import (
	"github.com/vatsalgp/monkey/console"
	"github.com/vatsalgp/monkey/lexer"
	"github.com/vatsalgp/monkey/token"
)

func main() {
	console.Start(token.NewString, lexer.CreateLexerString)
}
