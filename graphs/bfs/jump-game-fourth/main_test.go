package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minJumps(t *testing.T) {
	tests := []struct {
		arr []int
		ans int
	}{
		{arr: []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, ans: 3},
		{arr: []int{7, 6, 9, 6, 9, 6, 9, 7}, ans: 1},
		{arr: []int{7}, ans: 0},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, minJumps(tt.arr))
	}
}
