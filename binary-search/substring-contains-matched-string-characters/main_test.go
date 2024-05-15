package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minWindow(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		ans string
	}{
		{s: "ADOBECODEBANC", t: "ABC", ans: "BANC"},
		{s: "abc", t: "b", ans: "b"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.ans, minWindow(tt.s, tt.t))
	}
}
