package main

/*
https://leetcode.com/problems/set-matrix-zeroes/
*/
import "fmt"

func main() {

	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {

	rw := make([]int, len(matrix))
	cl := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rw[i] = -1
				cl[j] = -1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rw[i] == -1 || cl[j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
}
