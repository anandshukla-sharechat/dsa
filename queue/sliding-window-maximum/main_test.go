package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	testcases := []struct {
		nums   []int
		k      int
		output []int
	}{
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, output: []int{3, 3, 5, 5, 6, 7}},
		{nums: []int{1, -1}, k: 1, output: []int{1, -1}},
	}
	for _, testcase := range testcases {
		assert.Equal(t, testcase.output, maxSlidingWindow(testcase.nums, testcase.k))
	}
}
