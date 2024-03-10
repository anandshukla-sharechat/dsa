package main

import (
	"github.com/anandshukla-sharechat/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	rootTestCase := util.LevelOrder([]util.IntegerStruct{{3, true}, {5, true}, {1, true}, {6, true}, {2, true}, {0, true}, {8, true}, {-1, false}, {-1, false}, {7, true}, {4, true}})

	// Test Case One
	ansNodeTestCaseOne := lowestCommonAncestor(rootTestCase, rootTestCase.Left, rootTestCase.Right)
	assert.Equal(t, rootTestCase, ansNodeTestCaseOne)

	// Test Case Two
	ansNodeTestCaseTwo := lowestCommonAncestor(rootTestCase, rootTestCase.Left, rootTestCase.Left.Right.Right)
	assert.Equal(t, rootTestCase.Left, ansNodeTestCaseTwo)
}
