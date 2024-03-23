package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestIncreasingPath(t *testing.T) {
	tests := []struct {
		matrix [][]int
		ans    int
	}{
		{matrix: [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, ans: 4},
		{matrix: [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, ans: 4},
		{matrix: [][]int{{1}}, ans: 1},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.ans, longestIncreasingPath(tt.matrix))
	}
}
