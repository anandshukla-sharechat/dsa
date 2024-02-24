package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum(t *testing.T) {
	testcases := []struct {
		inputArr  []util.IntegerStruct
		targetSum int
		output    [][]int
	}{
		{inputArr: []util.IntegerStruct{{5, true}, {4, true}, {8, true}, {11, true}, {-1, false}, {13, true}, {4, true}, {7, true}, {2, true}, {-1, false}, {-1, false}, {-1, false}, {-1, false}, {5, true}, {1, true}}, targetSum: 22, output: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
	}
	for _, testcase := range testcases {
		root := util.LevelOrder(testcase.inputArr)
		assert.Equal(t, testcase.output, pathSum(root, testcase.targetSum))
	}
}
