package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	tests := []struct {
		graph [][]int
		ans   []int
	}{
		{graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}, ans: []int{2, 4, 5, 6}},
		{graph: [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}, ans: []int{4}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, eventualSafeNodes(tt.graph))
	}
}
