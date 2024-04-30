package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumPrefixScores(t *testing.T) {
	tests := []struct {
		words []string
		ans   []int
	}{
		{words: []string{"abc", "ab", "bc", "b"}, ans: []int{5, 4, 3, 2}},
		{words: []string{"abcd"}, ans: []int{4}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, sumPrefixScores(tt.words))
	}
}
