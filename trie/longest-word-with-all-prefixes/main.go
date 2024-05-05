package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func (t *TrieNode) insert(word string) {
	for _, w := range word {
		el := w - 'a'
		if t.children[el] == nil {
			t.children[el] = &TrieNode{}
		}
		t = t.children[el]
	}
	t.isEnd = true
}

func (t *TrieNode) sumOfPrefixCounts(word string) int {
	ct := 0
	for _, w := range word {
		el := w - 'a'
		t = t.children[el]
		if t.isEnd {
			ct++
		}
	}
	return ct
}

func longestWord(words []string) string {
	trie := &TrieNode{}
	for _, word := range words {
		trie.insert(word)
	}
	maxL, ans := 0, ""
	for _, word := range words {
		ct := trie.sumOfPrefixCounts(word)
		if ct == len(word) && ct >= maxL {
			if ct > maxL {
				maxL = ct
				ans = word
			} else {
				if word < ans {
					ans = word
				}
			}
		}
	}
	return ans
}

/*
Given an array of strings words, find the longest string in words such that every prefix of it is also in words.

For example, let words = ["a", "app", "ap"]. The string "app" has prefixes "ap" and "a", all of which are in words.
Return the string described above. If there is more than one string with the same length, return the lexicographically smallest one, and if no string exists, return "".
https://leetcode.com/problems/longest-word-with-all-prefixes/description/
*/
func main() {
	words := []string{"k", "ki", "kir", "kira", "kiran"}
	ans := longestWord(words)
	fmt.Println(ans)
}
