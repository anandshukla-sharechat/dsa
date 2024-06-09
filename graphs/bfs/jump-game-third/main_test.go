package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canReach(t *testing.T) {
	tests := []struct {
		arr   []int
		start int
		ans   bool
	}{
		{arr: []int{4, 2, 3, 0, 3, 1, 2}, start: 5, ans: true},
		{arr: []int{4, 2, 3, 0, 3, 1, 2}, start: 0, ans: true},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, canReach(tt.arr, tt.start))
	}
}
