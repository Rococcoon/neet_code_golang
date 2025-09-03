package validpalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tc := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single character", "a", true},
		{"simple palindrome", "madam", true},
		{"simple non-palindrome", "hello", false},
		{"palindrome with mixed case", "Racecar", true},
		{"non-palindrome with mixed case", "Go", false},
		{"palindrome with spaces", "a man a plan a canal panama", true},
		{"non-palindrome with spaces", "not a palindrome", false},
		{"palindrome with punctuation and numbers", "A man, a plan, a canal: Panama!", true},
		{"palindrome with a number string", "12321", true},
		{"non-palindrome with numbers", "12345", false},
		{"complex case with punctuation and mixed case", "Eva, can I see bees in a cave?", true},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.input)

			if result != tc.expected {
				t.Logf("-------------------------\n")
				t.Errorf("For Test: %s", tc.name)
				t.Errorf("Expected: %t", tc.expected)
				t.Errorf("Got: %t", result)
				t.Logf("-------------------------\n")
			} else {
				t.Logf("-------------------------\n")
				t.Logf("PASS: Test '%s'", tc.name)
				t.Logf("-------------------------\n")
			}
		})
	}
}
