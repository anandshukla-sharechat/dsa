package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	Input  []int
	Output int
}

func Test_largestRectangleArea_CaseOne(t *testing.T) {
	testCaseOne := TestStruct{
		Input:  []int{2, 1, 5, 6, 2, 3},
		Output: 10,
	}
	assert.Equal(t, testCaseOne.Output, largestRectangleArea(testCaseOne.Input))
}

func Test_largestRectangleArea_Two(t *testing.T) {
	testCaseOne := TestStruct{
		Input:  []int{2, 4},
		Output: 4,
	}
	assert.Equal(t, testCaseOne.Output, largestRectangleArea(testCaseOne.Input))
}
