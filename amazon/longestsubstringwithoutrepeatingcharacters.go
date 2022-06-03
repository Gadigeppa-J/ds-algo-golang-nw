package main

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	l, r := 0, 0
	m := make(map[byte]struct{})
	max := 0

	for r < len(s) {

		_, ok := m[s[r]]

		if ok {
			for true {
				delete(m, s[l])
				l++
				if _, ok := m[s[r]]; !ok {
					break
				}
			}

		}

		m[s[r]] = struct{}{}
		max = utils.Max(max, r-l+1)
		r++

	}

	return max

}
