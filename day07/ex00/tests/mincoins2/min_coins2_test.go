package mincoins2

import (
	"day07/ex00"
	"reflect"
	"testing"
)

func TestMinCoins2Correct(t *testing.T) {
	var (
		val            int
		coins          []int
		expectedResult []int
		realResult     []int
	)

	t.Run("Case#1", func(t *testing.T) {
		val = 13
		coins = []int{1, 5, 10}
		expectedResult = []int{1, 1, 1, 10}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})

	t.Run("Case#2", func(t *testing.T) {
		val = 15
		coins = []int{1, 5, 10}
		expectedResult = []int{5, 10}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})
}

func TestMinCoins2Incorrect(t *testing.T) {
	var (
		val            int
		coins          []int
		expectedResult []int
		realResult     []int
	)

	t.Run("Case#1", func(t *testing.T) {
		val = 12
		coins = []int{1, 6, 10}
		expectedResult = []int{6, 6}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})

	t.Run("Case#2", func(t *testing.T) {
		val = 24
		coins = []int{1, 2, 12, 18}
		expectedResult = []int{12, 12}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})

	t.Run("Case#3", func(t *testing.T) {
		val = 25
		coins = []int{10, 5, 1}
		expectedResult = []int{10, 10, 5}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})

	t.Run("Case#4", func(t *testing.T) {
		val = 6
		coins = []int{2, 3, 4, 5, 10, 100, 1000000}
		expectedResult = []int{2, 4}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})

	t.Run("Case#5", func(t *testing.T) {
		val = 15
		coins = []int{1, 3, 4, 7, 13, 15}
		expectedResult = []int{15}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})

	t.Run("Case#6", func(t *testing.T) {
		val = -15
		coins = []int{1, 3, 4, 7, 13, 15}
		expectedResult = []int{}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})

	t.Run("Case#7", func(t *testing.T) {
		val = -15
		coins = []int{}
		expectedResult = []int{}

		realResult = ex00.MinCoins2(val, coins)
		if !reflect.DeepEqual(expectedResult, realResult) {
			t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		}
	})
}
