package main

/*
	https://leetcode.com/problems/longest-palindromic-substring/
	TC: O(n^2)
	SC: O(1)
*/

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {

	resStr := ""
	resLen := 0

	sl := len(s)

	for i, _ := range s {

		// even palindrome
		r, l := i, i

		for l >= 0 && r < sl && s[l] == s[r] {
			if r-l+1 > resLen {
				resLen = r - l + 1
				resStr = s[l : r+1]
			}
			l--
			r++
		}

		// odd palindrome
		r, l = i+1, i

		for l >= 0 && r < sl && s[l] == s[r] {
			if r-l+1 > resLen {
				resLen = r - l + 1
				resStr = s[l : r+1]
			}
			l--
			r++
		}

	}

	return resStr
}
