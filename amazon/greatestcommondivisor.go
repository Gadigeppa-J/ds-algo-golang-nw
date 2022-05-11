package main

import "fmt"

/*
TC: O(min(a, b))
SC: O(min(a, b))
*/

func main() {
	fmt.Println(gcd(6, 14))
}

func gcd(a, b int) int {

	if b == 0 {
		return a
	}

	// GCD formula
	// gcd (b, a%b)

	return gcd(b, a%b)
}
