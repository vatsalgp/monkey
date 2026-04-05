package console

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Start(createTokenString func(string) string, createLexerString func(string) string) {
	printMenu()

	// Listen for interrupt signals to gracefully exit
	reader := bufio.NewReader(os.Stdin)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	option := 0

	for {
		fmt.Print("Selection: ")

		fmt.Scan(&option)

		if option != 0 && option != 1 && option != 2 {
			fmt.Println("Please enter a valid integer option (0 - 2).")
			continue
		}

		break
	}

	fmt.Println()

	for {

		switch option {
		case 1:
			printTokenPrompt()
		case 2:
			printLexerPrompt()
		case 0:
			printGoodbye()
			return
		default:
			fmt.Println("Invalid option selected. Please select either 0 - 2.")
		}

		// Read user input in a separate goroutine to allow for graceful shutdown
		lineReady := make(chan string, 1)

		// This goroutine will read a line from the console and send the result to the lineReady channel
		go func() {
			lineReady <- readLine(reader)
		}()

		select {
		// If the user sends an interrupt signal (e.g., Ctrl+C), we print a goodbye message and exit
		case <-ctx.Done():
			printGoodbye()
			return

			// If a line is ready from the console, we process it and print the corresponding token
		case result := <-lineReady:
			if result == "" {
				printGoodbye()
				return
			}

			// Actual work is done here
			switch option {
			case 1:
				printToken(createTokenString, result)
			case 2:
				printLexer(createLexerString, result)
			case 0:
				printGoodbye()
				return
			default:
				fmt.Println("Invalid option selected. Please select either 0 - 2.")
			}
		}
	}
}
