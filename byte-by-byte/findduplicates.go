/*
4. Find Duplicates
Question: Given an array of integers where each value 1 <= x <= len(array), write a
function that finds all the duplicates in the array.

findDuplicates([1, 2, 3]) = []
findDuplicates([1, 2, 2]) = [2]
findDuplicates([3, 3, 3]) = [3]
findDuplicates([2, 1, 2, 1]) = [1, 2]

LeetCode: https://leetcode.com/problems/find-all-duplicates-in-an-array/
*/

package main

func findDuplicates(nums []int) []int {

	m := make(map[int]int)
	o := []int{}

	for _, n := range nums {
		_, ok := m[n]
		if ok {
			m[n] += 1
		} else {
			m[n] = 1
		}
	}

	for k, v := range m {
		if v >= 2 {
			o = append(o, k)
		}
	}

	return o

}
