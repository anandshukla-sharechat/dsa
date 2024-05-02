package main

import (
	"fmt"
	"sort"
	"strconv"
)

func lexicalOrderBruteForce(n int) []int {
	l := make([]string, 0)
	for i := 1; i <= n; i++ {
		l = append(l, fmt.Sprintf("%d", i))
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	ans := make([]int, 0)
	for _, el := range l {
		v, _ := strconv.Atoi(el)
		ans = append(ans, v)
	}
	return ans
}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children [10]*TrieNode
	isEnd    bool
}

func (t Trie) Insert(num []int) {
	node := t.root
	for _, el := range num {
		if node.children[el] == nil {
			node.children[el] = &TrieNode{}
		}
		node = node.children[el]
	}
	node.isEnd = true
}

func (t *TrieNode) ConstructSortedList(ans *[]int, val int) {
	if t.isEnd {
		*ans = append(*ans, val)
	}
	val = val * 10
	for i := 0; i < 10; i++ {
		if t.children[i] != nil {
			val = val + i
			t.children[i].ConstructSortedList(ans, val)
			val = val - i
		}
	}
}

func lexicalOrder(n int) []int {
	trie := Trie{root: &TrieNode{}}
	ans := make([]int, 0)
	for u := 1; u <= n; u++ {
		temp := make([]int, 0)
		p := 0
		for i := u; i != 0 || p >= 0; i = i / 10 {
			temp = append(temp, i%10)
			p--
		}
		trie.Insert(reverseIntArr(temp))
	}
	trie.root.ConstructSortedList(&ans, 0)
	return ans
}

func reverseIntArr(arr []int) []int {
	ans := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		ans = append(ans, arr[i])
	}
	return ans
}

// https://leetcode.com/problems/lexicographical-numbers/description/
func main() {
	n := 13
	ans := lexicalOrder(n)
	fmt.Println(ans)
}
