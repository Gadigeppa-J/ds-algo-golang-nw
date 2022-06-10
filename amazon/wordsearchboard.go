package main

/*
https://leetcode.com/problems/word-search/
TC: O(m*n*dfs()) dsf()-> 4^len(word)
SC: O(len(word))
*/

import (
	"fmt"
)

func main() {

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCESEEEFS"

	fmt.Println(wordSearch(board, word))

}

type coord struct {
	i int
	j int
}

func wordSearch(board [][]byte, word string) bool {

	if len(board) == 0 || len(word) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if wordSearchRecursive(board, make(map[coord]struct{}), word, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func wordSearchRecursive(board [][]byte, path map[coord]struct{}, word string, pos int, i int, j int) bool {

	if pos >= len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		board[i][j] != word[pos] {
		return false
	}

	if _, ok := path[coord{i, j}]; ok {
		return false
	}

	path[coord{i, j}] = struct{}{}
	res := wordSearchRecursive(board, path, word, pos+1, i-1, j) ||
		wordSearchRecursive(board, path, word, pos+1, i+1, j) ||
		wordSearchRecursive(board, path, word, pos+1, i, j-1) ||
		wordSearchRecursive(board, path, word, pos+1, i, j+1)
	delete(path, coord{i, j})

	return res
}
