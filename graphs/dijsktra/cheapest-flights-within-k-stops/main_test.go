package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findCheapestPrice(t *testing.T) {
	tests := []struct {
		n            int
		flights      [][]int
		src          int
		dst          int
		k            int
		cheapestFare int
	}{
		{n: 4, flights: [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, src: 0, dst: 3, k: 1, cheapestFare: 6},
		{n: 4, flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, src: 0, dst: 3, k: 1, cheapestFare: 700},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.cheapestFare, findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k))
	}
}
