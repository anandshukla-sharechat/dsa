package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		ans      bool
	}{
		{s: "leetcode", wordDict: []string{"leet", "code"}, ans: true},
		{s: "applepenapple", wordDict: []string{"apple", "pen"}, ans: true},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, wordBreak(tt.s, tt.wordDict))
	}
}
