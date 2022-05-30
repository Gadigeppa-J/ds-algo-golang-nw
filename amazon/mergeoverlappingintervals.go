package main

/*
https://leetcode.com/problems/merge-intervals/
TC: O(nlogn)
SC: O(n) if result is considered
*/

import (
	"sort"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	fmt.Println(mergeIntervals(intervals))
}
*/

func mergeIntervals(intervals [][]int) [][]int {

	if intervals == nil || len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for _, inter := range intervals {

		if result[len(result)-1][1] >= inter[0] {
			result[len(result)-1][1] = utils.Max(result[len(result)-1][1], inter[1])
		} else {
			result = append(result, inter)
		}
	}

	return result
}
