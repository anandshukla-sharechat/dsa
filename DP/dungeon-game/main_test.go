package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculateMinimumHP(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Output int
	}{
		{Input: [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, Output: 7},
		{Input: [][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}, Output: 3},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.Output, calculateMinimumHP(testCase.Input))
	}
}
