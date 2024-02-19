package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	testcases := []struct {
		candidates []int
		target     int
		res        [][]int
	}{
		{candidates: []int{2, 3, 6, 7}, target: 7, res: [][]int{{2, 2, 3}, {7}}},
		{candidates: []int{2, 3, 5}, target: 8, res: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.res, combinationSum(testcase.candidates, testcase.target))
	}
}
