package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		words []string
		k     int
		ans   []string
	}{
		{words: []string{"i", "love", "leetcode", "i", "love", "coding"}, k: 2, ans: []string{"i", "love"}},
		{words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, k: 4, ans: []string{"the", "is", "sunny", "day"}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, topKFrequent(tt.words, tt.k))
	}
}
