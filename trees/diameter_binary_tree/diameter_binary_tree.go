// Package diameterbinarytree
package diameterbinarytree

import "Rococcoon/neet_code_golang/trees"

func diameterOfBinaryTree(root *trees.TreeNode) int {
	result := 0

	var dfs func(*trees.TreeNode) int
	dfs = func(root *trees.TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		result = max(result, left+right)

		return 1 + max(left, right)
	}

	dfs(root)
	return result
}
