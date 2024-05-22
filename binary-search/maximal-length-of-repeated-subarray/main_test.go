package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findLength(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		ans   int
	}{
		{nums1: []int{1, 2, 3, 2, 1}, nums2: []int{3, 2, 1, 4, 7}, ans: 3},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, findLength(tt.nums1, tt.nums2))
	}
}
