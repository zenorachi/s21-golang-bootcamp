package unroll

import (
	tree "day05/ex00"
)

func unrollGarland(root *tree.Node) []bool {
	if root == nil {
		return []bool{}
	}

	queue := []*tree.Node{root}
	values := []bool{root.HasToy}
	isEvenLevel := true

	for len(queue) > 0 {
		levelSize := len(queue)

		for i, j := 0, levelSize; i < levelSize; i, j = i+1, j-1 {
			var curNode *tree.Node
			if isEvenLevel {
				curNode = queue[i]
				if curNode.Left != nil {
					queue = append(queue, curNode.Left)
					values = append(values, curNode.Left.HasToy)
				}
				if curNode.Right != nil {
					queue = append(queue, curNode.Right)
					values = append(values, curNode.Right.HasToy)
				}
			} else {
				curNode = queue[j-1]
				if curNode.Right != nil {
					queue = append(queue, curNode.Right)
					values = append(values, curNode.Right.HasToy)
				}
				if curNode.Left != nil {
					queue = append(queue, curNode.Left)
					values = append(values, curNode.Left.HasToy)
				}
			}
		}

		isEvenLevel = !isEvenLevel
		queue = queue[levelSize:]
	}
	return values
}
