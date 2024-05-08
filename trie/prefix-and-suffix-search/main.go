package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	Index    int
}

type WordFilter struct {
	prefix *TrieNode
	suffix *TrieNode
}

func (trie *TrieNode) insert(word string, a, b, update, idx int) {

	for i := a; ; i = i + update {
		if update == 1 {
			if i > b {
				break
			}
		} else {
			if i < b {
				break
			}
		}
		el := word[i] - 'a'
		if trie.children[el] == nil {
			trie.children[el] = &TrieNode{}
		}
		trie = trie.children[el]
	}
	trie.Index = idx + 1
}

func (trie *TrieNode) getListOfIndices(freq *map[int]int, pref_suffix string) {

	for _, c := range pref_suffix {
		el := c - 'a'
		if trie.children[el] == nil {
			return
		}
		trie = trie.children[el]
	}
	if trie.Index > 0 {
		(*freq)[trie.Index]++
	} else {
		trie.getIndices(freq)
	}
}

func (trie *TrieNode) getIndices(freq *map[int]int) {
	if trie == nil {
		return
	}
	for _, child := range trie.children {
		if trie.Index > 0 {
			(*freq)[trie.Index]++
		} else {
			child.getIndices(freq)
		}
	}
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Constructor(words []string) WordFilter {
	prefix := &TrieNode{}
	suffix := &TrieNode{}
	for idx, word := range words {
		prefix.insert(word, 0, len(word)-1, 1, idx)
		suffix.insert(word, len(word)-1, 0, -1, idx)
	}
	return WordFilter{prefix: prefix, suffix: suffix}
}

func (this *WordFilter) F(pref string, suff string) int {
	mp_pref := make(map[int]int)
	mp_suff := make(map[int]int)
	this.prefix.getListOfIndices(&mp_pref, pref)
	this.suffix.getListOfIndices(&mp_suff, reverseString(suff))
	maxIdx := -1
	for k, v := range mp_pref {
		if v >= 1 && mp_suff[k] >= 1 {
			maxIdx = max(maxIdx, k)
		}
	}
	if maxIdx == -1 {
		return maxIdx
	} else {
		return maxIdx - 1
	}
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */

/*
https://leetcode.com/problems/prefix-and-suffix-search/description/
Passed 12/17 testcases
*/
func main() {
	wordFilter := Constructor([]string{"apple"})
	ans := wordFilter.F("b", "e")
	fmt.Println(ans)
}
