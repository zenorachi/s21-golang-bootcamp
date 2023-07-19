package benchmark

import (
	"day07/ex00"
	"testing"
)

func BenchmarkMinCoins2(b *testing.B) {
	coins := []int{1, 3, 4, 7, 13, 15, 100, 5000, 10000, 1_000_000, 1_000_000_000}
	val := 100

	for i := 0; i < b.N; i++ {
		_ = ex00.MinCoins(val, coins)
	}
}
