package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canFinish(t *testing.T) {
	testCases := []struct {
		numCourses    int
		prerequisites [][]int
		output        bool
	}{
		{numCourses: 1, prerequisites: [][]int{{1, 0}}, output: true},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.output, canFinish(testCase.numCourses, testCase.prerequisites))
	}
}
