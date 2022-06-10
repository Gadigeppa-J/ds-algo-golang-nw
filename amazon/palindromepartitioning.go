package main

/*
https://leetcode.com/problems/palindrome-partitioning/
*/

import "fmt"

func main() {
	fmt.Println(partition("cbbbcc"))
}

func partition(s string) [][]string {

	if len(s) == 0 {
		return [][]string{}
	}

	res := [][]string{}
	partitionRecursive(s, 0, []string{}, &res)
	return res
}

func partitionRecursive(s string, pos int, path []string, res *[][]string) {

	if pos >= len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := pos; i < len(s); i++ {
		wrd := s[pos : i+1]
		if isPalindrom(wrd) {
			path = append(path, wrd)
			partitionRecursive(s, i+1, path, res)
			path = path[:len(path)-1]
		}
	}

}

func isPalindrom(nums string) bool {

	if len(nums) == 0 {
		return false
	}

	l := 0
	h := len(nums) - 1

	for l <= h {
		if nums[l] != nums[h] {
			return false
		}
		l++
		h--
	}

	return true
}
