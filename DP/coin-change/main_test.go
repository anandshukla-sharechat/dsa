package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		ans    int
	}{
		{coins: []int{1, 3, 5}, amount: 11, ans: 3},
		{coins: []int{2}, amount: 3, ans: -1},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.ans, coinChange(tt.coins, tt.amount))
	}
}
