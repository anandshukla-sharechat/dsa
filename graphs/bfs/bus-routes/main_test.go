package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numBusesToDestination(t *testing.T) {
	tests := []struct {
		routes     [][]int
		source     int
		target     int
		numOfBuses int
	}{
		{routes: [][]int{{1, 2, 7}, {3, 6, 7}}, source: 1, target: 6, numOfBuses: 2},
		{routes: [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, source: 15, target: 12, numOfBuses: -1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.numOfBuses, numBusesToDestination(tt.routes, tt.source, tt.target))
	}
}
