package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_designAddAndSearchWords_ds(t *testing.T) {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	assert.Equal(t, false, wordDictionary.Search("pad"))
	assert.Equal(t, true, wordDictionary.Search("bad"))
	assert.Equal(t, true, wordDictionary.Search(".ad"))
	assert.Equal(t, true, wordDictionary.Search("b.."))
}
