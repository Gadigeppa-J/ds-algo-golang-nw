package main

import (
	"fmt"
	"sort"
)

func main() {

	meetings := [][]int{
		{3, 29},
		{50, 93},
		{88, 92},
		{54, 67},
		{50, 87},
	}

	fmt.Println(maxMeetings(meetings))
}

func maxMeetings(meetings [][]int) int {

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][1] < meetings[j][1]
	})

	max := 1
	lastEnd := meetings[0][1]

	for i := 1; i < len(meetings); i++ {
		if lastEnd <= meetings[i][0] {
			max++
			lastEnd = meetings[i][0]
		}
	}

	return max
}
