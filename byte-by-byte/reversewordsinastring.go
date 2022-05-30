package main

import "github.com/Gadigeppa-J/ds-algo-golang-nw/utils"

func reverseWords(s string) string {

	sk := utils.NewGenericStack()
	word := ""

	s = s + " "

	for _, c := range s {
		if string(c) != " " {
			word = word + string(c)
		} else {

			if word != "" {
				sk.Push(word)
				word = ""
			}

		}

	}

	reversed := ""

	for !sk.IsEmpty() {
		reversed = reversed + sk.Pop().(string) + " "
	}

	return reversed[:len(reversed)-1]

}
