package main

/*
Given an array containing all the numbers from 1 to n except two, find the two missing numbers.
missing([4, 2, 3]) = 1, 5

TC: O(n)
SC: O(1)

*/

/*
func main() {
	fmt.Println(twoMissinNums([]int{1, 2}))
}
*/

func twoMissinNums(nums []int) []int {

	size := len(nums) + 2
	totalSum := (size * (size + 1)) / 2
	arrSum := 0

	for _, n := range nums {
		arrSum += n
	}

	pivot := (totalSum - arrSum) / 2

	leftTotalXor := 0
	rightTotalXor := 0

	for i := 1; i <= pivot; i++ {
		leftTotalXor = leftTotalXor ^ i
	}

	for i := pivot + 1; i <= size; i++ {
		rightTotalXor = rightTotalXor ^ i
	}

	leftArrXor := 0
	rightArrXor := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] <= pivot {
			leftArrXor = leftArrXor ^ nums[i]
		} else {
			rightArrXor = rightArrXor ^ nums[i]
		}

	}

	return []int{leftTotalXor ^ leftArrXor, rightTotalXor ^ rightArrXor}

}
