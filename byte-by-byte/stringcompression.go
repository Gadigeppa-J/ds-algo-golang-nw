package main

/*
https://leetcode.com/problems/string-compression/

TC: O(n)
TC: O(1)
*/

import (
	"strconv"
)

/* func main() {

	chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(chars))

} */

func compress(chars []byte) int {

	i, j := 0, 0
	index := 0

	for j < len(chars) {

		i = j
		c := chars[i]

		for j < len(chars) && chars[i] == chars[j] {
			j++
		}

		chars[index] = c
		index++

		d := j - i
		if d > 1 {

			ns := strconv.FormatInt(int64(d), 10)

			for _, dc := range ns {
				chars[index] = byte(dc)
				index++
			}
		}

	}

	return index

}
