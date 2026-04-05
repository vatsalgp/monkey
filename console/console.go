package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func tokenize(reader *bufio.Reader, CreateTokenString func(string) string) {
	fmt.Print("Enter your token: ")

	line, _ := reader.ReadString('\n')
	word := strings.TrimSpace(line)

	token := CreateTokenString(word)
	fmt.Printf("Word '%s' is %s\n\n", word, token)
}

func Start(CreateTokenString func(string) string) {
	fmt.Printf("Hello from Monkey!\n\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		tokenize(reader, CreateTokenString)
	}
}
