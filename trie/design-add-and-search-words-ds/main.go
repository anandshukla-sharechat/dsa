package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (node *TrieNode) search(word string) bool {
	for i, c := range word {
		if c != '.' {
			if _, ok := node.children[c]; !ok {
				return false
			} else {
				node = node.children[c]
			}
		} else {
			for _, child := range node.children {
				if child.search(word[i+1:]) {
					return true
				}
			}
			return false
		}
	}
	return node.isEnd
}

func (this *WordDictionary) Search(word string) bool {
	node := this.root
	return node.search(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad"))
	fmt.Println(wordDictionary.Search("bad"))
	fmt.Println(wordDictionary.Search(".ad"))
	fmt.Println(wordDictionary.Search("b.."))
}
