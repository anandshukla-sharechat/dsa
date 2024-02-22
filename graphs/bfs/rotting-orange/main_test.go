package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_orangesRotting(t *testing.T) {
	testcases := []struct {
		grid   [][]int
		output int
	}{
		{grid: [][]int{{0, 1, 1, 2}, {0, 1, 1, 0}, {1, 1, 2, 0}}, output: 2},
		{grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, output: 4},
	}
	for _, testcase := range testcases {
		assert.Equal(t, testcase.output, orangesRotting(testcase.grid))
	}
}
