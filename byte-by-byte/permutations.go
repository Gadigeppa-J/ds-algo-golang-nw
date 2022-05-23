package main

/*
https://leetcode.com/problems/permutations/
TC: O(n! * n)
SC: O(1)
*/

/* func main() {
	nums := []int{1, 2, 3}

	fmt.Println(permute(nums))
} */

func permute(nums []int) [][]int {
	result := [][]int{}
	permuteHelp(0, nums, &result)
	return result
}

func permuteHelp(index int, nums []int, result *[][]int) {

	if index >= len(nums) {
		cp := make([]int, len(nums))

		for i, n := range nums {
			cp[i] = n
		}

		*result = append(*result, cp)
		return
	}

	for i := index; i < len(nums); i++ {
		swap(nums, i, index)
		permuteHelp(index+1, nums, result)
		swap(nums, i, index)
	}

	return
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
