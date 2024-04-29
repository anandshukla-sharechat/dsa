package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaximumXOR(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{nums: []int{3, 10, 5, 25, 2, 8}, ans: 28},
		{nums: []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}, ans: 127},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, findMaximumXOR(tt.nums))
	}
}
