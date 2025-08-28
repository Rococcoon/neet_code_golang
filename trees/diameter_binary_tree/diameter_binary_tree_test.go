package diameterbinarytree

import (
	"testing"

	"Rococcoon/neet_code_golang/trees"
)

func TestDiameterBinaryTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1: Basic Tree",
			input:    []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "Example 2: Skewed Tree",
			input:    []int{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			name:     "Empty Tree",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Single Node Tree",
			input:    []int{1},
			expected: 0,
		},
		{
			name:     "Complex Tree with Nulls",
			input:    []int{4, 2, 7, 1, 3, 6, 9},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputTree := trees.NewBinaryTree(tc.input)

			resultDiameter := diameterOfBinaryTree(inputTree)

			if resultDiameter != tc.expected {
				t.Errorf("For Test: %s, expected %d but got %d", tc.name, tc.expected, resultDiameter)
			} else {
				t.Logf("-------------------------\n")
				t.Logf("PASS: Test '%s'", tc.name)
				t.Logf("-------------------------\n")
			}
		})
	}
}
