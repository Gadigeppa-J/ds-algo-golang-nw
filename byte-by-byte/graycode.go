package main

/*
35. Gray Code
Question:​ Given two integers, write a function to determine whether or not their binary representations differ by a single bit.

eg.
gray​(​0​, ​1​) = true
gray​(​1​, ​2​) = false
*/

func gray(i, j int) bool {
	k := i ^ j
	for k > 0 {
		if k%2 == 1 {
			if k>>1 == 0 {
				return true
			} else {
				return false
			}
		}
		k = k >> 1
	}
	return false
}
