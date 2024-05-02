package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	tests := []struct {
		n   int
		ans []int
	}{
		{n: 13, ans: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{n: 2, ans: []int{1, 2}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, lexicalOrder(tt.n))
		assert.Equal(t, tt.ans, lexicalOrderBruteForce(tt.n))
	}
}
