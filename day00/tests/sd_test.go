package tests

import (
	"math"
	"stats/statistics"
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	t.Run("Data from file", func(t *testing.T) {
		var sm statistics.StatMetrics
		sm.ReadFromFile("test.txt")
		sd := sm.StandardDeviation()
		realSD := 0.816
		if math.Abs(sd-realSD) > 1e-03 {
			t.Errorf("[Expected] %f != %f [Real]\n", sd, realSD)
		}
	})

	t.Run("Data from slice", func(t *testing.T) {
		var sm statistics.StatMetrics
		sm.SetArray([]int{1, 2, 3, 4, -5, 4, 0, 99, 314, 14113, 0, -991})
		sd := sm.StandardDeviation()
		realSD := 3926.201
		if math.Abs(sd-realSD) > 1e-03 {
			t.Errorf("[Expected] %f != %f [Real]\n", sd, realSD)
		}
	})
}
