// Package sametree
package sametree

import "Rococcoon/neet_code_golang/trees"

// Depth First Search

func isSameTreeDFS(p, q *trees.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return isSameTreeDFS(p.Left, q.Left) && isSameTreeDFS(p.Right, q.Right)
	}

	return false
}

// Iterative Depth First Search

func isSameTreeIDFS(p, q *trees.TreeNode) bool {
	type pair struct {
		first  *trees.TreeNode
		second *trees.TreeNode
	}

	stack := []pair{{p, q}}

	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		node1 := stack[lastIdx].first
		node2 := stack[lastIdx].second
		stack = stack[:lastIdx]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		stack = append(stack, pair{node1.Right, node2.Right})
		stack = append(stack, pair{node1.Left, node2.Left})
	}

	return true
}

// Breadth First Search

func isSameTreeBFS(p, q *trees.TreeNode) bool {
	queue1 := []*trees.TreeNode{p}
	queue2 := []*trees.TreeNode{q}

	for len(queue1) > 0 && len(queue2) > 0 {
		for i := len(queue1); i > 0; i-- {
			nodeP := queue1[0]
			nodeQ := queue2[0]
			queue1 = queue1[1:]
			queue2 = queue2[1:]

			if nodeP == nil && nodeQ == nil {
				continue
			}
			if nodeP == nil || nodeQ == nil || nodeP.Val != nodeQ.Val {
				return false
			}

			queue1 = append(queue1, nodeP.Left, nodeP.Right)
			queue2 = append(queue2, nodeQ.Left, nodeQ.Right)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0
}
