package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findOrder(t *testing.T) {
	testCases := []struct {
		numCourses    int
		prerequisites [][]int
		ans           []int
	}{
		{numCourses: 2, prerequisites: [][]int{{1, 0}}, ans: []int{0, 1}},
		{numCourses: 4, prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, ans: []int{0, 2, 1, 3}},
		{numCourses: 1, prerequisites: nil, ans: []int{0}},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.ans, findOrder(testCase.numCourses, testCase.prerequisites))
	}
}
