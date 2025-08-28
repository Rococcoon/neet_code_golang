package invertbinarytree

import (
	"Rococcoon/neet_code_golang/trees"
	"testing"
)

// function to check if output tree is that same as the expected output tree
func isSameTree(p, q *trees.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestInvertTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{4, 2, 7, 1, 3, 6, 9},
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "Example 2",
			input:    []int{2, 1, 3},
			expected: []int{2, 3, 1},
		},
		{
			name:     "Example 3 - Empty Tree",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single Node Tree",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Tree with a null child",
			input:    []int{3, 1, 0, 0, 2},
			expected: []int{3, 0, 1, 2, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputTree := trees.NewBinaryTree(tc.input)
			expectedTree := trees.NewBinaryTree(tc.expected)

			resultTree := InvertTree(inputTree)
			if !isSameTree(resultTree, expectedTree) {
				t.Errorf("For input tree %v, expected %v but got %v", tc.input, tc.expected, resultTree)
			}
		})
	}
}
