package main

/*
func main() {
	words := []string{"abc", "adb", "a", "bcm", "ttm", "llc", "ll"}
	a := NewAutocomplete(words)

	fmt.Println(a.findAllWordsByPrefix("adbc"))
}
*/

type node struct {
	data     string
	isWord   bool
	children map[rune]*node
}

type Autocomplete struct {
	trie *node
}

func NewAutocomplete(words []string) *Autocomplete {

	a := &Autocomplete{
		trie: &node{
			data:     "",
			children: make(map[rune]*node),
		},
	}

	for _, w := range words {
		a.insertPrefix(w)
	}

	return a
}

func (a *Autocomplete) insertPrefix(prefix string) {

	curr := a.trie

	for i := 0; i < len(prefix); i++ {
		c := rune(prefix[i])
		if _, ok := curr.children[c]; !ok {

			nwnode := &node{
				data:     prefix[:i+1],
				children: map[rune]*node{},
			}

			curr.children[c] = nwnode
		}

		// check if is a word
		if i == len(prefix)-1 {
			curr.children[c].isWord = true
		}

		curr = curr.children[c]
	}

}

func (a *Autocomplete) findAllWordsByPrefix(prefix string) []string {

	curr := a.trie

	for i := 0; i < len(prefix); i++ {
		c := rune(prefix[i])
		if _, ok := curr.children[c]; ok {
			curr = curr.children[c]
		} else {
			return []string{}
		}
	}

	var result []string

	a.findAllWords(curr, &result)

	return result
}

func (a *Autocomplete) findAllWords(node *node, result *[]string) {

	if node.isWord {
		*result = append(*result, node.data)
	}
	for _, n := range node.children {
		a.findAllWords(n, result)
	}

}
