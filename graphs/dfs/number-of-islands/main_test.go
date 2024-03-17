package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numIslands(t *testing.T) {
	testCases := []struct {
		grid   [][]byte
		output int
	}{
		{grid: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, output: 3},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.output, numIslands(testCase.grid))
	}
}
