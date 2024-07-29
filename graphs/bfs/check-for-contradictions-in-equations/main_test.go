package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkContradictions(t *testing.T) {
	tests := []struct {
		equations [][]string
		values    []float64
		ans       bool
	}{
		{equations: [][]string{{"a", "b"}, {"b", "c"}, {"a", "c"}}, values: []float64{3, 0.5, 1.5}, ans: false},
		{equations: [][]string{{"le", "et"}, {"le", "code"}, {"code", "et"}}, values: []float64{2, 5, 0.5}, ans: true},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, checkContradictions(tt.equations, tt.values))
	}
}
