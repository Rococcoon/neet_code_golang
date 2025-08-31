package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tc := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "basic case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "no solution",
			nums:     []int{3, 2, 4},
			target:   10,
			expected: nil, // Or an empty slice, depending on function spec
		},
		{
			name:     "target is first element",
			nums:     []int{3, 2, 4},
			target:   5,
			expected: []int{0, 1},
		},
		{
			name:     "target is last element",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "zero in array",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "duplicate numbers",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "large numbers",
			nums:     []int{1000, 2000, 3000},
			target:   3000,
			expected: []int{0, 1},
		},
		{
			name:     "negative target",
			nums:     []int{5, 2, -3, 8},
			target:   -1,
			expected: []int{1, 2},
		},
		{
			name:     "single element array (edge case)",
			nums:     []int{5},
			target:   5,
			expected: nil, // Or an empty slice, since there's no pair
		},
		{
			name:     "long array with solution at ends",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   11,
			expected: []int{0, 9},
		},
		{
			name:     "duplicate target number",
			nums:     []int{1, 5, 9, 5},
			target:   10,
			expected: []int{1, 3},
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.nums, tc.target)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Logf("-------------------------\n")
				t.Errorf("Test: %s\n expected: %d\n recieve: %d\n", tc.name, tc.expected, result)
				t.Logf("-------------------------\n")
			} else {
				t.Logf("-------------------------\n")
				t.Logf("PASS: Test '%s'", tc.name)
				t.Logf("-------------------------\n")
			}
		})
	}
}
