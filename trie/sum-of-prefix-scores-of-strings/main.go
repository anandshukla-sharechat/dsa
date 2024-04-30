package main

import "fmt"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
	freq     int
}

func InititialiseTrie() Trie {
	return Trie{root: &TrieNode{}}
}

func (t Trie) insert(word string) {
	node := t.root
	for _, el := range word {
		idx := el - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
		node.freq++
	}
	node.isEnd = true
}

func (t Trie) search(p string) int {
	node := t.root
	ans := 0
	for _, el := range p {
		idx := el - 'a'
		if node.children[idx] == nil {
			return 0
		}
		node = node.children[idx]
		ans += node.freq
	}
	return ans
}

func sumPrefixScores(words []string) []int {
	trie := InititialiseTrie()
	n := len(words)
	for _, word := range words {
		trie.insert(word)
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = trie.search(words[i])
	}
	return ans
}

func main() {
	words := []string{"abc", "ab", "bc", "b"}
	ans := sumPrefixScores(words)
	fmt.Println(ans)
}
