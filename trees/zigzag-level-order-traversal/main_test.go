package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	testCases := []struct {
		root *util.TreeNode
		ans  [][]int
	}{
		{root: util.LevelOrder([]util.IntegerStruct{{3, true}, {9, true}, {20, true}, {-1, false}, {-1, false}, {15, true}, {7, true}}), ans: [][]int{{3}, {20, 9}, {15, 7}}},
		{root: util.LevelOrder([]util.IntegerStruct{{1, true}}), ans: [][]int{{1}}},
		{root: nil, ans: [][]int{}},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.ans, zigzagLevelOrder(testCase.root))
	}
}
