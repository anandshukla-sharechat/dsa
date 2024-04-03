package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	tests := []struct {
		isConnected [][]int
		ans         int
	}{
		{isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, ans: 2},
		{isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, ans: 3},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, findCircleNum(tt.isConnected))
	}
}
