package tree

type Node struct {
	HasToy bool
	Left   *Node
	Right  *Node
}

func Create(toy bool) *Node {
	return &Node{
		HasToy: toy,
		Left:   nil,
		Right:  nil,
	}
}

func AreToysBalanced(root *Node) bool {
	return getToysCount(root.Left) == getToysCount(root.Right)
}

func getToysCount(tree *Node) int {
	if tree == nil {
		return 0
	}

	if tree.HasToy {
		return getToysCount(tree.Left) + getToysCount(tree.Right) + 1
	}

	return getToysCount(tree.Left) + getToysCount(tree.Right)
}
