package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	testcase := struct {
		root   *util.TreeNode
		output []int
	}{
		root:   util.LevelOrder([]util.IntegerStruct{{1, true}, {2, true}, {3, true}, {-1, false}, {5, true}, {-1, false}, {4, true}}),
		output: []int{1, 3, 4},
	}
	assert.Equal(t, testcase.output, rightSideView(testcase.root))
}
