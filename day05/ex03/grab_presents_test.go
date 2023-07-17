package knapsack

import (
	ph "day05/ex02"
	"testing"
)

func TestGrabPresents(t *testing.T) {
	presents := make([]ph.Present, 4)
	presents[0].Value = 5
	presents[0].Size = 1
	presents[1].Value = 4
	presents[1].Size = 5
	presents[2].Value = 3
	presents[2].Size = 1
	presents[3].Value = 5
	presents[3].Size = 2

	var profitPresents []ph.Present
	var maxSize int

	t.Run("MaxSize - 3", func(t *testing.T) {
		maxSize = 3
		profitPresents = grabPresents(presents, maxSize)
		expectedRes := []ph.Present{
			{5, 1},
			{5, 2},
		}
		for i, present := range profitPresents {
			if present != expectedRes[i] {
				t.Errorf("[Expected] %v != [Real] %v\n", expectedRes[i], present)
			}
		}
	})

	t.Run("MaxSize - 8", func(t *testing.T) {
		maxSize = 8
		profitPresents = grabPresents(presents, maxSize)
		expectedRes := []ph.Present{
			{5, 1},
			{4, 5},
			{5, 2},
		}
		for i, present := range profitPresents {
			if present != expectedRes[i] {
				t.Errorf("[Expected] %v != [Real] %v\n", expectedRes[i], present)
			}
		}
	})
}
