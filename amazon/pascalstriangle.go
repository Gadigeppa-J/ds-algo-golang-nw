package main

/*
https://leetcode.com/problems/pascals-triangle/
TC: O(n^2)
SC: O(n)

Your triangle has 1 + 2 + ... + n elements. This is arithmetic progression that sums to n*(n+1)/2, which is in O(n^2)

*/

/*
func main() {

	numRows := 0
	fmt.Println(generate(numRows))

}
*/

func generate(numRows int) [][]int {

	if numRows <= 0 {
		return [][]int{}
	}

	triangle := [][]int{{1}}

	for i := 1; i <= numRows-1; i++ {

		row := []int{}
		prev := append(append([]int{0}, triangle[i-1]...), 0)

		j := 0
		k := 1

		for k < len(prev) {
			row = append(row, prev[j]+prev[k])
			j++
			k++
		}
		triangle = append(triangle, row)
	}

	return triangle
}
