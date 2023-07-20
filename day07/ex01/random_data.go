package cpu

import (
	"math/rand"
	"time"
)

func getRandomData(cap int) []int {
	rand.Seed(time.Now().UnixNano())
	res := make([]int, cap)
	for i := 0; i < cap; i++ {
		res[i] = rand.Int()
	}
	return res
}
