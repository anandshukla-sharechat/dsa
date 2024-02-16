package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	tt := struct {
		a, b []int
		want []int
	}{
		a:    []int{1, 2},
		b:    []int{3, 4, 5, 6},
		want: []int{1, 3, 2, 4, 5, 6},
	}
	assert.Equal(t, tt.want, solve(tt.a, tt.b))
}
