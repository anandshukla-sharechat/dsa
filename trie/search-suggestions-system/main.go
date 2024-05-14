package main

import (
	"fmt"
	"sort"
)

type TrieNode struct {
	children  [26]*TrieNode
	isEnd     bool
	wordsList []string
}

func (t *TrieNode) insert(word string) {
	for _, c := range word {
		el := c - 'a'
		if t.children[el] == nil {
			t.children[el] = &TrieNode{wordsList: make([]string, 0)}
		}
		t = t.children[el]
		t.wordsList = append(t.wordsList, word)
	}
	t.isEnd = true
}

func (t *TrieNode) search(searchWord string) [][]string {
	ans := make([][]string, len(searchWord))
	for i, c := range searchWord {
		el := c - 'a'
		if t.children[el] == nil {
			break
		}
		t = t.children[el]
		sort.Slice(t.wordsList, func(i, j int) bool {
			return t.wordsList[i] < t.wordsList[j]
		})
		ans[i] = t.wordsList[:min(len(t.wordsList), 3)]
	}
	return ans
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := &TrieNode{}
	for _, word := range products {
		trie.insert(word)
	}
	ans := trie.search(searchWord)
	return ans
}

func main() {
	products, searchedWord := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"
	ans := suggestedProducts(products, searchedWord)
	for _, el := range ans {
		fmt.Println(el)
	}
}
