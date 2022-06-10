package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	res := []string{}
	generateParenthesisRecursive(n, "(", 1, 0, &res)
	return res
}

func generateParenthesisRecursive(n int, b string, o int, c int, res *[]string) {

	if len(b) == n*2 {
		*res = append(*res, b)
		return
	}

	if o < n {
		generateParenthesisRecursive(n, b+"(", o+1, c, res)
	}

	if c < o {
		generateParenthesisRecursive(n, b+")", o, c+1, res)
	}

}
