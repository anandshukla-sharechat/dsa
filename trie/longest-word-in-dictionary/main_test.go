package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestWord(t *testing.T) {
	tests := []struct {
		words []string
		ans   string
	}{
		{words: []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, ans: "apple"},
		{words: []string{"w", "wo", "wor", "worl", "world"}, ans: "world"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, longestWord(tt.words))
	}
}
