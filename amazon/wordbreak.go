package main

import "fmt"

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak(s, wordDict))

}

func wordBreak(s string, wordDict []string) bool {

	if s == "" || len(wordDict) == 0 {
		return false
	}

	m := make(map[string]struct{})

	for _, wrd := range wordDict {
		m[wrd] = struct{}{}
	}

	return wordBreakRecursive(s, m, 0, make(map[int]bool))
}

func wordBreakRecursive(s string, wordDict map[string]struct{}, pos int, mem map[int]bool) bool {

	if pos >= len(s) {
		return true
	}

	if _, ok := mem[pos]; ok {
		fmt.Println("entered ", pos)
		return mem[pos]
	}

	for i := pos; i < len(s); i++ {

		wrd := s[pos : i+1]
		_, ok := wordDict[wrd]

		if ok && wordBreakRecursive(s, wordDict, i+1, mem) {
			mem[pos] = true
			return true
		}

	}

	mem[pos] = false
	return false

}
