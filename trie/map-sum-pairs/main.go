package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
	freq     int
}
type MapSum struct {
	root *TrieNode
	mp   map[string]int
}

func Constructor() MapSum {
	return MapSum{root: &TrieNode{}, mp: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	node := this.root
	olderVal := this.mp[key]
	for _, el := range key {
		idx := el - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
		node.freq += val - olderVal
	}
	this.mp[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	node := this.root
	for _, el := range prefix {
		idx := el - 'a'
		if node.children[idx] == nil {
			return 0
		}
		node = node.children[idx]
	}
	return node.freq
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

// https://leetcode.com/problems/map-sum-pairs/description/
func main() {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	fmt.Println(mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	fmt.Println(mapSum.Sum("ap"))
}
