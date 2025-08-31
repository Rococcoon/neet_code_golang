// Package hasduplicate
package hasduplicate

import (
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	tc := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "contains duplicates",
			input:    []int{1, 2, 3, 2, 5},
			expected: true,
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: false,
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: false,
		},
		{
			name:     "all elements are duplicates",
			input:    []int{7, 7, 7, 7},
			expected: true,
		},
		{
			name:     "duplicates at start and end",
			input:    []int{1, 2, 3, 4, 1},
			expected: true,
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "zero in slice",
			input:    []int{0, 1, 0, 2},
			expected: true,
		},
		{
			name:     "large numbers",
			input:    []int{1000000, 2000000, 1000000},
			expected: true,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			result := hasDuplicate(tc.input)

			if result != tc.expected {
				t.Logf("-------------------------\n")
				t.Errorf("Test: %s\n expected: %t\n recieve: %t\n", tc.name, tc.expected, result)
				t.Logf("-------------------------\n")
			} else {
				t.Logf("-------------------------\n")
				t.Logf("PASS: Test '%s'", tc.name)
				t.Logf("-------------------------\n")
			}
		})
	}
}
