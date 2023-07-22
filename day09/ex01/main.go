package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	urls := make(chan string)

	go func() {
		for i := 0; i < 100; i++ {
			urls <- "https://www.example.com"
		}
		close(urls)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Crawling interrupted. Stopping gracefully...")
		cancel()
	}()

	bodies := crawlWeb(ctx, urls)
	for b := range bodies {
		fmt.Println(*b)
	}
}
