package main

/*
https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
TC: O(n)
SC:O(1)
*/

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	s := "tryhard"
	k := 4
	fmt.Println(maxVowels(s, k))
}

func maxVowels(s string, k int) int {

	if len(s) == 0 {
		return 0
	}

	max := 0
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	swcount := 0

	for i := 0; i < k; i++ {
		if _, ok := vowels[s[i]]; ok {
			swcount++
		}
	}

	max = swcount

	for i := k; i < len(s); i++ {

		if _, ok := vowels[s[i-k]]; ok { // exiting
			swcount--
		}

		if _, ok := vowels[s[i]]; ok { // entering
			swcount++
		}

		max = utils.Max(max, swcount)
	}

	return max
}
