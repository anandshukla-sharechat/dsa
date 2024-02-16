package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solve(t *testing.T) {
	tt := struct {
		elements []int
		window   int
		output   []float64
	}{
		elements: []int{1, 10, 3, 5},
		window:   3,
		output:   []float64{1.0, 5.5, 4.666666666666667, 6.0},
	}
	assert.Equal(t, tt.output, solve(tt.window, tt.elements))
}
