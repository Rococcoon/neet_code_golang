// package maxdepthbinarytree
package maxdepthbinarytree

import (
	"testing"

	"Rococcoon/neet_code_golang/trees"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1: Balanced Tree",
			input:    []int{3, 9, 20, 0, 0, 15, 7},
			expected: 3,
		},
		{
			name:     "Example 2: Skewed Tree",
			input:    []int{1, 0, 2, 0, 0, 0, 3},
			expected: 2,
		},
		{
			name:     "Empty Tree",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Single Node Tree",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "Tree with Nulls",
			input:    []int{1, 2, 3, 4, 0, 0, 5},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputTree := trees.NewBinaryTree(tc.input)

			resultMax := maxDepth(inputTree)

			if resultMax != tc.expected {
				t.Errorf("For Test: %s, expected %d but got %d", tc.name, tc.expected, resultMax)
			} else {
				t.Logf("PASS: Test '%s'", tc.name)
			}
		})
	}
}
