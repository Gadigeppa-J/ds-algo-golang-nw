package main

import (
	"fmt"
	"strings"

	"github.com/Gadigeppa-J/ds-algo-golang-nw/utils"
)

func main() {
	document := "Every book is a quotation; and every house is a quotation out of all forests, and mines, and stone quarries; and every man is a quotation from all his ancestors. "

	fmt.Println(WordCountEngine(document))
}

func WordCountEngine(document string) [][]string {

	words := tokenizeString(document)

	m := make(map[string]int)
	ordersKeys := []string{}

	largest := 0

	for i := 0; i < len(words); i++ {

		if strings.TrimSpace(words[i]) == "" {
			continue
		}

		if m[words[i]] == 0 {
			ordersKeys = append(ordersKeys, words[i])
		}
		m[words[i]] = m[words[i]] + 1
		largest = utils.Max(largest, m[words[i]])
	}

	arr := make([][]string, largest+1)

	for _, k := range ordersKeys {
		tmp := arr[m[k]]

		if tmp == nil {
			tmp = []string{}
		}

		tmp = append(tmp, k)
		arr[m[k]] = tmp
	}

	res := [][]string{}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == nil || len(arr[i]) == 0 {
			continue
		}

		for j := 0; j < len(arr[i]); j++ {
			res = append(res, []string{arr[i][j], fmt.Sprint(i)})
		}
	}

	return res
}

func tokenizeString(document string) []string {

	document = strings.ToLower(document)

	//punctuations := "!\"#$%&'()*+,-./:;?@[\\]^_`{|}~"

	// remove punctuations
	/*for i := 0; i < len(punctuations); i++ {
		document = strings.ReplaceAll(document, string(punctuations[i]), "")
	}*/

	str := ""

	for i := 0; i < len(document); i++ {

		if (document[i] >= 'a' && document[i] <= 'z') || document[i] == ' ' {
			str = str + string(document[i])
		}

	}

	return strings.Split(str, " ")
}
