package main

import "fmt"

func main() {

	k := 5
	nums := []int{1, 3, 5, 7, 9}

	fmt.Println(RotateArrayInplace(nums, k))
}

func RotateArray(nums []int, k int) []int {

	n := len(nums)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		index := (i + k) % n
		res[index] = nums[i]
	}

	return res
}

func RotateArrayInplace(nums []int, k int) []int {

	reverse := func(nums []int, l, h int) {
		for h >= l {
			tmp := nums[l]
			nums[l] = nums[h]
			nums[h] = tmp
			h--
			l++
		}

	}

	// reverse full array
	reverse(nums, 0, len(nums)-1)
	k = k % len(nums)

	// reverse first k
	reverse(nums, 0, k-1)

	// reverse next
	reverse(nums, k, len(nums)-1)

	return nums
}
