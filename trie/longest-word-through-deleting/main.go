package main

import (
	"fmt"
	"math"
)

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func (t *TrieNode) insert(word string) {
	for _, c := range word {
		el := c - 'a'
		if t.children[el] == nil {
			t.children[el] = &TrieNode{}
		}
		t = t.children[el]
	}
	t.isEnd = true
}

func (t *TrieNode) search(str, word string) bool {
	idx := 0
	for i := 0; i < len(word); i++ {
		el := word[i] - 'a'
		for idx < len(str) && (str[idx]-'a') != el {
			idx++
		}
		if idx >= len(str) {
			return false
		}
		t = t.children[el]
		idx++
	}
	return true
}

func findLongestWord(s string, words []string) string {
	trie := &TrieNode{}
	for _, word := range words {
		trie.insert(word)
	}
	ans, maxL := "", math.MinInt
	for _, word := range words {
		if trie.search(s, word) {
			if maxL <= len(word) {
				if maxL < len(word) {
					ans = word
					maxL = len(ans)
				} else {
					if ans > word {
						ans = word
					}
				}
			}
		}
	}
	return ans
}

// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/
func main() {
	s, words := "abpcplea", []string{"ale", "apple", "monkey", "plea"}
	ans := findLongestWord(s, words)
	fmt.Println(ans)
}
