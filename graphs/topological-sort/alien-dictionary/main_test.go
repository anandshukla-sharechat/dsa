package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_alienOrder(t *testing.T) {
	tests := []struct {
		words []string
		ans   string
	}{
		{words: []string{"wrt", "wrf", "er", "ett", "rftt"}, ans: "wertf"},
		{words: []string{"z", "x"}, ans: "zx"},
		{words: []string{"z", "x", "z"}, ans: ""},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.ans, alienOrder(tt.words))
	}
}
