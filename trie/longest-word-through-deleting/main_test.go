package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findLongestWord(t *testing.T) {
	tests := []struct {
		words []string
		s     string
		ans   string
	}{
		{words: []string{"ale", "apple", "monkey", "plea"}, s: "abpcplea", ans: "apple"},
		{words: []string{"a", "b", "c"}, s: "abpcplea", ans: "a"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, findLongestWord(tt.s, tt.words))
	}
}
