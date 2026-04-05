package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vatsalgp/monkey/token"
)

func readAndProcessToken(reader *bufio.Reader) {
	fmt.Print("Enter your token: ")

	line, _ := reader.ReadString('\n')
	word := strings.TrimSpace(line)

	token := token.CreateToken(word)
	fmt.Printf("Word '%s' is %s\n\n", word, token)
}

func main() {
	fmt.Printf("Hello from Monkey!\n\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		readAndProcessToken(reader)
	}
}
