package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	tests := []struct {
		products   []string
		searchWord string
		list       [][]string
	}{
		{products: []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, searchWord: "mouse", list: [][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}}},
		{products: []string{"havana"}, searchWord: "havana", list: [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.list, suggestedProducts(tt.products, tt.searchWord))
	}
}
