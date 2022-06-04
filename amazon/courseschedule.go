package main

/*
https://leetcode.com/problems/course-schedule/submissions/
TC: O(numCourses + len(prerequisites))
SC: O(numCourses)
*/

import "fmt"

func main() {

	numCource := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCource, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	almap := make(map[int][]int)

	for i := 0; i < len(prerequisites); i++ {
		crs := prerequisites[i][0]
		pre := almap[crs]
		pre = append(pre, prerequisites[i][1])
		almap[crs] = pre
	}

	visited := make(map[int]struct{})
	for i := 0; i < numCourses; i++ {
		if !canFinishDFS(almap, i, visited) {
			return false
		}
	}

	return true
}

func canFinishDFS(almap map[int][]int, crs int, visited map[int]struct{}) bool {

	if len(almap[crs]) == 0 {
		return true
	}

	if _, ok := visited[crs]; ok {
		return false
	}

	visited[crs] = struct{}{}

	pre := almap[crs]

	for i := 0; i < len(pre); i++ {
		if !canFinishDFS(almap, pre[i], visited) {
			return false
		}
	}

	almap[crs] = []int{}

	return true
}
