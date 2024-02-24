package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	testcases := []struct {
		NumberArr []util.IntegerStruct
		Output    bool
	}{
		{NumberArr: []util.IntegerStruct{{2, true}, {1, true}, {3, true}}, Output: true},
		{NumberArr: []util.IntegerStruct{{5, true}, {1, true}, {4, true}, {-1, false}, {-1, false}, {3, true}, {6, true}}, Output: false},
	}

	for _, testcase := range testcases {
		root := util.LevelOrder(testcase.NumberArr)
		assert.Equal(t, testcase.Output, isValidBST(root))
	}

}
