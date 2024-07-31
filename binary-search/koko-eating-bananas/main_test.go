package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minEatingSpeed(t *testing.T) {
	tests := []struct {
		piles []int
		h     int
		ans   int
	}{
		{piles: []int{3, 6, 7, 11}, h: 8, ans: 4},
		{piles: []int{30, 11, 23, 4, 20}, h: 5, ans: 30},
		{piles: []int{30, 11, 23, 4, 20}, h: 6, ans: 23},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, minEatingSpeed(tt.piles, tt.h))
	}
}
