package main

/*
func main() {

	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	fmt.Println(search(nums, target))
}
*/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
