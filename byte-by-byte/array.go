package main

/*
5. Consecutive Array
Question:​ Given an unsorted array, find the length of the longest sequence of consecutive numbers in the array.
*/
func longestConsecutive(arr []int) int {
	set := make(map[int]bool)

	for _, val := range arr {
		set[val] = true
	}

	max, len := 0, 0

	for key, _ := range set {

		curr := key

		if _, ok := set[curr-1]; ok {
			continue
		}

		for ok := set[curr]; ok; ok = set[curr] {
			len++
			curr++
		}

		if len > max {
			max = len
		}

		len = 0

	}

	return max
}
