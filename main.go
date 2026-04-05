package main

import (
	"fmt"
	"github.com/vatsalgp/monkey/token"
)

func main() {
	fmt.Println("Hello from Monkey!")
	four := token.Token{Type: token.INTEGER_LITERAL, Literal: "4"}
	fmt.Println(four)
}
