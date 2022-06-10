package main

import "fmt"

func main() {
	nums := []int{3, 34, 4, 12, 5, 2}
	target := 34

	fmt.Println(subsetSum(nums, target))
}

func subsetSum(nums []int, target int) bool {

	if len(nums) == 0 {
		return false
	}

	return subsetSumRecusrsive(nums, target, 0)
}

func subsetSumRecusrsive(nums []int, target int, pos int) bool {

	if target == 0 {
		return true
	}

	if target < 0 {
		return false
	}

	for i := pos; i < len(nums); i++ {
		if subsetSumRecusrsive(nums, target-nums[i], i+1) {
			return true
		}
	}

	return false

}
