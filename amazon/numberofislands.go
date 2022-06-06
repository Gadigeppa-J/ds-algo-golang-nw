package main

import "fmt"

func main() {

	/*
		grid := [][]byte{
			{1, 1, 1, 1, 0},
			{1, 1, 0, 1, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
	*/

	/*
		grid := [][]byte{
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
		}
	*/

	//	[["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]

	grid := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	fmt.Println(numIslands(grid))

}

func numIslands(grid [][]byte) int {

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res := numIslandsHelp(grid, i, j)

			if res == 1 {
				count++
			}
		}
	}

	return count

}

func numIslandsHelp(grid [][]byte, i, j int) int {

	// check boundary
	if i < 0 || i > len(grid)-1 {
		return -1
	}

	// check boundary
	if j < 0 || j > len(grid[0])-1 {
		return -1
	}

	if string(grid[i][j]) == "2" || string(grid[i][j]) == "0" {
		return -1
	}
	grid[i][j] = []byte("2")[0]

	// top
	numIslandsHelp(grid, i-1, j)

	// down
	numIslandsHelp(grid, i+1, j)

	// left
	numIslandsHelp(grid, i, j-1)

	// right
	numIslandsHelp(grid, i, j+1)

	return 1
}
