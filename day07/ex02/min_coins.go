/*
How to reproduce docs generation:
	First way:
		1. Install the godoc util: run $ go get golang.org/x/tools/cmd/godoc or $ make install
		2. Run the command: $ make or $ ~/go/bin/godoc -http=:6060
		3. Go to http://localhost:6060/pkg/day07/ex02, and you will see the documentation for the mincoins package.
	Second way:
		1. Run the command: $ make extract-doc
		2. Open index.html
*/

// Package mincoins
package mincoins

import "math"

/*
MinCoins - default function, which is given as an example and works incorrectly with some cases.

For example:
coins = []int{1, 2, 12, 18}, val = 24. Function output will be {18, 1, 1, 1, 1}, although it should be {12, 12}.

This is because MinCoins uses a greedy algorithm that takes the largest coin denomination if it does not exceed the val value.
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

Some reasons, why MinCoins2 is better than MinCoins:

1. MinCoins2 is faster.
       The algorithm in the MinCoins2 function performs the computation only once for each sum from 1 to val.
       While the algorithm in the MinCoins function uses two nested loops to calculate the sum at each iteration.
       This makes the MinCoins2 algorithm more time efficient.
2. Dynamic programming.
       The algorithm in the MinCoins2 function uses the dynamic programming method to find the optimal solution.
       It builds an array of sums that contains the minimum number of coins for each sum from 1 to val.
       This avoids repeated computations and solves the problem more efficiently,
       especially for large values of val and complex coin sets.
3. Correctness.
       The algorithm in the MinCoins2 function guarantees that an optimal solution is found.
       It checks that the optimal number of coins for a given val exists before continuing the computation.
       If no optimal solution exists, it will return an empty slice.

Algorithm explanation:

1. Check for incorrect input.

2. Create sums slice.
	Create sums slice. of length val+1, where sums[i] will contain the minimum number of coins for the set of sum i.
	Initialize all array elements with the value math.MaxUint32 to make sure
	that the elements for which no optimal set of coins is found will remain with the maximum value.

3. Filling the sums slice.
	This loop considers each sum 'sum' from 1 to val and for each sum iterates over all coins in the set coins.
	For each sum 'sum', we check if some coin from the set of coins can improve the current sums[sum]
	for the set of that sum. If it can, we update sums[sum] to sums[sum-coin]+1,
	which means that we have found a better solution for that sum.

4. Check that we really can find the optimal solution.

5. Reconstruct the coins.
	This loop reconstructs the optimal set of coins for a given amount val.
	We start with val and at each iteration we select the coin 'coin' that caused the minimum number of coins
	to decrease for the current amount val.
	We then add that coin to the result and reduce val by its denomination to move on to the next sum.
	In the end, the result variable will contain the optimal set of coins for the set of sum val,
	and the minimum number of coins for that sum will be sums[val].

6. Return the result.
*/
func MinCoins2(val int, coins []int) (result []int) {
	// Check for incorrect input.
	if len(coins) == 0 || val <= 0 || coins == nil {
		return []int{}
	}

	/*
		Create sums slice of length val+1,
		where sums[i] will contain the minimum number of coins for the set of sum i.
		Initialize all array elements with the value math.MaxUint32 to make sure
		that the elements for which no optimal set of coins is found will remain with the maximum value.
	*/
	sums := make([]int, val+1)
	for i := 1; i <= val; i++ {
		sums[i] = math.MaxUint32
	}

	/*
		This loop considers each sum 'sum' from 1 to val and for each sum iterates over all coins in the set coins.
		For each sum 'sum', we check if some coin from the set of coins can improve the current sums[sum]
		for the set of that sum. If it can, we update sums[sum] to sums[sum-coin]+1,
		which means that we have found a better solution for that sum.
	*/
	for sum := 1; sum <= val; sum++ {
		for _, coin := range coins {
			if sum >= coin && sums[sum-coin]+1 < sums[sum] {
				sums[sum] = sums[sum-coin] + 1
			}
		}
	}

	// Check that we really can find the optimal solution.
	if sums[val] == math.MaxUint32 {
		return []int{}
	}

	/*
		This loop reconstructs the optimal set of coins for a given amount val.
		We start with val and at each iteration we select the coin 'coin' that caused the minimum number of coins
		to decrease for the current amount val.
		We then add that coin to the result and reduce val by its denomination to move on to the next sum.
		In the end, the result variable will contain the optimal set of coins for the set of sum val,
		and the minimum number of coins for that sum will be sums[val].
	*/
	for val > 0 {
		for _, coin := range coins {
			if val >= coin && sums[val-coin] == sums[val]-1 {
				result = append(result, coin)
				val -= coin
				break
			}
		}
	}

	// Return the result
	return result
}
