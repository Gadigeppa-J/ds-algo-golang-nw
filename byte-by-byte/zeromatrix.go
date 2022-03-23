package main

/*
6. Zero Matrix
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place.

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Time Complexity: O(m + n)
Space Complexity: 0(1)

LeetCode: https://leetcode.com/problems/set-matrix-zeroes/
*/

func setZeroes(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])

	cl := make([]bool, m)
	rw := make([]bool, n)

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {

			if matrix[i][j] == 0 {
				cl[i] = true
				rw[j] = true
			}
		}
	}

	for i, v := range rw {
		if v == true {
			for j := 0; j < m; j++ {
				matrix[j][i] = 0
			}
		}
	}

	for i, v := range cl {
		if v == true {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

}
