package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDistance(t *testing.T) {
	tests := []struct {
		position             []int
		m                    int
		maximumMagneticForce int
	}{
		{position: []int{1, 2, 3, 4, 7}, m: 3, maximumMagneticForce: 3},
		{position: []int{5, 4, 3, 2, 1, 1000000000}, m: 2, maximumMagneticForce: 999999999},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.maximumMagneticForce, maxDistance(tt.position, tt.m))
	}
}
