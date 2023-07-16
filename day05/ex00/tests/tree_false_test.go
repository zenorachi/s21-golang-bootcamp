package tests

import (
	tree "day05/ex00"
	"testing"
)

func TestTreeFalse1(t *testing.T) {
	t.Run("Three nodes", func(t *testing.T) {
		/*
			  1
			 /

			1   0
		*/
		root := &tree.Node{HasToy: true}
		root.Left = tree.Create(true)
		root.Right = tree.Create(false)
		areBalanced := tree.AreToysBalanced(root)

		if false != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})

	t.Run("Five nodes", func(t *testing.T) {
		/*
			  0
			 /

			1   0



			  1   1
		*/
		root := &tree.Node{HasToy: false}
		root.Left = tree.Create(true)
		root.Right = tree.Create(false)
		root.Right.Left = tree.Create(true)
		root.Right.Right = tree.Create(true)
		areBalanced := tree.AreToysBalanced(root)

		if false != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})

	t.Run("Seven nodes", func(t *testing.T) {
		/*
			    1
			   /

			  1     0
			 /
			       /

			1   0 1   0
		*/
		root := &tree.Node{HasToy: true}
		root.Left = tree.Create(true)
		root.Left.Left = tree.Create(true)
		root.Left.Right = tree.Create(false)
		root.Right = tree.Create(false)
		root.Right.Left = tree.Create(true)
		root.Right.Right = tree.Create(false)
		areBalanced := tree.AreToysBalanced(root)

		if false != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})
}
