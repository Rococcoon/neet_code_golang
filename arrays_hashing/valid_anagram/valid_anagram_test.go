// Package validanagram
package validanagram

import (
	"testing"
)

func TestIsValidAnagram(t *testing.T) {
	tc := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "valid anagram",
			s1:       "listen",
			s2:       "silent",
			expected: true,
		},
		{
			name:     "not an anagram - different letters",
			s1:       "hello",
			s2:       "world",
			expected: false,
		},
		{
			name:     "not an anagram - different length",
			s1:       "a",
			s2:       "aa",
			expected: false,
		},
		{
			name:     "valid anagram - with spaces",
			s1:       "go go go",
			s2:       "oog ggo",
			expected: true,
		},
		{
			name:     "empty strings",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "one empty string",
			s1:       "a",
			s2:       "",
			expected: false,
		},
		{
			name:     "different case",
			s1:       "Abc",
			s2:       "cba",
			expected: false, // Assuming case sensitivity, otherwise this would be true
		},
		{
			name:     "valid anagram - complex characters",
			s1:       "résumé",
			s2:       "smereú",
			expected: true,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidAnagram(tc.s1, tc.s2)

			if result != tc.expected {
				t.Logf("-------------------------\n")
				t.Errorf("Test: %s\n expected: %t\n recieve: %t\n", tc.name, tc.expected, result)
				t.Logf("-------------------------\n")
				t.Logf("PASS: Test '%s'", tc.name)
				t.Logf("-------------------------\n")
				t.Logf("-------------------------\n")
			} else {
				t.Logf("-------------------------\n")
				t.Logf("PASS: Test '%s'", tc.name)
				t.Logf("-------------------------\n")
			}
		})
	}
}
