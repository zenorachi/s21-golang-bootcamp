package ex00

import "math"

func MinCoins(val int, coins []int) []int {
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

func MinCoins2(val int, coins []int) (result []int) {
	if len(coins) == 0 || val <= 0 || coins == nil {
		return []int{}
	}

	sums := make([]int, val+1)
	for i := 1; i <= val; i++ {
		sums[i] = math.MaxUint32
	}

	for sum := 1; sum <= val; sum++ {
		for _, coin := range coins {
			if sum >= coin && sums[sum-coin]+1 < sums[sum] {
				sums[sum] = sums[sum-coin] + 1
			}
		}
	}

	if sums[val] == math.MaxInt {
		return []int{}
	}

	kek := val
	for kek > 0 {
		for _, coin := range coins {
			if kek >= coin && sums[kek-coin] == sums[kek]-1 {
				result = append(result, coin)
				kek -= coin
				break
			}
		}
	}

	return result
}
