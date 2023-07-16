package main

import (
	"testing"
)

func TestCoolestPresents(t *testing.T) {
	presents := make([]Present, 4)
	presents[0].Value = 5
	presents[0].Size = 1
	presents[1].Value = 4
	presents[1].Size = 5
	presents[2].Value = 3
	presents[2].Size = 1
	presents[3].Value = 5
	presents[3].Size = 2

	var coolest []Present
	var err error

	t.Run("Two coolest presents", func(t *testing.T) {
		coolest, err = getNCoolestPresents(presents, 2)
		realCoolest := []Present{
			{5, 1},
			{5, 2},
		}
		for i, present := range realCoolest {
			if present != coolest[i] {
				t.Errorf("[Expected] %v != [Real] %v\n", realCoolest, coolest)
			}
		}
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Four coolest presents", func(t *testing.T) {
		coolest, err = getNCoolestPresents(presents, 4)
		realCoolest := []Present{
			{5, 1},
			{5, 2},
			{4, 5},
			{3, 1},
		}
		for i, present := range realCoolest {
			if present != coolest[i] {
				t.Errorf("[Expected] %v != [Real] %v\n", realCoolest, coolest)
			}
		}
		if err != nil {
			t.Error(err)
		}
	})

}
