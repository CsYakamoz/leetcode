package golang

import (
	_ "fmt"
	"strconv"
)

type TrieNode struct {
	isWorld  bool
	children []*TrieNode
}

func suggestedProducts(products []string, searchWord string) [][]string {
	root := newTrieNode()
	for _, word := range products {
		insertWordIntoTrieTree(root, word)
	}

	return search(searchWord, root)
}

func insertWordIntoTrieTree(curr *TrieNode, word string) {
	for _, c := range word {
		idx := int(c - 'a')

		if curr.children[idx] == nil {
			curr.children[idx] = newTrieNode()
		}

		curr = curr.children[idx]
	}

	curr.isWorld = true
}

func newTrieNode() *TrieNode {
	return &TrieNode{isWorld: false, children: make([]*TrieNode, 26)}
}

func search(searchWord string, root *TrieNode) [][]string {
	ans := make([][]string, 0, len(searchWord))

	prefix := ""
	for _, c := range searchWord {
		prefix += string(c)
		result := []string{}

		dfsWithPrefix(root, prefix, &result)
		ans = append(ans, result)
	}

	return ans
}

func dfsWithPrefix(root *TrieNode, prefix string, result *[]string) {
	if len(*result) > 3 {
		return
	}

	if root.isWorld {
		*result = append(*result, prefix)
	}

	for idx, next := range root.children {
		if next == nil {
			continue
		}

		prefix += strconv.Itoa(97 + idx)
		dfsWithPrefix(next, prefix, result)
		prefix = prefix[:len(prefix)-1]
	}
}
