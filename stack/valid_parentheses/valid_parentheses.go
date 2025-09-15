// Package validparentheses
package validparentheses

import (
	"github.com/golang-collections/collections/stack"
)

func isValid(s string) bool {
	// Initialize a new stack.
	stack := stack.New()

	// Map closing brackets to their corresponding opening brackets.
	closeToOpenMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		// Check if the current character is a closing bracket.
		if open, exists := closeToOpenMap[c]; exists {
			// If it's a closing bracket, the stack must not be empty,
			// and the top of the stack must match the required opening bracket.
			if stack.Len() == 0 {
				return false
			}

			// Pop the top element and check for a match.
			// Note: The Pop() method from this library returns an interface{}, so we need to cast it back to rune.
			top := stack.Pop().(rune)
			if top != open {
				return false
			}

		} else {
			// If it's an opening bracket, push it onto the stack.
			stack.Push(c)
		}
	}

	// The string is valid only if the stack is empty at the end.
	return stack.Len() == 0
}
