package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
	idx      int
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

func (t *TrieNode) searchWord(word string) int {
	for _, c := range word {
		el := c - 'a'
		if t.children[el] == nil {
			return -1
		}
		if t.children[el].isEnd {
			return t.children[el].idx
		}
		t = t.children[el]
	}
	return -1
}

func addSeparator(idx int, l int) string {
	if idx == l-1 {
		return ""
	} else {
		return " "
	}
}

func replaceWords(dictionary []string, sentence string) string {
	trie := &TrieNode{}
	for idx, word := range dictionary {
		trie.insert(word, idx)
	}
	sentWords := strings.Split(sentence, " ")
	ans, l := "", len(sentWords)
	for i, word := range sentWords {
		idx := trie.searchWord(word)
		if idx == -1 {
			ans += word + addSeparator(i, l)
		} else {
			ans += dictionary[idx] + addSeparator(i, l)
		}
	}
	return ans
}

// https://leetcode.com/problems/replace-words/
func main() {
	dictionary, sentence := []string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"
	ans := replaceWords(dictionary, sentence)
	fmt.Println(ans)
}
