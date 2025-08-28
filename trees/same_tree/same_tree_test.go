package sametree

import (
	"testing"

	"Rococcoon/neet_code_golang/trees"
)

func TestIsSameTreeDFS(t *testing.T) {
	tc := []struct {
		name     string
		input1   []int
		input2   []int
		expected bool
	}{
		{
			name:     "both empty",
			input1:   []int{},
			input2:   []int{},
			expected: true,
		},
		{
			name:     "one empty, one not",
			input1:   []int{},
			input2:   []int{1},
			expected: false,
		},
		{
			name:     "simple identical trees",
			input1:   []int{1, 2, 3},
			input2:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "same values, different structure",
			input1:   []int{1, 2, 0, 3},
			input2:   []int{1, 0, 2, 0, 3},
			expected: false,
		},
		{
			name:     "different values at same position",
			input1:   []int{1, 2, 3},
			input2:   []int{1, 3, 2},
			expected: false,
		},
		{
			name:     "perfectly identical trees",
			input1:   []int{1, 2, 3, 4, 5, 6, 7},
			input2:   []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "nulls at different positions",
			input1:   []int{1, 2, 0, 3},
			input2:   []int{1, 0, 2, 3},
			expected: false,
		},
		{
			name:     "one tree is a subtree of the other",
			input1:   []int{1, 2},
			input2:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "mismatched tree shapes and values",
			input1:   []int{1, 0, 2, 3},
			input2:   []int{1, 2, 0, 3},
			expected: false,
		},
		{
			name:     "symmetric but not same",
			input1:   []int{1, 2, 2},
			input2:   []int{1, 2, 2},
			expected: true,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			inputTree1 := trees.NewBinaryTree(tc.input1)
			inputTree2 := trees.NewBinaryTree(tc.input2)

			result := isSameTreeDFS(inputTree1, inputTree2)

			if tc.expected != isSameTreeDFS(inputTree1, inputTree2) {
				t.Errorf("For Test: %s, expected %t but got %t", tc.name, tc.expected, result)
			}
		})
	}
}

func TestIsSameTreeIDFS(t *testing.T) {
	tc := []struct {
		name     string
		input1   []int
		input2   []int
		expected bool
	}{
		{
			name:     "both empty",
			input1:   []int{},
			input2:   []int{},
			expected: true,
		},
		{
			name:     "one empty, one not",
			input1:   []int{},
			input2:   []int{1},
			expected: false,
		},
		{
			name:     "simple identical trees",
			input1:   []int{1, 2, 3},
			input2:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "same values, different structure",
			input1:   []int{1, 2, 0, 3},
			input2:   []int{1, 0, 2, 0, 3},
			expected: false,
		},
		{
			name:     "different values at same position",
			input1:   []int{1, 2, 3},
			input2:   []int{1, 3, 2},
			expected: false,
		},
		{
			name:     "perfectly identical trees",
			input1:   []int{1, 2, 3, 4, 5, 6, 7},
			input2:   []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "nulls at different positions",
			input1:   []int{1, 2, 0, 3},
			input2:   []int{1, 0, 2, 3},
			expected: false,
		},
		{
			name:     "one tree is a subtree of the other",
			input1:   []int{1, 2},
			input2:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "mismatched tree shapes and values",
			input1:   []int{1, 0, 2, 3},
			input2:   []int{1, 2, 0, 3},
			expected: false,
		},
		{
			name:     "symmetric but not same",
			input1:   []int{1, 2, 2},
			input2:   []int{1, 2, 2},
			expected: true,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			inputTree1 := trees.NewBinaryTree(tc.input1)
			inputTree2 := trees.NewBinaryTree(tc.input2)

			result := isSameTreeIDFS(inputTree1, inputTree2)

			if tc.expected != isSameTreeDFS(inputTree1, inputTree2) {
				t.Errorf("For Test: %s, expected %t but got %t", tc.name, tc.expected, result)
			}
		})
	}
}

func TestIsSameTreeBFS(t *testing.T) {
	tc := []struct {
		name     string
		input1   []int
		input2   []int
		expected bool
	}{
		{
			name:     "both empty",
			input1:   []int{},
			input2:   []int{},
			expected: true,
		},
		{
			name:     "one empty, one not",
			input1:   []int{},
			input2:   []int{1},
			expected: false,
		},
		{
			name:     "simple identical trees",
			input1:   []int{1, 2, 3},
			input2:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "same values, different structure",
			input1:   []int{1, 2, 0, 3},
			input2:   []int{1, 0, 2, 0, 3},
			expected: false,
		},
		{
			name:     "different values at same position",
			input1:   []int{1, 2, 3},
			input2:   []int{1, 3, 2},
			expected: false,
		},
		{
			name:     "perfectly identical trees",
			input1:   []int{1, 2, 3, 4, 5, 6, 7},
			input2:   []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "nulls at different positions",
			input1:   []int{1, 2, 0, 3},
			input2:   []int{1, 0, 2, 3},
			expected: false,
		},
		{
			name:     "one tree is a subtree of the other",
			input1:   []int{1, 2},
			input2:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "mismatched tree shapes and values",
			input1:   []int{1, 0, 2, 3},
			input2:   []int{1, 2, 0, 3},
			expected: false,
		},
		{
			name:     "symmetric but not same",
			input1:   []int{1, 2, 2},
			input2:   []int{1, 2, 2},
			expected: true,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			inputTree1 := trees.NewBinaryTree(tc.input1)
			inputTree2 := trees.NewBinaryTree(tc.input2)

			result := isSameTreeBFS(inputTree1, inputTree2)

			if tc.expected != isSameTreeDFS(inputTree1, inputTree2) {
				t.Errorf("For Test: %s, expected %t but got %t", tc.name, tc.expected, result)
			}
		})
	}
}
