package main

/*
	no leetcode
	TC: O(nlogn)
	SC: O(1)
*/

import (
	"fmt"
	"sort"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	/*
		fmt.Println(minPlatforms([][]int{
			{200, 230},
			{210, 340},
			{300, 320},
			{320, 430},
			{350, 400},
			{500, 520},
		}))
	*/

	fmt.Println(minPlatforms([][]int{
		{120, 600},
		{50, 550},
		{550, 700},
		{200, 500},
		{700, 900},
		{850, 1000},
	}))
}

func minPlatforms(timings [][]int) int {

	if len(timings) <= 0 {
		return 0
	}

	arrival := make([]int, len(timings))
	departure := make([]int, len(timings))

	for i, v := range timings {
		arrival[i] = v[0]
		departure[i] = v[1]
	}

	sort.Ints(arrival)
	sort.Ints(departure)

	minPlat := 0
	aptr := 1
	dptr := 0
	platCount := 1

	for aptr < len(timings) && dptr < len(timings) {

		if arrival[aptr] <= departure[dptr] {
			platCount++
			aptr++
		} else {
			platCount--
			dptr++
		}
		minPlat = utils.Max(minPlat, platCount)
	}

	return minPlat
}
