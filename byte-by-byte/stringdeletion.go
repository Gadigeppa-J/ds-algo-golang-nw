package main

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/*
Given a string and a dictionary HashSet, write a function to determine the
minimum number of characters to delete to make a word.

dictionary: ["a", "aa", "aaa"]
query: "abc"
output: 2

*/

func main() {
	fmt.Println(deleteString("abc", []string{"a", "aa", "aaa"}))
}

func deleteString(query string, dict []string) int {

	dictmap := make(map[string]struct{})

	for _, d := range dict {
		dictmap[d] = struct{}{}
	}

	q := utils.NewGenericQueue()
	m := make(map[string]struct{})

	q.Enqueue(query)
	m[query] = struct{}{}

	for !q.IsEmpty() {
		item := q.Dequeue().(string)
		delete(m, query)

		if _, ok := dictmap[item]; ok {
			return len(query) - len(item)
		}

		for i := 0; i < len(item); i++ {
			ss := item[:i] + item[i+1:]
			q.Enqueue(ss)
			m[ss] = struct{}{}
		}
	}

	return -1
}
