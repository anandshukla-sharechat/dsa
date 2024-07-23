package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		x   int
		ans int
	}{
		{x: 2, ans: 1},
		{x: 10, ans: 3},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, mySqrt(tt.x))
	}
}
