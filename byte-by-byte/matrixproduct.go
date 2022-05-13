package main

/*
TC: 2^(r+c)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	//fmt.Println(maxProductPath([][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}))

	fmt.Println(maxProductPath([][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}))
}

func maxProductPath(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	return maxProductPathHelpRecurive(grid, 0, 0, 1)

}

/*
func maxProductPathHelp(grid [][]int, i, j, product int, maxi *int) {

	if i >= len(grid) || j >= len(grid[0]) {
		return
	}
	product = product * grid[i][j]

	if i == len(grid)-1 && j == len(grid[0])-1 {
		*maxi = utils.Max(*maxi, product)
	}

	maxProductPathHelp(grid, i, j+1, product, maxi)
	maxProductPathHelp(grid, i+1, j, product, maxi)

}
*/

func maxProductPathHelpRecurive(grid [][]int, i, j, product int) int {

	if i >= len(grid)-1 && j >= len(grid[0])-1 {
		return product * grid[i][j]
	}

	downProduct := -1
	rightProduct := -1

	product = product * grid[i][j]

	if j < len(grid[0])-1 {
		rightProduct = maxProductPathHelpRecurive(grid, i, j+1, product)
	}

	if i < len(grid)-1 {
		downProduct = maxProductPathHelpRecurive(grid, i+1, j, product)
	}

	return utils.Max(downProduct, rightProduct)
}
