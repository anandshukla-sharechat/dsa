package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	testcase := struct {
		candidate []int
		target    int
		res       [][]int
	}{
		candidate: []int{10, 1, 2, 7, 6, 1, 5},
		target:    8,
		res:       [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
	}
	assert.Equal(t, testcase.res, combinationSum2(testcase.candidate, testcase.target))
}
