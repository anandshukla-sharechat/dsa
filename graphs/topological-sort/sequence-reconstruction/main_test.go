package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sequenceReconstruction(t *testing.T) {
	testCases := []struct {
		nums             []int
		sequences        [][]int
		canBeConstructed bool
	}{
		{nums: []int{1, 2, 3}, sequences: [][]int{{1, 2}, {1, 3}}, canBeConstructed: false},
		{nums: []int{1, 2, 3}, sequences: [][]int{{1, 2}, {1, 3}, {2, 3}}, canBeConstructed: true},
		{nums: []int{1}, sequences: [][]int{{1}}, canBeConstructed: true},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.canBeConstructed, sequenceReconstruction(testCase.nums, testCase.sequences))
	}
}
