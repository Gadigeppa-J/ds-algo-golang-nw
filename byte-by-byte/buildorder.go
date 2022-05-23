package main

/*
Given a list of packages that need to be built and the dependencies for each
package, determine a valid order in which to build the packages.

0:
1: 0
2: 0
3: 1, 2
4: 3
output: 0, 1, 2, 3, 4

TC: O(n)
SC: O(n)
*/

/* func main() {

	build := [][]int{{}, {0}, {0}, {1, 2}, {3}}
	result := []int{}
	isvisited := make([]bool, len(build))
	temp := make([]bool, len(build))

	for i := range result {
		result[i] = -1
	}

	for i := range build {
		buildOrder(i, build, isvisited, temp, &result)
	}

	fmt.Println(result)

} */

func buildOrder(inx int, build [][]int, isvisited []bool, temp []bool, result *[]int) {

	if isvisited[inx] == true {
		return
	}

	if temp[inx] {
		panic("Cycle detected!")
	}

	temp[inx] = true

	for i := 0; i < len(build[inx]); i++ {
		if !isvisited[build[inx][i]] {
			buildOrder(build[inx][i], build, isvisited, temp, result)
		}
	}

	isvisited[inx] = true
	temp[inx] = false

	*result = append(*result, inx)

}
