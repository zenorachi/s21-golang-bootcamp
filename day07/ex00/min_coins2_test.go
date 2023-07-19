package ex00

import (
	"fmt"
	"testing"
)

func TestMinCoins2(t *testing.T) {
	var (
		val   int
		coins []int
		//expectedResult []int
		//realResult     []int
	)

	//t.Run("Case#1", func(t *testing.T) {
	//	val = 13
	//	coins = []int{1, 5, 10}
	//	expectedResult = []int{10, 1, 1, 1}
	//
	//	realResult = minCoins2(val, coins)
	//	if !reflect.DeepEqual(expectedResult, realResult) {
	//		t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
	//	}
	//})
	//
	//t.Run("Case#2", func(t *testing.T) {
	//	val = 15
	//	coins = []int{1, 5, 10}
	//	expectedResult = []int{10, 5}
	//
	//	realResult = minCoins2(val, coins)
	//	if !reflect.DeepEqual(expectedResult, realResult) {
	//		t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
	//	}
	//})

	t.Run("Case#3", func(t *testing.T) {
		val = 3303
		coins = []int{1, 100, 9, 500, 1000, 100, 1, 500, 1000, 9, 5000, 9}
		//expectedResult = []int{1000, 1000, 1000, 100, 100, 100, 1, 1, 1}

		minCoins2(val, coins)
		fmt.Println(coins)
		//if !reflect.DeepEqual(expectedResult, realResult) {
		//	t.Errorf("[Expected] %v != %v [Real]\n", expectedResult, realResult)
		//}
	})
}
