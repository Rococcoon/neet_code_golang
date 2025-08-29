// Package subtreeofanothertree
package subtreeofanothertree

import "Rococcoon/neet_code_golang/trees"

func isSubTreeDFS(root *trees.TreeNode, subRoot *trees.TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if sameTree(root, subRoot) {
		return true
	}

	return isSubTreeDFS(root.Left, subRoot) || isSubTreeDFS(root.Right, subRoot)
}

func sameTree(root *trees.TreeNode, subRoot *trees.TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root != nil && subRoot != nil && root.Val == subRoot.Val {
		return sameTree(root.Left, subRoot.Left) && sameTree(root.Right, subRoot.Right)
	}

	return false
}
