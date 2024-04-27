package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximumSafenessFactor(t *testing.T) {
	tests := []struct {
		grid           [][]int
		safestDistance int
	}{
		{grid: [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}, safestDistance: 2},
		{grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}, safestDistance: 0},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.safestDistance, maximumSafenessFactor(tt.grid))
	}
}
