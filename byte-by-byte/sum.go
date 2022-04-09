package main

/*
19. Sum
Question:​ Given two integers, write a function to sum the numbers without using any arithmetic operators.

https://leetcode.com/problems/sum-of-two-integers/
*/

func getSum(a int, b int) int {

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	sum := 0

	for b != 0 {
		sum = a ^ b
		carry := a & b
		b = carry << 1
		a = sum
	}

	return sum
}
