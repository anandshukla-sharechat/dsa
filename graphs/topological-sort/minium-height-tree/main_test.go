package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	testCases := []struct {
		n      int
		edges  [][]int
		output []int
	}{
		{n: 4, edges: [][]int{{1, 0}, {1, 2}, {1, 3}}, output: []int{1}},
		{n: 6, edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}, output: []int{3, 4}},
		{n: 1, edges: nil, output: []int{0}},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.output, findMinHeightTrees(testCase.n, testCase.edges))
	}
}
