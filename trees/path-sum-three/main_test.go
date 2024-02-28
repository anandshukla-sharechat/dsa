package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum(t *testing.T) {
	testcases := []struct {
		nums      []util.IntegerStruct
		targetSum int
		ans       int
	}{
		{nums: []util.IntegerStruct{{10, true}, {5, true}, {-3, true}, {3, true}, {2, true}, {-1, false}, {11, true}, {3, true}, {-2, true}, {1, true}}, targetSum: 8, ans: 3},
	}

	for _, testcase := range testcases {
		root := util.LevelOrder(testcase.nums)
		assert.Equal(t, testcase.ans, pathSum(root, testcase.targetSum))
	}
}
