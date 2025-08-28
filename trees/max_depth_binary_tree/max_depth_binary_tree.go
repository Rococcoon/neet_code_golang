// Package maxdepthbinarytree
package maxdepthbinarytree

import "Rococcoon/neet_code_golang/trees"

func maxDepth(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), max(maxDepth(root.Right)))
}
