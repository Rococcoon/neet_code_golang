// Package invertbinarytree
package invertbinarytree

import (
	"Rococcoon/neet_code_golang/trees"
)

func InvertTree(root *trees.TreeNode) *trees.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}
