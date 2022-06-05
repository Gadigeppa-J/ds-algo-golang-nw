package main

/*
no leetcode
TC: O(m*n)
SC: O(m*n)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	/*
		matrix := [][]int{
			{2, 1, 0},
			{1, 1, 0},
			{0, 1, 1},
		}
	*/

	matrix := [][]int{
		{2, 1, 0, 2, 1},
		{0, 0, 1, 2, 1},
		{1, 0, 0, 2, 1},
	}

	fmt.Println(getDaysToRot(matrix))

}

type RotNode struct {
	timefarme int
	coord     []int
}

func getDaysToRot(matrix [][]int) int {

	m := len(matrix)
	n := len(matrix[0])

	days := 0
	q := utils.NewGenericQueue()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 2 {
				q.Enqueue(&RotNode{
					timefarme: 0,
					coord:     []int{i, j},
				})
			}
		}
	}

	for !q.IsEmpty() {
		node := q.Dequeue().(*RotNode)
		i := node.coord[0]
		j := node.coord[1]
		t := node.timefarme

		// check adjacent

		// up
		if i > 0 && matrix[i-1][j] == 1 {
			matrix[i-1][j] = 2
			q.Enqueue(&RotNode{
				timefarme: t + 1,
				coord:     []int{i - 1, j},
			})
		}

		// down
		if i < m-1 && matrix[i+1][j] == 1 {
			matrix[i+1][j] = 2
			q.Enqueue(&RotNode{
				timefarme: t + 1,
				coord:     []int{i + 1, j},
			})
		}

		// right
		if j < n-1 && matrix[i][j+1] == 1 {
			matrix[i][j+1] = 2
			q.Enqueue(&RotNode{
				timefarme: t + 1,
				coord:     []int{i, j + 1},
			})
		}

		// left
		if j > 0 && matrix[i][j-1] == 1 {
			matrix[i][j-1] = 2
			q.Enqueue(&RotNode{
				timefarme: t + 1,
				coord:     []int{i, j - 1},
			})
		}

		days = t
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				return -1
			}
		}
	}

	return days
}
