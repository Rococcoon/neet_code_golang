package balancedbinarytree

import (
	"testing"

	"Rococcoon/neet_code_golang/trees"
)

func TestBalancedBinaryTreeDFS(t *testing.T) {
	tc := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "empty tree",
			input:    []int{},
			expected: true,
		},
		{
			name:     "single node",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "perfectly balanced",
			input:    []int{3, 9, 20, 0, 0, 15, 7},
			expected: true,
		},
		{
			name:     "unbalanced_1",
			input:    []int{1, 2, 2, 3, 3, 0, 0, 4, 4},
			expected: false,
		},
		{
			name:     "unbalanced_2",
			input:    []int{1, 2, 0, 3, 4, 0, 0, 5, 6},
			expected: false,
		},
		{
			name:     "left-heavy",
			input:    []int{5, 3, 7, 2, 4, 6, 8, 1},
			expected: true,
		},
		{
			name:     "right-heavy",
			input:    []int{5, 3, 7, 2, 4, 6, 8, 0, 1},
			expected: true,
		},
		{
			name:     "null values at various positions",
			input:    []int{1, 2, 0, 3},
			expected: false,
		},
		{
			name:     "large balanced tree",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: true,
		},
		{
			name:     "large unbalanced tree",
			input:    []int{1, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7},
			expected: false,
		},
		{
			name:     "single branch",
			input:    []int{1, 0, 2, 0, 3, 0, 4, 0, 5},
			expected: false,
		},
		{
			name:     "zigzag tree",
			input:    []int{1, 2, 3, 4, 0, 5, 0, 6},
			expected: false,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			inputTree := trees.NewBinaryTree(tc.input)

			result := isBalancedDFS(inputTree)

			if result != tc.expected {
				t.Errorf("For Test: %s, expected %t, but got %t", tc.name, tc.expected, result)
			} else {
				t.Logf("-------------------------\n")
				t.Logf("PASS: Test '%s'", tc.name)
				t.Logf("-------------------------\n")
			}
		})
	}
}

func TestBalancedBinaryTreeIDFS(t *testing.T) {
	tc := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "empty tree",
			input:    []int{},
			expected: true,
		},
		{
			name:     "single node",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "perfectly balanced",
			input:    []int{3, 9, 20, 0, 0, 15, 7},
			expected: true,
		},
		{
			name:     "unbalanced_1",
			input:    []int{1, 2, 2, 3, 3, 0, 0, 4, 4},
			expected: false,
		},
		{
			name:     "unbalanced_2",
			input:    []int{1, 2, 0, 3, 4, 0, 0, 5, 6},
			expected: false,
		},
		{
			name:     "left-heavy",
			input:    []int{5, 3, 7, 2, 4, 6, 8, 1},
			expected: true,
		},
		{
			name:     "right-heavy",
			input:    []int{5, 3, 7, 2, 4, 6, 8, 0, 1},
			expected: true,
		},
		{
			name:     "null values at various positions",
			input:    []int{1, 2, 0, 3},
			expected: false,
		},
		{
			name:     "large balanced tree",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: true,
		},
		{
			name:     "large unbalanced tree",
			input:    []int{1, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7},
			expected: false,
		},
		{
			name:     "single branch",
			input:    []int{1, 0, 2, 0, 3, 0, 4, 0, 5},
			expected: false,
		},
		{
			name:     "zigzag tree",
			input:    []int{1, 2, 3, 4, 0, 5, 0, 6},
			expected: false,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			inputTree := trees.NewBinaryTree(tc.input)

			result := isBalancedIDFS(inputTree)

			if result != tc.expected {
				t.Errorf("For Test: %s, expected %t, but got %t", tc.name, tc.expected, result)
			} else {
				t.Logf("-------------------------\n")
				t.Logf("PASS: Test '%s'", tc.name)
				t.Logf("-------------------------\n")
			}
		})
	}
}
