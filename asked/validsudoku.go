package main

/*
https://leetcode.com/problems/valid-sudoku/submissions/
TC: O(n)
SC: O(n)
*/

import "fmt"

func main() {

	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {

	m := make(map[string]struct{})

	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[0]); j++ {

			if board[i][j] == '.' {
				continue
			}

			// row
			r := fmt.Sprintf("%s%d%v", "r", i, board[i][j])
			if _, ok := m[r]; ok {
				return false
			}
			m[r] = struct{}{}

			// col
			c := fmt.Sprintf("%s%d%v", "c", j, board[i][j])
			if _, ok := m[c]; ok {
				return false
			}
			m[c] = struct{}{}

			// box
			b := fmt.Sprintf("%s%d-%d%v", "b", (i / 3), (j / 3), board[i][j])
			if _, ok := m[b]; ok {
				return false
			}
			m[b] = struct{}{}

		}
	}

	return true

}
