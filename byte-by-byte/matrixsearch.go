package main

/*
9. Matrix Search
Question:​ Given an n x m array where all rows and columns are in sorted order, write a function to determine whether the array contains an element x.

https://leetcode.com/problems/search-a-2d-matrix/

Time Complexity: O(m+n)
Space Complexity: O(1)

*/

/*
func main() {

	fmt.Println(matrixSearch([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, 13))

}
*/

func matrixSearch(matrix [][]int, n int) bool {

	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {

		if matrix[row][col] == n {
			return true
		} else if matrix[row][col] < n {
			row++
		} else if matrix[row][col] > n {
			col--
		}
	}

	return false
}
