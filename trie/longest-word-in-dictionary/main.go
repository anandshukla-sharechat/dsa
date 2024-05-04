package main

import (
	"fmt"
	"math"
)

type TrieNode struct {
	children [26]*TrieNode
	IsEnd    bool
}

func InitialiseDS() *TrieNode {
	return &TrieNode{}
}

func (t *TrieNode) insert(word string) {
	for _, c := range word {
		el := c - 'a'
		if t.children[el] == nil {
			t.children[el] = &TrieNode{}
		}
		t = t.children[el]
	}
	t.IsEnd = true
}

func (t *TrieNode) numberOfEndingChar(word string) int {
	ct := 0
	for _, c := range word {
		el := c - 'a'
		t = t.children[el]
		if t.IsEnd {
			ct++
		}
	}
	return ct
}

func longestWord(words []string) string {
	trie := InitialiseDS()
	for _, word := range words {
		trie.insert(word)
	}
	maxL, ans := math.MinInt, ""
	for _, word := range words {
		wLength := trie.numberOfEndingChar(word)
		if len(word) == wLength {
			if maxL <= wLength {
				if maxL == wLength {
					if ans > word {
						ans = word
					}
				} else {
					maxL = wLength
					ans = word
				}
			}
		}
	}
	return ans
}

// https://leetcode.com/problems/longest-word-in-dictionary/description/
func main() {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	ans := longestWord(words)
	fmt.Println(ans)
}
