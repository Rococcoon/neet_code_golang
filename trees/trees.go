// Package trees contains Helper functions and an api for the tree
// pacakges to use
package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))
	for i, val := range nums {
		if val != 0 {
			nodes[i] = &TreeNode{Val: val}
		}
	}

	root := nodes[0]
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		current := queue[0]
		queue = queue[1:]

		// Assign left child
		if nodes[i] != nil {
			current.Left = nodes[i]
			queue = append(queue, current.Left)
		}
		i++

		// Assign right child
		if i < len(nums) && nodes[i] != nil {
			current.Right = nodes[i]
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}
