package main

import (
	"fmt"
	"strings"
)

const alphabets = 26

//node for each children; i.e alphabets 1 - 26
type node struct {
	children [alphabets]*node
	isEnd    bool
}

// a particulat node
type trie struct {
	root *node
}

// initialize our tries
func InitTrie() *trie {
	trie := &trie{root: &node{}}
	return trie
}

// Insert puts a new word in the trie
func (t *trie) insert(s string) {
	w := strings.ToLower(s)
	currentNode := t.root

	for i := range w {
		childIndex := w[i] - 'a'

		if currentNode.children[childIndex] == nil {
			currentNode.children[childIndex] = &node{}
		}

		currentNode = currentNode.children[childIndex]
	}

	currentNode.isEnd = true
}

func (t *trie) search(s string) bool {
	w := strings.ToLower(s)
	currentNode := t.root

	for i := range w {
		childIndex := w[i] - 'a'

		if currentNode.children[childIndex] == nil {
			return false
		}

		currentNode = currentNode.children[childIndex]
	}

	return currentNode.isEnd
}

func main() {
	rootTrie := InitTrie()

	wordsToInsert := []string{
		"tree",
		"free",
		"trim",
		"trip",
		"fret",
		"frey",
	}

	for _, v := range wordsToInsert {
		rootTrie.insert(v)
	}

	fmt.Println(rootTrie.search("frey"))
}
