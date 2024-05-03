package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type StreamChecker struct {
	root *TrieNode
	str  string
}

func (this *StreamChecker) insert(word string) {
	node := this.root
	for i := len(word) - 1; i >= 0; i-- {
		el := word[i] - 'a'
		if node.children[el] == nil {
			node.children[el] = &TrieNode{}
		}
		node = node.children[el]
	}
	node.isEnd = true
}

func Constructor(words []string) StreamChecker {
	trie := &StreamChecker{root: &TrieNode{}, str: ""}
	for _, word := range words {
		trie.insert(word)
	}
	return *trie
}

func (this *StreamChecker) Query(letter byte) bool {
	this.str = this.str + string(letter)
	node := this.root
	for i := len(this.str) - 1; i >= 0; i-- {
		el := this.str[i] - 'a'
		if node.children[el] == nil {
			return false
		} else {
			node = node.children[el]
			if node.isEnd {
				return true
			}
		}
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

func main() {
	streamChecker := Constructor([]string{"cd", "f", "kl"})
	fmt.Println(streamChecker.Query('a')) // return False
	fmt.Println(streamChecker.Query('b')) // return False
	fmt.Println(streamChecker.Query('c')) // return False
	fmt.Println(streamChecker.Query('d')) // return True, because 'cd' is in the wordlist
	fmt.Println(streamChecker.Query('e')) // return False
	fmt.Println(streamChecker.Query('f')) // return True, because 'f' is in the wordlist
	fmt.Println(streamChecker.Query('g')) // return False
	fmt.Println(streamChecker.Query('h')) // return False
	fmt.Println(streamChecker.Query('i')) // return False
	fmt.Println(streamChecker.Query('j')) // return False
	fmt.Println(streamChecker.Query('k')) // return False
	fmt.Println(streamChecker.Query('l')) // return True, because 'kl' is in the wordlist
}
