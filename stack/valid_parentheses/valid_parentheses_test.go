package validparentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tc := []struct {
		name     string
		input    string
		expected bool
	}{
		// Valid strings
		{"empty string", "", true},
		{"simple valid string", "()", true},
		{"valid nested string", "()[]{}", true},
		{"valid mixed nested string", "{[()]}", true},
		{"another valid nested string", "()[]{}{()}", true},
		{"valid with other characters", "a(b[c]d)e", true},

		// Invalid strings
		{"unclosed string", "(", false},
		{"unopened string", "]", false},
		{"unmatched string", "([)]", false},
		{"unclosed first element", "({[", false},
		{"unclosed string", "([", false},
		{"mismatched closing string", "(]", false},
		{"mismatched closing string", "([)]", false},
		{"incorrect order", ")(", false},
		{"invalid with other characters", "a(b[c)d]e", false},
	}

	for _, tc := range tc {
		result := isValid(tc.input)

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

	}

}
