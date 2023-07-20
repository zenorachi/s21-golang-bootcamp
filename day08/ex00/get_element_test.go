package ptr

import (
	"testing"
)

func TestGetElementValid(t *testing.T) {
	var (
		arr []int
		idx int
	)

	t.Run("Case#1", func(t *testing.T) {
		arr = []int{1, 5, 10, 100, 1000, 10000, 100000, 1000000, 1000000000}
		for i := 0; i < len(arr); i++ {
			idx = i
			value, err := getElement(arr, idx)
			if err != nil {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if value != arr[i] {
				t.Errorf("[Expected] - %v != [Real] - %v\n", arr[i], value)
			}
		}
	})

	t.Run("Case#2", func(t *testing.T) {
		arr = []int{1000000000, 1, 5, 10, 100, 1000, 10000, 100000, 1000000, 0, 0, 0, -1, 1000000, 438943, -43141398}
		for i := 0; i < len(arr); i++ {
			idx = i
			value, err := getElement(arr, idx)
			if err != nil {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if value != arr[i] {
				t.Errorf("[Expected] - %v != [Real] - %v\n", arr[i], value)
			}
		}
	})

	t.Run("Case#3", func(t *testing.T) {
		arr = []int{1000000000, 1, 51000000, 438943, -43141398, 1, 1, 1, 1, 1, 1, 2, 3, 1, 43, 431}
		for i := 0; i < len(arr); i++ {
			idx = i
			value, err := getElement(arr, idx)
			if err != nil {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if value != arr[i] {
				t.Errorf("[Expected] - %v != [Real] - %v\n", arr[i], value)
			}
		}
	})
}

func TestGetElementInvalid(t *testing.T) {
	var (
		arr []int
		idx int
	)

	t.Run("Case#1", func(t *testing.T) {
		arr = []int{1, 5, 10, 100, 1000, 10000, 100000, 1000000, 1000000000}
		idx = -1
		_, err := getElement(arr, idx)
		if err == nil {
			t.Errorf("Unexpected error: %v\n", err)
		}
	})

	t.Run("Case#2", func(t *testing.T) {
		arr = []int{1, 5, 10, 100, 1000, 10000, 100000, 1000000, 1000000000}
		idx = 10000
		_, err := getElement(arr, idx)
		if err == nil {
			t.Errorf("Unexpected error: %v\n", err)
		}
	})

	t.Run("Case#3", func(t *testing.T) {
		arr = []int{}
		idx = 0
		_, err := getElement(arr, idx)
		if err == nil {
			t.Errorf("Unexpected error: %v\n", err)
		}
	})

	t.Run("Case#4", func(t *testing.T) {
		idx = 1
		_, err := getElement(nil, idx)
		if err == nil {
			t.Errorf("Unexpected error: %v\n", err)
		}
	})
}
