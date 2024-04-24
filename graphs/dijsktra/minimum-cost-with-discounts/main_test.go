package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumCost(t *testing.T) {
	tests := []struct {
		n         int
		highways  [][]int
		discounts int
		cost      int
	}{
		{n: 5, highways: [][]int{{0, 1, 4}, {2, 1, 3}, {1, 4, 11}, {3, 2, 3}, {3, 4, 2}}, discounts: 1, cost: 9},
		{n: 2, highways: [][]int{{1, 0, 0}}, discounts: 0, cost: 0},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.cost, minimumCost(tt.n, tt.highways, tt.discounts))
	}
}
