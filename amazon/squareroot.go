package main

/*
https://leetcode.com/problems/sqrtx/
TC = O(log sqrt(x) )
SC = O(1)
*/

/*
func main() {

	fmt.Println(mySqrt(9))
}
*/

func mySqrt(x int) int {

	l := 0
	h := x
	ans := 0

	for l <= h {

		mid := (l + h) / 2

		if (mid * mid) <= x {
			ans = mid
			l = mid + 1
		} else {
			h = mid - 1
		}
	}

	return ans
}
