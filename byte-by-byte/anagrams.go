package main

/*
14. Anagrams
Question:​ Given two strings, write a function to determine whether they are anagrams.

isAnagram​(​""​, ​""​) = true
isAnagram​(​"A"​, ​"A"​) = true
isAnagram​(​"A"​, ​"B"​) = false
isAnagram​(​"ab"​, ​"ba"​) = true
isAnagram​(​"AB"​, ​"ab"​) = true

https://leetcode.com/problems/valid-anagram/

Time Complexity: O(n)
Space Complexity: O(n)

*/

import "strings"

/*
func main() {
	fmt.Println(isAnagram("rat", "car"))
}
*/

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	s = strings.ToLower(s)
	t = strings.ToLower(t)

	m := make(map[rune]int)

	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c] = m[c] + 1
		} else {
			m[c] = 1
		}
	}

	for _, c := range t {
		if v, ok := m[c]; ok {
			if v == 0 {
				return false
			} else {
				m[c] = m[c] - 1
			}
		} else {
			return false
		}
	}

	return true
}
