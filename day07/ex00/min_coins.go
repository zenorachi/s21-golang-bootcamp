package ex00

import (
	"fmt"
	"sort"
)

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func minCoins2(val int, coins []int) []int {
	if len(coins) == 0 {
		return []int{}
	}

	sort.Ints(coins)
	removeDuplicates(coins)
	return nil
}

func removeDuplicates(coins []int) {
	length := len(coins)

	if length <= 1 {
		return
	}

	i := 0

	for j := 1; j < length; j++ {
		if coins[i] != coins[j] {
			i++
			coins[i] = coins[j]
		}
	}
	fmt.Println(coins)
}
