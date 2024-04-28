package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		} else {
			node = node.children[c]
		}
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		} else {
			node = node.children[c]
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
