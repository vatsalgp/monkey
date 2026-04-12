package main

import (
	"github.com/vatsalgp/monkey/lexer"
	"github.com/vatsalgp/monkey/repl"
	"github.com/vatsalgp/monkey/token"
)

func main() {
	repl.Start(token.NewString, lexer.NewString)
}
