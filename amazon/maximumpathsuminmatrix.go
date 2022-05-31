package main

/*
no leetcode
DP
TC: O(m*n)
SC: O(1)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	matrix := [][]int{
		{10, 10, 2, 0, 20, 4},
		{1, 0, 0, 30, 2, 5},
		{0, 10, 4, 0, 2, 0},
		{1, 0, 2, 20, 0, 4},
	}

	/*
		matrix := [][]int{
			{12, 22, 5, 0, 20, 18},
			{0, 2, 5, 25, 18, 17},
			{12, 10, 5, 4, 2, 1},
			{3, 8, 2, 20, 10, 8},
		}
	*/

	fmt.Println(maxPathInMatrix(matrix))
}

func maxPathInMatrix(matrix [][]int) int {

	r := len(matrix)
	c := len(matrix[0])
	max := 0

	for i := 1; i < r; i++ {
		for j := 0; j < c; j++ {

			if c == 1 {
				matrix[i][j] = matrix[i][j] + matrix[i-1][j]
			} else if j == 0 {
				matrix[i][j] = matrix[i][j] + utils.Max(matrix[i-1][j], matrix[i-1][j+1])
			} else if j == c-1 {
				matrix[i][j] = matrix[i][j] + utils.Max(matrix[i-1][j], matrix[i-1][j-1])
			} else {
				matrix[i][j] = matrix[i][j] + utils.Max(utils.Max(matrix[i-1][j], matrix[i-1][j-1]), matrix[i-1][j+1])
			}

		}
	}

	for i := 0; i < c; i++ {
		max = utils.Max(max, matrix[r-1][i])
	}

	return max
}
