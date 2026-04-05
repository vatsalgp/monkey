package console

import (
	"bufio"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Start(createTokenString func(string) string) {
	printHello()

	// Listen for interrupt signals to gracefully exit
	reader := bufio.NewReader(os.Stdin)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for {
		// Read user input in a separate goroutine to allow for graceful shutdown
		lineReady := make(chan string, 1)

		// This goroutine will read a line from the console and send the result to the lineReady channel
		go func() {
			lineReady <- readLine(reader)
		}()

		printPrompt()

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
			printToken(createTokenString, result)
		}
	}
}
