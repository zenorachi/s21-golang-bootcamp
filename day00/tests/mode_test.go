package tests

import (
	"stats/statistics"
	"testing"
)

func TestMode(t *testing.T) {
	t.Run("Data from file", func(t *testing.T) {
		var sm statistics.StatMetrics
		sm.ReadFromFile("big_test.txt")
		mode := sm.Mode()
		realMode := 1
		if mode != realMode {
			t.Errorf("[Expected] %d != %d [Real]\n", mode, realMode)
		}
	})

	t.Run("Data from slice", func(t *testing.T) {
		var sm statistics.StatMetrics
		sm.SetArray([]int{1, 2, 3, 4, -5, 4})
		mode := sm.Mode()
		realMode := 4
		if mode != realMode {
			t.Errorf("[Expected] %d != %d [Real]\n", mode, realMode)
		}
	})
}
