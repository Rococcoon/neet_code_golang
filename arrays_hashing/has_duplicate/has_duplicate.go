// Package hasduplicate...
package hasduplicate

/* HOW TO SOLVE

* initialize an empty hash map
  * map[int]bool
* loop through nums array
* for each check if value is in the hashmap
* if true
  * return true
* if false
  * Add value to hash map
* return false

*/

func hasDuplicate(nums []int) bool {
	hmSeen := make(map[int]bool)

	for _, n := range nums {
		if hmSeen[n] {
			return true
		}
		hmSeen[n] = true
	}

	return false
}
