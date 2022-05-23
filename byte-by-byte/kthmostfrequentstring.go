package main

/*
Given a list of strings, write a function to get the kth most frequently
occurring string.
kthMostFrequent({"a","b","c","a","b","a"}, 0) = "a"
kthMostFrequent({"a","b","c","a","b","a"}, 1) = "b"
kthMostFrequent({"a","b","c","a","b","a"}, 2) = "c"
kthMostFrequent({"a","b","c","a","b","a"}, 3) = null

TC: O(nlogn)
SC: O(n)
*/

import (
	"sort"
)

/* func main() {
	fmt.Println(kthMostFrequestString([]string{"a", "b", "c", "a", "b", "a"}, 2))
} */

type word struct {
	w    string
	freq int
}

func kthMostFrequestString(input []string, k int) string {

	m := make(map[string]int)

	for _, i := range input {
		if _, ok := m[i]; !ok {
			m[i] = 0
		}
		m[i] = m[i] + 1
	}

	list := []word{}
	for k, v := range m {
		list = append(list, word{
			w:    k,
			freq: v,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].freq > list[j].freq
	})

	if k < len(list) {
		return list[k].w
	}

	return ""
}
