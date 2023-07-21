package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const MaxConcurrency = 8

func crawlWeb(ctx context.Context, input chan string, done chan struct{}) chan *string {
	semaphore := make(chan struct{}, MaxConcurrency)
	output := make(chan *string)
	var wg sync.WaitGroup

	go func() {
		defer close(output)
		for {
			select {
			case url, ok := <-input:
				if !ok {
					// input channel closed, no more URLs to process
					wg.Wait() // Wait for all goroutines to complete
					done <- struct{}{}
					return
				}

				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					semaphore <- struct{}{} // acquire semaphore

					body, err := fetchURL(url)
					if err == nil {
						output <- &body
					}

					<-semaphore // release semaphore
				}(url)

			case <-ctx.Done():
				// Cancel the crawling process
				return
			}
		}
	}()

	return output
}

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func handleOutput(output chan *string) {
	for bodyPtr := range output {
		if bodyPtr != nil {
			body := *bodyPtr
			fmt.Printf("Received body:%s\n", body)
		}
	}
}
