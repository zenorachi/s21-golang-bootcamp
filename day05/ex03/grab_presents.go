package knapsack

import (
	ph "day05/ex02"
	"math"
)

func grabPresents(presents []ph.Present, maxSize int) (profitPresents []ph.Present) {
	if maxSize <= 0 {
		return []ph.Present{}
	}

	weightsLen := len(presents)
	table := make([][]int, weightsLen+1)
	for i := range table {
		table[i] = make([]int, maxSize+1)
	}

	for i := 0; i <= weightsLen; i++ {
		for j := 0; j <= maxSize; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else {
				if presents[i-1].Size > j {
					table[i][j] = table[i-1][j]
				} else {
					table[i][j] = max(table[i-1][j], table[i-1][j-presents[i-1].Size]+presents[i-1].Value)
				}
			}
		}
	}

	tracePresents(table, weightsLen, maxSize, &presents, &profitPresents)
	return profitPresents
}

func tracePresents(table [][]int, i, j int, presents, profitPresents *[]ph.Present) {
	if table[i][j] == 0 {
		return
	}

	if table[i-1][j] == table[i][j] {
		tracePresents(table, i-1, j, presents, profitPresents)
	} else {
		tracePresents(table, i-1, j-(*presents)[i-1].Size, presents, profitPresents)
		*profitPresents = append(*profitPresents, (*presents)[i-1])
	}
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
