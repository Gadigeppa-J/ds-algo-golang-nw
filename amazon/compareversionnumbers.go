package main

/*
https://leetcode.com/problems/compare-version-numbers/
*/

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	fmt.Println(compareVersion("1.0", "1.0.1"))
}

func compareVersion(version1 string, version2 string) int {

	v1parts := strings.Split(version1, ".")
	v2parts := strings.Split(version2, ".")
	max := utils.Max(len(v1parts), len(v2parts))

	for i := 0; i < max; i++ {

		if i >= len(v1parts) {
			v2conv, _ := strconv.ParseInt(v2parts[i], 10, 32)
			if v2conv > 0 {
				return -1
			}
		} else if i >= len(v2parts) {
			v1conv, _ := strconv.ParseInt(v1parts[i], 10, 32)
			if v1conv > 0 {
				return 1
			}
		} else {
			v1conv, _ := strconv.ParseInt(v1parts[i], 10, 32)
			v2conv, _ := strconv.ParseInt(v2parts[i], 10, 32)

			if v1conv > v2conv {
				return 1
			} else if v1conv < v2conv {
				return -1
			}
		}

	}

	return 0

}
