package main

/*
sliding window
no leetcode
TC: O(n)
SC: O(k)
*/

import "fmt"

func main() {

	nums := []int{1, 1, 2, 1, 1}
	fmt.Println(distintNumbersInWindow(nums, 2))

}

func distintNumbersInWindow(nums []int, k int) []int {

	m := make(map[int]int)
	res := []int{}

	for i := 0; i < k; i++ {
		v, ok := m[nums[i]]
		if ok {
			m[nums[i]] = v + 1
		} else {
			m[nums[i]] = 1
		}
	}

	res = append(res, len(m))

	for i := k; i < len(nums); i++ {

		// remove
		r := nums[i-k]
		v, ok := m[r]

		if ok {
			if v > 1 {
				m[r] = v - 1
			} else {
				delete(m, r)
			}
		}

		// add
		a := nums[i]
		v, ok = m[a]
		if ok {
			m[a] = v + 1
		} else {
			m[a] = 1
		}

		res = append(res, len(m))
	}

	return res
}
