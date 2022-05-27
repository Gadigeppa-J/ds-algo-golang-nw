package main

/*
https://leetcode.com/problems/rotate-image/
TC: O(n^2)
SC: 1
*/

import "fmt"

func main() {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotateMatrix(matrix)

	fmt.Println(matrix)

}

func rotateMatrix(matrix [][]int) {

	transposeMatrix(matrix)

	n := len(matrix[0])

	for _, m := range matrix {

		i := 0
		j := n - 1

		for i < j {
			tmp := m[i]
			m[i] = m[j]
			m[j] = tmp
			i++
			j--
		}

	}

}

func transposeMatrix(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := i; j < n; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

}
