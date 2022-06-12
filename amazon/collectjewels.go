package main

import (
	"fmt"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {

	jewels := []JewelStone{
		{10, 60},
		{20, 100},
		{30, 120},
	}

	cap := 50

	fmt.Println(collectJewels(jewels, cap))

}

type JewelStone struct {
	w int
	v int
}

func collectJewels(jewels []JewelStone, cap int) int {
	res := 0
	collectJewelsHelp(jewels, cap, 0, 0, &res)
	return res
}

func collectJewelsHelp(jewels []JewelStone, cap int, pos int, curr int, res *int) {

	if cap == 0 {
		*res = utils.Max(*res, curr)
		return
	}

	if cap < 0 {
		return
	}

	for i := pos; i < len(jewels); i++ {
		curr = curr + jewels[i].v
		collectJewelsHelp(jewels, cap-jewels[i].w, i+1, curr, res)
		curr = curr - jewels[i].v
	}

}
