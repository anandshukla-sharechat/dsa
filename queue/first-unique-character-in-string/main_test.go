package main

import (
	heap_approach "first-unique-character-in-string/heap-approach"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	testCases := []struct {
		s   string
		res int
	}{
		{s: "leetcode", res: 0},
		{s: "loveleetcode", res: 2},
	}

	for _, tt := range testCases {
		assert.Equal(t, tt.res, firstUniqChar(tt.s))
		assert.Equal(t, tt.res, heap_approach.FirstUniqChar(tt.s))
	}
}
