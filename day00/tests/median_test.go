package tests

import (
	"stats/statistics"
	"testing"
)

func TestMedian(t *testing.T) {
	t.Run("Data from file", func(t *testing.T) {
		var sm statistics.StatMetrics
		sm.ReadFromFile("test.txt")
		median := sm.Median()
		realMedian := 2.0
		if median != realMedian {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", median, realMedian)
		}
	})

	t.Run("Data from slice", func(t *testing.T) {
		var sm statistics.StatMetrics
		sm.SetArray([]int{1, 2, 3, 4, -5})
		mean := sm.Median()
		realMean := 2.0
		if mean != realMean {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", mean, realMean)
		}
	})
}
