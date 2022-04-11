package main

/*
8. Merge K Arrays
Question:​ Given k sorted arrays, merge them into a single sorted array.

 merge({{​1​, ​4​, ​7​},{​2​, ​5​, ​8​},{​3​, ​6​, ​9​}}) = {​1​, ​2​, ​3​, ​4​, ​5​, ​6​, ​7​, ​8​, ​9​}

 https://leetcode.com/problems/merge-k-sorted-lists/

 Time & Space complexity: https://www.youtube.com/watch?v=ptYUCjfNhJY

*/

import (
	"fmt"

	ge "github.com/Gadigeppa-J/ds-algo-golang-nw/general"
)

/*
func main() {

	matrix := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	fmt.Println(mergeKArrays(matrix))

}
*/

type arrItem struct {
	val      int
	colIndex int
	rowIndex int
}

func mergeKArrays(matrix [][]int) []int {

	r := []int{}

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return r
	}

	k := len(matrix)
	pq := ge.NewPriorityQueue()

	// add first k items to queue
	for i := 0; i < k; i++ {
		item := &ge.PriorityQueueItem{
			Val: matrix[i][0],
			Meta: arrItem{
				val:      matrix[i][0],
				colIndex: 0,
				rowIndex: i,
			},
		}
		pq.Insert(item)
	}

	for !pq.IsEmpty() {
		item := pq.Remove()
		fmt.Println(item.Val)

		r = append(r, item.Val)

		ar := item.Meta.(arrItem)

		colIndex := ar.colIndex
		rowIndex := ar.rowIndex

		if len(matrix[rowIndex]) > (colIndex + 1) {

			// add new item to queue
			item := &ge.PriorityQueueItem{
				Val: matrix[rowIndex][colIndex+1],
				Meta: arrItem{
					val:      matrix[rowIndex][colIndex+1],
					colIndex: colIndex + 1,
					rowIndex: rowIndex,
				},
			}
			pq.Insert(item)

		}

	}

	return r
}
