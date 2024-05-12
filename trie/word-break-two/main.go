package main

import (
	"container/list"
	"fmt"
	"strings"
)

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
	idx      int
}

type WordJoin struct {
	idx      int
	wordlist []string
}

func (t *TrieNode) insert(word string, idx int) {
	for _, c := range word {
		el := c - 'a'
		if t.children[el] == nil {
			t.children[el] = &TrieNode{}
		}
		t = t.children[el]
	}
	t.isEnd = true
	t.idx = idx
}

func wordBreak(s string, wordDict []string) []string {
	trie := &TrieNode{}
	for idx, word := range wordDict {
		trie.insert(word, idx)
	}
	q := list.New()
	sLen := len(s)
	ans := make([]string, 0)
	q.PushBack(WordJoin{idx: 0, wordlist: make([]string, 0)})
	for q.Len() > 0 {
		sz := q.Len()
		for ; sz > 0; sz-- {
			frontEl := q.Remove(q.Front()).(WordJoin)
			if frontEl.idx >= sLen {
				if frontEl.idx == sLen {
					constructedString := strings.Join(frontEl.wordlist, " ")
					ans = append(ans, constructedString)
				}
				continue
			}
			node := trie
			for frontEl.idx < sLen && node.children[s[frontEl.idx]-'a'] != nil {
				node = node.children[s[frontEl.idx]-'a']
				if node.isEnd {
					q.PushBack(WordJoin{idx: frontEl.idx + 1, wordlist: append(frontEl.wordlist, wordDict[node.idx])})
				}
				frontEl.idx++
			}
		}
	}
	return ans
}

// https://leetcode.com/problems/word-break-ii/description/
// 25 / 27 testcases passed
func main() {
	s, wordDict := "catsanddog", []string{"cat", "cats", "and", "sand", "dog"}
	ans := wordBreak(s, wordDict)
	for _, el := range ans {
		fmt.Println(el)
	}
}
