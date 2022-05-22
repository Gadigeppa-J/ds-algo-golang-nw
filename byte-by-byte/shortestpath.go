package main

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	gn3 := &GraphNode{
		Val: 3,
	}

	gn4 := &GraphNode{
		Val: 4,
	}

	gn5 := &GraphNode{
		Val: 5,
	}

	gn2 := &GraphNode{
		Val: 2,
	}

	gn1 := &GraphNode{
		Val: 1,
	}

	gn1.Next = []*GraphNode{gn2, gn3}
	gn2.Next = []*GraphNode{gn5}
	gn5.Next = []*GraphNode{gn4}
	gn4.Next = []*GraphNode{gn3, gn1}

	sp := shortestpath(gn2, gn3)

	for _, n := range sp {
		fmt.Println(n.Val)
	}

}

type GraphNode struct {
	Val  int
	Next []*GraphNode
}

func shortestpath(a *GraphNode, b *GraphNode) []*GraphNode {

	if a == nil || b == nil {
		return nil
	}

	q := utils.NewGenericQueue()
	q.Enqueue(a)
	visited := make(map[*GraphNode]struct{})
	parents := make(map[*GraphNode]*GraphNode)

	var curr *GraphNode
	parents[curr] = nil

	for !q.IsEmpty() {
		curr = q.Dequeue().(*GraphNode)

		if curr == b {
			break
		}

		if _, ok := visited[curr]; ok {
			panic("Cycle Detected!")
		}

		visited[curr] = struct{}{}

		for _, nxt := range curr.Next {
			q.Enqueue(nxt)
			parents[nxt] = curr
		}

	}

	res := []*GraphNode{}

	for curr != nil {
		res = append(res, curr)
		curr = parents[curr]
	}

	return res
}
