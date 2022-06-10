package main

/*
https://leetcode.com/problems/restore-ip-addresses/
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

func restoreIpAddresses(s string) []string {

	if len(s) < 4 {
		return []string{}
	}

	intermRes := [][]string{}
	restoreIpAddressesRecursive(s, 0, []string{}, &intermRes)
	res := []string{}

	for _, ir := range intermRes {
		res = append(res, strings.Join(ir, "."))
	}

	return res
}

func restoreIpAddressesRecursive(s string, pos int, path []string, res *[][]string) {

	if pos >= len(s) || len(path) == 4 {
		if pos >= len(s) && len(path) == 4 {
			tmp := make([]string, len(path))
			copy(tmp, path)
			*res = append(*res, tmp)
		}
		return
	}

	for i := pos; i < len(s); i++ {
		oct := s[pos : i+1]

		//  cannot have leading zeros
		if len(oct) > 1 && oct[:1] == "0" {
			break
		}

		v, _ := strconv.ParseInt(oct, 10, 16)

		if v > 255 {
			break
		}

		path = append(path, oct)
		restoreIpAddressesRecursive(s, i+1, path, res)
		path = path[:len(path)-1]
	}

}
