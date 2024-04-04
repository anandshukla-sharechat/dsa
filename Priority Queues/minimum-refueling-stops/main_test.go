package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minRefuelStops(t *testing.T) {
	tests := []struct {
		target           int
		startFuel        int
		stations         [][]int
		minimumFuelStops int
	}{
		{target: 100, startFuel: 10, stations: [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}, minimumFuelStops: 2},
		{target: 100, startFuel: 50, stations: [][]int{{25, 25}, {50, 50}}, minimumFuelStops: 1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.minimumFuelStops, minRefuelStops(tt.target, tt.startFuel, tt.stations))
	}
}
