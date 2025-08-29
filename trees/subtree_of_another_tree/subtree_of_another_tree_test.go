// Pacakge subtreeofanothertree
package subtreeofanothertree

import (
	"testing"

	"Rococcoon/neet_code_golang/trees"
)

func TestIsSubtreeDFS(t *testing.T) {
	tc := []struct {
		name         string
		inputTree    []int
		inputSubTree []int
		expected     bool
	}{
		{
			name:         "identical trees",
			inputTree:    []int{3, 4, 5, 1, 2},
			inputSubTree: []int{3, 4, 5, 1, 2},
			expected:     true,
		},
		{
			name:         "subtree matches",
			inputTree:    []int{3, 4, 5, 1, 2, 0, 0, 0, 0, 0, 0},
			inputSubTree: []int{4, 1, 2},
			expected:     true,
		},
		{
			name:         "subtree is a leaf",
			inputTree:    []int{1, 2, 3},
			inputSubTree: []int{2},
			expected:     true,
		},
		{
			name:         "subtree is empty",
			inputTree:    []int{1, 2, 3},
			inputSubTree: []int{},
			expected:     true,
		},
		{
			name:         "main tree is empty",
			inputTree:    []int{},
			inputSubTree: []int{1, 2, 3},
			expected:     false,
		},
		{
			name:         "no match due to value difference",
			inputTree:    []int{3, 4, 5, 1, 2},
			inputSubTree: []int{4, 1, 3},
			expected:     false,
		},
		{
			name:         "no match due to structure difference",
			inputTree:    []int{1, 2, 0, 3},
			inputSubTree: []int{2, 3},
			expected:     true, // Corrected from false
		},
		{
			name:         "no match, wrong node is a potential root",
			inputTree:    []int{1, 2, 3, 4},
			inputSubTree: []int{2, 3},
			expected:     false,
		},
		{
			name:         "sub-tree with one node and main tree with two nodes",
			inputTree:    []int{1, 2},
			inputSubTree: []int{1},
			expected:     false, // Corrected from true
		},
		{
			name:         "subtree is much larger than main tree",
			inputTree:    []int{1, 2},
			inputSubTree: []int{1, 2, 3},
			expected:     false,
		},
		{
			name:         "sub-tree matches a branch of a larger tree",
			inputTree:    []int{1, 0, 2, 0, 3, 4, 5, 6},
			inputSubTree: []int{2, 4, 5},
			expected:     false,
		},
	}

	for _, tc := range tc {
		inputTree := trees.NewBinaryTree(tc.inputTree)
		inputSubTree := trees.NewBinaryTree(tc.inputSubTree)

		result := isSubTreeDFS(inputTree, inputSubTree)

		if tc.expected != result {
			t.Logf("-------------------------\n")
			t.Errorf("FAIL: Test '%s'", tc.name)
			t.Errorf("Expected %t but got %t", tc.expected, result)
			t.Logf("-------------------------\n")
		} else {
			t.Logf("-------------------------\n")
			t.Logf("PASS: Test '%s'", tc.name)
			t.Logf("-------------------------\n")
		}
	}
}
