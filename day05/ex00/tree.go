package main

import "errors"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func create(toy bool) *TreeNode {
	return &TreeNode{
		HasToy: toy,
		Left:   nil,
		Right:  nil,
	}
}

func areToysBalanced(root *TreeNode) (bool, error) {
	if root == nil {
		return false, errors.New("root pointer cannot be nil")
	}

	return getToysCount(root.Left) == getToysCount(root.Right), nil
}

func getToysCount(tree *TreeNode) int {
	if tree == nil {
		return 0
	}

	if tree.HasToy {
		return getToysCount(tree.Left) + getToysCount(tree.Right) + 1
	}

	return getToysCount(tree.Left) + getToysCount(tree.Right)
}
