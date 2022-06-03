package main

/*
no leetcode
TC: O(n)
SC: O(n)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	fmt.Println(lengthOfLongestSubstringKUnique("mississippi", 4))
}

func lengthOfLongestSubstringKUnique(s string, k int) int {

	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	l, r := 0, 0
	m := make(map[byte]int)

	max := 0
	for r < len(s) {

		_, ok := m[s[r]]
		if ok {
			m[s[r]] = m[s[r]] + 1
		} else {
			m[s[r]] = 1
		}

		if len(m) > k {

			for len(m) > k {
				lv := m[s[l]]

				if lv == 1 {
					delete(m, s[l])
				} else {
					m[s[l]] = m[s[l]] - 1
				}

				l++
			}

		}

		max = utils.Max(max, r-l+1)
		r++
	}

	if len(m) < 3 {
		return -1
	}

	return max

}
