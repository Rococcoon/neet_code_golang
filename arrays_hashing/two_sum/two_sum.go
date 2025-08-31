/*  HOW TO SOLVE

* Initialize a has map of map[int]int
* Get the difference
* Loop over map
* Check if it is in the hashmap
* if so return []int{first instance index of nums, second instande index of nums}
* if not append nums value as key, and index as value
* return empty []int{}

 */

package twosum

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if j, found := hm[diff]; found {
			return []int{j, i}
		}
		hm[n] = i
	}

	return []int{}
}
