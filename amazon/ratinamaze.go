package main

/*
no leetcode
Time Complexity: O(2^(n^2))
Space Complexity: O(n^2)
*/

import "fmt"

func main() {

	matrix := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 1},
	}

	fmt.Println(solveRatMaze(matrix, 0, 0))

}

func solveRatMaze(matrix [][]int, i int, j int) bool {

	if i == len(matrix)-1 && j == len(matrix[0])-1 && matrix[i][j] == 1 {
		return true
	}

	if matrix[i][j] == 1 {

		if i < len(matrix)-1 {
			if solveRatMaze(matrix, i+1, j) {
				return true
			}
		}

		if j < len(matrix[0])-1 {
			return solveRatMaze(matrix, i, j+1)
		}
	}

	return false

}
