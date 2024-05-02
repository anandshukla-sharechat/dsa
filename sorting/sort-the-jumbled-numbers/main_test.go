package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortJumbled(t *testing.T) {
	tests := []struct {
		mapping []int
		nums    []int
		ans     []int
	}{
		{mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, nums: []int{991, 338, 38}, ans: []int{338, 38, 991}},
		{mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nums: []int{789, 456, 123}, ans: []int{123, 456, 789}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, sortJumbled(tt.mapping, tt.nums))
	}
}
