package main

import "fmt"

/*
https://leetcode.com/problems/fibonacci-number/
TC: O(n)
SC: O(1)
*/

func main() {
	fmt.Println(fib(8))
}

func fib(n int) int {

	if n <= 0 {
		return 0
	}

	i := 1
	j := 1
	n = n - 2

	for k := 0; k < n; k++ {
		tmp := i + j
		i = j
		j = tmp
	}

	return j

}
