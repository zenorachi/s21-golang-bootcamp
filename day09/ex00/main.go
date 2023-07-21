package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		ch   chan int
		done = make(chan struct{})
		arr  []int
	)

	arr = getRandomData(10)
	fmt.Println(arr)

	ch = sleepSort(arr)
	go func() {
		for i := 0; i < len(arr); i++ {
			fmt.Println(<-ch)
		}
		done <- struct{}{}
	}()

	<-done
}

func sleepSort(arr []int) chan int {
	ch := make(chan int)
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		go func() {
			time.Sleep(time.Duration(val) * time.Second)
			ch <- val
		}()
	}
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
