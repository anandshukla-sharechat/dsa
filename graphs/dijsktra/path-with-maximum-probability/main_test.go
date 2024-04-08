package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProbability(t *testing.T) {
	tests := []struct {
		n           int
		edges       [][]int
		succProb    []float64
		start_node  int
		end_node    int
		maximumProb float64
	}{
		{n: 3, edges: [][]int{{0, 1}, {1, 2}, {0, 2}}, succProb: []float64{0.5, 0.5, 0.2}, start_node: 0, end_node: 2, maximumProb: 0.25},
		{n: 3, edges: [][]int{{0, 1}, {1, 2}, {0, 2}}, succProb: []float64{0.5, 0.5, 0.3}, start_node: 0, end_node: 2, maximumProb: 0.3},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.maximumProb, maxProbability(tt.n, tt.edges, tt.succProb, tt.start_node, tt.end_node))
	}
}
