package main

/*
11. Zero Sum Subarray

Question:  Given an array, write a function to find any subarray that sums to zero, if one exists.
zeroSum({ 1 ,  2 ,  -5 ,  1 ,  2 ,  -1 }) = [ 2 ,  -5 ,  1 ,  2 ]
zeroSum({ 4, 2, -3, 1, 6}) = [ 2, -3, 1]
zeroSum({ 4, 2, 0, 1, 6}) = [0]
zeroSum({ -3, 2, 3, 1, 6}) = []

Time Complexity: O(n)
Space Complexity: O(n)
*/

func zeroSum(a []int) []int {

	if len(a) == 0 {
		return a
	}

	m := make(map[int]int)
	sum := 0

	for i, v := range a {

		sum += v
		idx, ok := m[sum]

		if ok {
			return a[idx+1 : i+1]
		}
		m[sum] = i
	}

	return []int{}
}
