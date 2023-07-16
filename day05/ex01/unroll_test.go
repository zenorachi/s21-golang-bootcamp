package unroll

import (
	tree "day05/ex00"
	"reflect"
	"testing"
)

func TestUnroll(t *testing.T) {
	root := &tree.Node{HasToy: true}
	root.Left = tree.Create(true)
	root.Right = tree.Create(false)
	root.Left.Left = tree.Create(true)
	root.Left.Right = tree.Create(false)
	root.Right.Left = tree.Create(true)
	root.Right.Right = tree.Create(true)
	expectedResult := []bool{true, true, false, true, true, false, true}
	myResult := unrollGarland(root)
	t.Run(
		"TestCase1", func(t *testing.T) {
			if !reflect.DeepEqual(expectedResult, myResult) {
				t.Errorf("[Expected] %v != %v [Real]\\n", expectedResult, myResult)
			}
		},
	)
}
