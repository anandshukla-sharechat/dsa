package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_replaceWords(t *testing.T) {
	tests := []struct {
		dictionary       []string
		sentence         string
		replacedSentence string
	}{
		{dictionary: []string{"cat", "bat", "rat"}, sentence: "the cattle was rattled by the battery", replacedSentence: "the cat was rat by the bat"},
		{dictionary: []string{"a", "b", "c"}, sentence: "aadsfasf absbs bbab cadsfafs", replacedSentence: "a a b c"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.replacedSentence, replaceWords(tt.dictionary, tt.sentence))
	}
}
