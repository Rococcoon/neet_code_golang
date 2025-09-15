// Package validpalindrome
package validpalindrome

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}

		if toLowerCase(rune(s[l])) != toLowerCase(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c rune) bool {
	if 'A' <= c && c <= 'Z' {
		return true
	}

	if 'a' <= c && c <= 'z' {
		return true
	}

	if '0' <= c && c <= '9' {
		return true
	}

	return false
}

func toLowerCase(c rune) rune {
	if 'A' <= c && c <= 'Z' {
		return c - ('A' - 'a')
	}

	return c
}
