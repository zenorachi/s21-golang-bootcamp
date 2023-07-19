/*
How to reproduce docs generation:
	First way:
		1. Install the godoc util: go get golang.org/x/tools/cmd/godoc
		2. Run the command: $ ~/go/bin/godoc -http=:6060 or $ make dvi
		3. Go to http://localhost:6060/pkg/day07/ex02, and you will see the documentation for the mincoins package.
	Second way:
		1. Run the command: $ make extract-docs
		2. Open index.html
*/

// Package mincoins
package mincoins

import "math"

/*
MinCoins - default function, which is given as an example and works incorrectly with some cases.

For example:
coins = []int{1, 2, 12, 18}, val = 24. Function output will be {18, 1, 1, 1, 1}, although it should be {12, 12}.

This is because it uses a greedy algorithm that takes the largest coin denomination if it does not exceed the val value.
In addition to such cases, cases are not processed when coins = nil or coins = {}.
*/
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

/*
MinCoins2 - new version of MinCoins.
It runs a bit faster and handles cases that the first version of it doesn't handle (you can see these cases in ex00/tests).
*/
func MinCoins2(val int, coins []int) (result []int) {
	// Check for incorrect input.
	if len(coins) == 0 || val <= 0 || coins == nil {
		return []int{}
	}

	// Create the slice to store... TODO
	sums := make([]int, val+1)
	for i := 1; i <= val; i++ {
		sums[i] = math.MaxUint32
	}

	// For loop for... TODO
	for sum := 1; sum <= val; sum++ {
		for _, coin := range coins {
			if sum >= coin && sums[sum-coin]+1 < sums[sum] {
				sums[sum] = sums[sum-coin] + 1
			}
		}
	}

	// Check that we really can find the optimal solution.
	if sums[val] == math.MaxInt {
		return []int{}
	}

	// For loop for... TODO
	for val > 0 {
		for _, coin := range coins {
			if val >= coin && sums[val-coin] == sums[val]-1 {
				result = append(result, coin)
				val -= coin
				break
			}
		}
	}

	return result
}
