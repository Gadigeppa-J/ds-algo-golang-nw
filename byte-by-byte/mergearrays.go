package main

/*
10. Merge Arrays
Question:​ Given 2 sorted arrays, A and B, where A is long enough to hold the contents of A and B, write a function to copy the contents of B into A without using any buffer or additional memory.

A = {​1​,​3​,​5​,​0​,​0​,​0​} B = {​2​,​4​,​6​} mergeArrays(A, B) A = {​1​,​2​,​3​,​4​,​5​,​6​}

https://leetcode.com/problems/merge-sorted-array/

Time Complexity: O(n)
Space Complexity: O(1)

*/

/*
func main() {

	a := []int{1, 3, 5, 0, 0, 0}
	b := []int{2, 4, 6}

	mergeArrays(a, 3, b, 3)
	fmt.Println(a)
}
*/

func mergeArrays(a []int, m int, b []int, n int) {

	aptr := m - 1
	aendptr := len(a) - 1
	bptr := n - 1

	for true {

		if aptr < 0 && bptr < 0 {
			break
		}

		if aptr < 0 {
			a[aendptr] = b[bptr]
			bptr--
		} else if bptr < 0 {
			a[aendptr] = a[aptr]
			aptr--
		} else {
			if a[aptr] > b[bptr] {
				a[aendptr] = a[aptr]
				aptr--
			} else {
				a[aendptr] = b[bptr]
				bptr--
			}
		}

		aendptr--
	}

}
