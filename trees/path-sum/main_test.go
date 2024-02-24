package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	testcases := []struct {
		inputArr  []util.IntegerStruct
		targetSum int
		output    bool
	}{
		{inputArr: []util.IntegerStruct{{5, true}, {4, true}, {8, true}, {11, true}, {-1, false}, {13, true}, {4, true}, {7, true}, {2, true}, {-1, false}, {-1, false}, {-1, false}, {-1, false}, {-1, false}, {1, true}}, targetSum: 22, output: true},
	}

	for _, testcase := range testcases {
		root := util.LevelOrder(testcase.inputArr)
		assert.Equal(t, testcase.output, hasPathSum(root, testcase.targetSum))
	}
}
