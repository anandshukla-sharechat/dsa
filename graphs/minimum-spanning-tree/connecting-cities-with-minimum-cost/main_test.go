package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumCost(t *testing.T) {
	tests := []struct {
		n           int
		connections [][]int
		cost        int
	}{
		{n: 3, connections: [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}, cost: 6},
		{n: 4, connections: [][]int{{1, 2, 3}, {3, 4, 4}}, cost: -1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.cost, minimumCost(tt.n, tt.connections))
	}
}
