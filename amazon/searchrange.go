package main

/*
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
TC: O(logn)
SC: O(1)
*/

/*
func main() {

	nums := []int{}
	target := 6

	fmt.Println(searchRange(nums, target))
}
*/

func searchRange(nums []int, target int) []int {

	l := binSearch(nums, target, true)
	r := binSearch(nums, target, false)

	return []int{l, r}
}

func binSearch(nums []int, target int, leftBais bool) int {

	l := 0
	h := len(nums) - 1
	m := -1

	for h >= l {
		mid := (h + l) / 2

		if nums[mid] == target {
			m = mid
			if leftBais {
				h = mid - 1
			} else {
				l = mid + 1
			}

		} else if nums[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return m
}
