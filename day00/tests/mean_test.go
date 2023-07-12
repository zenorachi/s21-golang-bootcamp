package tests

import (
	"stats/statistics"
	"testing"
)

func TestMean(t *testing.T) {
	t.Run("Data from file", func(t *testing.T) {
		var sm statistics.StatMetrics
		sm.ReadFromFile("test.txt")
		mean := sm.Mean()
		realMean := 2.0
		if mean != realMean {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", mean, realMean)
		}
	})

	t.Run("Data from slice", func(t *testing.T) {
		var sm statistics.StatMetrics
		sm.SetArray([]int{1, 2, 3, 4, -5})
		mean := sm.Mean()
		realMean := 1.0
		if mean != realMean {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", mean, realMean)
		}
	})
}
