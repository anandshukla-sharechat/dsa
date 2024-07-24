package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calcEquation(t *testing.T) {
	tests := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		ans       []float64
	}{
		{equations: [][]string{{"a", "b"}, {"b", "c"}}, values: []float64{2.0, 3.0}, queries: [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}, ans: []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}},
		{equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, values: []float64{1.5, 2.5, 5.0}, queries: [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}, ans: []float64{3.75000, 0.40000, 5.00000, 0.20000}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, calcEquation(tt.equations, tt.values, tt.queries))
	}
}
