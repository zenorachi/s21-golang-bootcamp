package ph

import (
	"container/heap"
	"errors"
)

func getNCoolestPresents(presents []Present, n int) ([]Present, error) {
	if n > len(presents) || n < 0 {
		return []Present{}, errors.New("invalid argument")
	}

	ph := make(PresentHeap, len(presents))
	for i, present := range presents {
		ph[i].Value = present.Value
		ph[i].Size = present.Size
	}
	heap.Init(&ph)

	coolestPresents := make([]Present, n)
	for i := 0; i < n; i++ {
		present := heap.Pop(&ph).(Present)
		coolestPresents[i] = present
	}

	return coolestPresents, nil
}
