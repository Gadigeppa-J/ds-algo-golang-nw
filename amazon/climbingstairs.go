package main

/*
https://leetcode.com/problems/climbing-stairs/

climbStairsWithOutMemoization
TC: O(2^n)
SC:

climbStairsWithMemotization
TC: O(n)
SC:


*/

func climbStairs(n int) int {

	return climbStairsDP(n)
	//return climbStairsWithMemotization(n, 0, make(map[int]int))
}

func climbStairsDP(n int) int {

	if n <= 1 {
		return n
	}

	p := 1
	q := 1
	ans := p + q

	for i := 0; i < n-2; i++ {
		q = p
		p = ans
		ans = p + q
	}

	return ans
}

func climbStairsWithMemoization(n int, s int, m map[int]int) int {

	if v, ok := m[s]; ok {
		return v
	}

	if n == s {
		return 1
	}

	if n < s {
		return 0
	}

	n1 := climbStairsWithMemoization(n, s+1, m)
	n2 := climbStairsWithMemoization(n, s+2, m)

	m[s] = n1 + n2

	return n1 + n2
}
