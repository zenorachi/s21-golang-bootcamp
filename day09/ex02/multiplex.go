package main

import "sync"

func multiplex(input ...<-chan interface{}) chan interface{} {
	fanIn := make(chan interface{})
	wg := &sync.WaitGroup{}

	for _, ch := range input {
		wg.Add(1)
		go func(c <-chan interface{}) {
			defer wg.Done()
			for data := range c {
				fanIn <- data
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(fanIn)
	}()

	return fanIn
}
