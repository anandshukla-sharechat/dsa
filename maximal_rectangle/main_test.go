package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximalRectangle(t *testing.T) {
	tt := struct {
		matrix [][]byte
		output int
	}{
		matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
		output: 6,
	}
	assert.Equal(t, tt.output, maximalRectangle(tt.matrix))
}
