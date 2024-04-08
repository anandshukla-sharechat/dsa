package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minCostConnectPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		ans    int
	}{
		{points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, ans: 20},
		{points: [][]int{{3, 12}, {-2, 5}, {-4, 1}}, ans: 18},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, minCostConnectPoints(tt.points))
	}
}
