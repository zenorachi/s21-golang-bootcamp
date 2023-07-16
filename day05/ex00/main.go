package main

import "fmt"

func main() {
	tree := &TreeNode{true, nil, nil}
	tree.Right = create(true)
	tree.Left = create(false)
	fmt.Println(areToysBalanced(tree))
}
