package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Prepare a context to manage the goroutines and enable cancellation
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	input := make(chan string)
	done := make(chan struct{})

	// Start the crawling process and receive the output channel
	output := crawlWeb(ctx, input, done)

	// Start handling the output concurrently
	go handleOutput(output)

	// Add URLs to the input channel for crawling
	urls := []string{
		"https://example.com",
		"https://example.com/page1",
		"https://example.com/page2",
		// Add more URLs as needed...
	}

	for _, url := range urls {
		input <- url
	}

	close(input)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		fmt.Println("Context is done. Crawling process completed.")
	case <-sigChan:
		fmt.Println("\nTerminating the crawling process...")
		cancelFunc()
	case <-done:
	}

	// Wait for the handleOutput goroutine to complete
	// to make sure all the results have been printed
	<-output
}
