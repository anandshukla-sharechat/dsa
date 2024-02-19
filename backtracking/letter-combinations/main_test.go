package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	testcases := []struct {
		digits string
		output []string
	}{
		{digits: "23", output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{digits: "", output: []string{}},
		{digits: "2", output: []string{"a", "b", "c"}},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.output, letterCombinations(testcase.digits))
	}
}
