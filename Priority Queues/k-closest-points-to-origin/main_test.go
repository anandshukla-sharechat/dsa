package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kClosest(t *testing.T) {
	tests := []struct {
		points [][]int
		k      int
		ans    [][]int
	}{
		{points: [][]int{{3, 3}, {5, -1}, {-2, 4}}, k: 2, ans: [][]int{{-2, 4}, {3, 3}}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, kClosest(tt.points, tt.k))
	}
}
