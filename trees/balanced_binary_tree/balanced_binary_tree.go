// Package balancedbinarytree
package balancedbinarytree

import (
	"Rococcoon/neet_code_golang/trees"
)

// Depth First Search

type result struct {
	balanced bool
	height   int
}

func isBalancedDFS(root *trees.TreeNode) bool {
	return dfs(root).balanced
}

func dfs(root *trees.TreeNode) result {
	if root == nil {
		return result{true, 0}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	balanced := left.balanced && right.balanced && abs(left.height-right.height) <= 1

	return result{balanced, 1 + max(left.height, right.height)}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Iterative Depth First Search

func isBalancedIDFS(root *trees.TreeNode) bool {
	stack := []*trees.TreeNode{}
	node := root
	last := (*trees.TreeNode)(nil)
	depth := make(map[*trees.TreeNode]int)

	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			if node.Right == nil || last == node.Right {
				stack = stack[:len(stack)-1]
				left := depth[node.Left]
				right := depth[node.Right]

				if abs(left-right) > 1 {
					return false
				}

				depth[node] = 1 + max(left, right)
				last = node
				node = nil
			} else {
				node = node.Right
			}
		}
	}
	return true
}
