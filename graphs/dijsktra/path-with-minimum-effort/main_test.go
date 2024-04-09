package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumEffortPath(t *testing.T) {
	tests := []struct {
		heights [][]int
		ans     int
	}{
		{heights: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}, ans: 2},
		{heights: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}, ans: 1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, minimumEffortPath(tt.heights))
	}
}
