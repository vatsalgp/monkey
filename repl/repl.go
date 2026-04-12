package repl

import (
	"bufio"
	"fmt"
	"strings"
)

func printTokenPrompt() {
	fmt.Print("Enter your token: ")
}

func printLexerPrompt() {
	fmt.Print("Enter your statement: ")
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return line
}

func printToken(createTokenString func(string) string, line string) {
	word := strings.TrimSpace(line)
	tok := createTokenString(word)
	fmt.Printf("Word '%s' is %s\n\n", word, tok)
}

func printLexer(createLexerString func(string) string, line string) {
	word := strings.TrimSpace(line)
	lexer := createLexerString(word)
	fmt.Println(lexer)
}

func printGoodbye() {
	fmt.Println("\nGoodbye from Monkey!")
}

func printMenu() {
	fmt.Println("Hello from Monkey!")
	fmt.Println()
	fmt.Println("1. Tokenizeing")
	fmt.Println("2. Lexing")
	fmt.Println("0. Exit")
	fmt.Println("Press Ctrl+D to exit")
	fmt.Println()
	fmt.Println("Enter (0-2) to select an option:")
}
