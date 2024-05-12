package main

import (
	"container/list"
	"fmt"
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

func wordBreak(s string, wordDict []string) bool {
	trie := &TrieNode{}
	sLen := len(s)
	for _, word := range wordDict {
		trie.insert(word)
	}
	// apply BFS on trie
	q := list.New()
	q.PushBack(0)
	vis := make([]bool, sLen)
	for q.Len() > 0 {
		sz := q.Len()
		for ; sz > 0; sz-- {
			el := q.Remove(q.Front()).(int)
			if vis[el] {
				continue
			}
			vis[el] = true
			node := trie
			// search through trie
			for el < sLen && node.children[s[el]-'a'] != nil {
				node = node.children[s[el]-'a']
				if node.isEnd {
					q.PushBack(el + 1)
				}
				el++
			}
			if el == sLen && node.isEnd {
				return true
			}
		}
	}
	return false
}

// https://leetcode.com/problems/word-break/description/
func main() {
	s, wordDict := "leetcode", []string{"leet", "code"}
	ans := wordBreak(s, wordDict)
	fmt.Println(ans)
}
