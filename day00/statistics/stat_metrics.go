package statistics

import "math"

func (sm *StatMetrics) Mean() float64 {
	var sum float64
	for _, num := range sm.arrayOfIntegers {
		sum += float64(num)
	}
	return sum / float64(sm.lengthOfArray)
}

func (sm *StatMetrics) Median() float64 {
	if sm.lengthOfArray > 0 {
		mid := sm.lengthOfArray / 2
		if sm.lengthOfArray%2 != 0 {
			return float64(sm.arrayOfIntegers[mid])
		}
		return float64(sm.arrayOfIntegers[mid]+sm.arrayOfIntegers[mid-1]) / 2
	}
	return math.NaN()
}

func (sm *StatMetrics) Mode() int {
	if sm.lengthOfArray > 0 {
		minElAbs := int(math.Abs(float64(sm.minElem)))
		size := sm.maxElem + minElAbs
		indices := make([]int, size+1, size+1)

		var index, max int
		for i := 0; i < sm.lengthOfArray; i++ {
			indices[sm.arrayOfIntegers[i]+minElAbs]++
			tmp := indices[sm.arrayOfIntegers[i]+minElAbs]
			if tmp > max {
				max = tmp
				index = i
			}
		}

		return sm.arrayOfIntegers[index]
	}
	return 0
}

func (sm *StatMetrics) StandardDeviation() float64 {
	if sm.lengthOfArray > 0 {
		average := sm.Mean()
		var sumOfResiduals float64

		for _, elem := range sm.arrayOfIntegers {
			sumOfResiduals += (float64(elem) - average) * (float64(elem) - average)
		}

		return math.Sqrt(sumOfResiduals / float64(sm.lengthOfArray))
	}
	return math.NaN()
}
