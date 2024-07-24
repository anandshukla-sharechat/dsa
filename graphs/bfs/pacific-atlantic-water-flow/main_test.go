package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pacificAtlantic(t *testing.T) {
	tests := []struct {
		height [][]int
		ans    [][]int
	}{
		{height: [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}, ans: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
		{height: [][]int{{1}}, ans: [][]int{{0, 0}}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, pacificAtlantic(tt.height))
	}
}
