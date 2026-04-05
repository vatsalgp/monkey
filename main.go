package main

import (
	"fmt"

	"github.com/vatsalgp/monkey/token"
)

func main() {
	fmt.Printf("Hello from Monkey!\n\n")

	for {
		var s string
		fmt.Print("Enter your token: ")
		fmt.Scan(&s)
		fmt.Printf("String %s is %s\n\n", s, token.CreateToken(s))
	}
}
