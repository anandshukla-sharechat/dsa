package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tt := struct {
		input int
		res   []string
	}{
		input: 3,
		res:   []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	}

	assert.Equal(t, tt.res, generateParenthesis(tt.input))
}
