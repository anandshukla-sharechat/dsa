package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LRU(t *testing.T) {
	expectedOutput := []int{1, -1, -1, 3, 4}
	res := make([]int, 0)
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)                 // cache is {1=1}
	lRUCache.Put(2, 2)                 // cache is {1=1, 2=2}
	res = append(res, lRUCache.Get(1)) // return 1
	lRUCache.Put(3, 3)                 // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	res = append(res, lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)                 // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	res = append(res, lRUCache.Get(1)) // return -1 (not found)
	res = append(res, lRUCache.Get(3)) // return 3
	res = append(res, lRUCache.Get(4)) // return 4

	assert.Equal(t, expectedOutput, res)
}
