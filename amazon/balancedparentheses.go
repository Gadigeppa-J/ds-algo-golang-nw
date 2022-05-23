package main

import (
	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

/* func main() {
	fmt.Println(isValid("(("))
} */

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	stack := utils.NewStack()

	for _, c := range s {

		if string(c) == "(" || string(c) == "[" || string(c) == "{" {
			stack.Push(int(c))
		} else if string(c) == ")" {
			if stack.IsEmpty() {
				return false
			}

			c := string(rune(stack.Pop()))

			if c != "(" {
				return false
			}

		} else if string(c) == "]" {
			if stack.IsEmpty() {
				return false
			}

			c := string(rune(stack.Pop()))

			if c != "[" {
				return false
			}

		} else if string(c) == "}" {
			if stack.IsEmpty() {
				return false
			}

			c := string(rune(stack.Pop()))

			if c != "{" {
				return false
			}
		}

	}

	return stack.IsEmpty()
}
