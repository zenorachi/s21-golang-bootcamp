package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		ch  chan int
		arr []int
	)

	arr = getRandomData(10)
	fmt.Println(arr)

	ch = sleepSort(arr)
	for i := range ch {
		fmt.Println(i)
	}
}

func sleepSort(arr []int) chan int {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(val) * time.Second)
			ch <- val
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func getRandomData(cap int) []int {
	rand.Seed(time.Now().UnixNano())
	res := make([]int, cap)
	for i := 0; i < cap; i++ {
		res[i] = rand.Intn(7)
	}
	return res
}
