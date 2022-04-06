package main

/*
37. Number of Ones in a Binary Number
Question:​ Given an integer, write a function to compute the number of ones in the binary representation of the number.

eg.
ones​(​0​) = ​0
ones​(​1​) = ​1
ones​(​2​) = ​1
ones​(​3​) = ​2
ones​(​7​) = ​3

https://leetcode.com/problems/number-of-1-bits/
*/

func ones(num uint32) int {

	count := 0

	for num != 0 {
		if num%2 == 1 {
			count++
		}
		num = num >> 1
	}

	return count

}
