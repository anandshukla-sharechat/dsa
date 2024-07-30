package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repairCars(t *testing.T) {
	tests := []struct {
		ranks []int
		cars  int
		ans   int64
	}{
		{ranks: []int{4, 2, 3, 1}, cars: 10, ans: 16},
		{ranks: []int{5, 1, 8}, cars: 6, ans: 16},
		{ranks: []int{1, 1, 3, 3}, cars: 74, ans: 576},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, repairCars(tt.ranks, tt.cars))
	}
}
