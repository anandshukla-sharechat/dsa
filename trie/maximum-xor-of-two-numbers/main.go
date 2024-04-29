package main

import (
	"fmt"
)

type TrieNode struct {
	children [2]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func reverseInts(input []int) []int {
	newArr := make([]int, 0)
	for i := len(input) - 1; i >= 0; i-- {
		newArr = append(newArr, input[i])
	}
	return newArr
}

func constructBitSet(num int) [64]int {
	var bitset [64]int
	otherBits := make([]int, 0)
	for num > 0 {
		rem := num % 2
		otherBits = append(otherBits, rem)
		num = num / 2
	}
	otherBits = reverseInts(otherBits)
	idx := 0
	for i := 64 - len(otherBits); i < 64; i++ {
		bitset[i] = otherBits[idx]
		idx++
	}
	return bitset
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (t Trie) insert(nums [64]int) {
	node := t.root
	for _, el := range nums {
		if node.children[el] == nil {
			node.children[el] = &TrieNode{}
		}
		node = node.children[el]
	}
}

func revBit(n int) int {
	if n == 1 {
		return 0
	} else {
		return 1
	}
}

func (t *Trie) maximumXOR(nums [64]int) int {
	node := t.root
	idx := 63
	ans := 0
	for _, el := range nums {
		if node.children[revBit(el)] != nil {
			ans += 1 << idx
			node = node.children[revBit(el)]
		} else {
			node = node.children[el]
		}
		idx--
	}
	return ans
}

func findMaximumXOR(nums []int) int {
	trie := Constructor()
	bitsetArr := make([][64]int, 0)
	for _, el := range nums {
		bitset := constructBitSet(el)
		bitsetArr = append(bitsetArr, bitset)
		trie.insert(bitset)
	}
	ans := 0
	for _, bitset := range bitsetArr {
		ans = max(ans, trie.maximumXOR(bitset))
	}
	return ans
}

/*
https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/description/
Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 <= i <= j < n.
*/
func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	ans := findMaximumXOR(nums)
	fmt.Println(ans)
}
