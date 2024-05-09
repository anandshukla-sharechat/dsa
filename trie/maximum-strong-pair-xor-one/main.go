package main

import (
	"fmt"
)

type TrieNode struct {
	children [2]*TrieNode
	isEnd    bool
}

func reverseArray(nums []int) []int {
	rev := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		rev = append(rev, nums[i])
	}
	return rev
}

func returnBitSet(num int) [32]int {
	lsbBits := make([]int, 0)
	var bitset [32]int
	idx := 0
	for i := num; i > 0; i = i / 2 {
		lsbBits = append(lsbBits, i%2)
	}
	lsbBits = reverseArray(lsbBits)

	for i := 32 - len(lsbBits); i < 32; i++ {
		bitset[i] = lsbBits[idx]
		idx++
	}
	return bitset
}

func (t *TrieNode) insert(bitset [32]int) {
	for i := 0; i < 32; i++ {
		if t.children[bitset[i]] == nil {
			t.children[bitset[i]] = &TrieNode{}
		}
		t = t.children[bitset[i]]
	}
	t.isEnd = true
}

func revBit(bit int) int {
	if bit == 1 {
		return 0
	} else {
		return 1
	}
}

func (t *TrieNode) getMaxXor(bitset [32]int) int {
	xor := 0
	for i := 0; i < 32; i++ {
		reverseBit := revBit(bitset[i])
		if t.children[reverseBit] == nil {
			t = t.children[bitset[i]]
		} else {
			xor += 1 << (31 - i)
			t = t.children[reverseBit]
		}
	}
	return xor
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

func maximumStrongPairXor(nums []int) int {
	trie := &TrieNode{}
	storeBitSets := make([][32]int, 0)
	for _, el := range nums {
		bitset := returnBitSet(el)
		trie.insert(bitset)
		storeBitSets = append(storeBitSets, bitset)
	}
	maxXor := 0
	for idx, val := range nums {
		xor := trie.getMaxXor(storeBitSets[idx])
		otherval := xor ^ val
		if abs(otherval-val) <= min(otherval, val) && maxXor < xor {
			maxXor = xor
		}
	}
	return maxXor
}

/*
https://leetcode.com/problems/maximum-strong-pair-xor-i/
782 / 860 testcases passed
*/
func main() {
	nums := []int{5, 6, 25, 30}
	ans := maximumStrongPairXor(nums)
	fmt.Println(ans)
}
