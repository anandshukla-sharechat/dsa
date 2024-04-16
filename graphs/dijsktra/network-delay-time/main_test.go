package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_networkDelayTime(t *testing.T) {
	tests := []struct {
		times [][]int
		n     int
		k     int
		ans   int
	}{
		{times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, n: 4, k: 2, ans: 2},
		{times: [][]int{{1, 2, 1}}, n: 2, k: 2, ans: -1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, networkDelayTime(tt.times, tt.n, tt.k))
	}
}
