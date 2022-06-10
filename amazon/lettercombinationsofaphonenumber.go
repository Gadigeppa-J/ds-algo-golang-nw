package main

/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number/
TC: O(n4^n)
SC:
*/

import "fmt"

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {

	res := []string{}

	if digits == "" {
		return res
	}

	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	letterCombinationsRecursive(digits, 0, "", m, &res)

	return res
}

func letterCombinationsRecursive(digits string, index int, curr string, m map[string]string, res *[]string) {

	if len(curr) == len(digits) {
		*res = append(*res, curr)
		return
	}

	v := m[string(digits[index])]

	for i := 0; i < len(v); i++ {
		letterCombinationsRecursive(digits, index+1, curr+string(v[i]), m, res)
	}

}
