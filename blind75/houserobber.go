package main

/*
func main() {
	nums := []int{2, 7, 9, 3, 1}

	fmt.Println(rob(nums, 0, make(map[int]int)))
}
*/

func rob(nums []int, pos int, mem map[int]int) int {

	if pos > len(nums)-1 {
		return 0
	}

	if v, ok := mem[pos]; ok {
		return v
	}

	pick := nums[pos] + rob(nums, pos+2, mem)
	notPick := 0 + rob(nums, pos+1, mem)

	mem[pos] = max(pick, notPick)

	return mem[pos]
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}
