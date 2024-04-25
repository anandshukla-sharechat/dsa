package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minCostToSupplyWater(t *testing.T) {
	tests := []struct {
		n     int
		wells []int
		pipes [][]int
		cost  int
	}{
		{n: 3, wells: []int{1, 2, 2}, pipes: [][]int{{1, 2, 1}, {2, 3, 1}}, cost: 3},
		{n: 2, wells: []int{1, 1}, pipes: [][]int{{1, 2, 1}, {1, 2, 2}}, cost: 2},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.cost, minCostToSupplyWater(tt.n, tt.wells, tt.pipes))
	}
}
