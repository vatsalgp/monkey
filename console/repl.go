package console

import (
	"bufio"
	"fmt"
	"strings"
)

func printPrompt() {
	fmt.Print("Enter your token: ")
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return line
}

func printToken(createTokenString func(string) string, line string) {
	word := strings.TrimSpace(line)
	token := createTokenString(word)
	fmt.Printf("Word '%s' is %s\n\n", word, token)
}

func printGoodbye() {
	fmt.Println("\nGoodbye from Monkey!")
}

func printHello() {
	fmt.Printf("Hello from Monkey!\n\n")
}
