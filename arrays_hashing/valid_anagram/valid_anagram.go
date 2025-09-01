// Package validanagram
package validanagram

/* HOW TO SOLVE

* Check same length
* Create 2 hashmaps, one for each string
* Loop length of one string
  * add value to hash map as key
  * add number/increment value at key
* Loop
  * Check if both have the same number for each key

*/

func isValidAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	hmS1 := make(map[rune]int)
	hmS2 := make(map[rune]int)

	for i, c := range s1 {
		hmS1[c]++
		hmS2[rune(s2[i])]++
	}

	for _, c := range s1 {
		if hmS1[c] != hmS2[c] {
			return false
		}
	}

	return true
}
