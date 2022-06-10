package main

/*
no leetcode
*/

import "fmt"

func main() {

	s := "abcd"
	fmt.Println(getAllPermutations(s))

}

func getAllPermutations(s string) []string {

	res := []string{}

	if s == "" {
		return res
	}

	used := make([]bool, len(s))
	getAllPermutationsRecursive(s, "", used, &res)
	return res
}

func getAllPermutationsRecursive(s string, curr string, used []bool, res *[]string) {

	if len(curr) == len(s) {
		*res = append(*res, curr)
		return
	}

	for i := 0; i < len(s); i++ {

		if used[i] {
			continue
		}
		used[i] = true
		getAllPermutationsRecursive(s, curr+string(s[i]), used, res)
		used[i] = false
	}

}
