package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StreamOfCharacter(t *testing.T) {
	streamChecker := Constructor([]string{"cd", "f", "kl"})
	assert.Equal(t, false, streamChecker.Query('a'))
	assert.Equal(t, false, streamChecker.Query('b'))
	assert.Equal(t, false, streamChecker.Query('c'))
	assert.Equal(t, true, streamChecker.Query('d'))
	assert.Equal(t, false, streamChecker.Query('e'))
	assert.Equal(t, true, streamChecker.Query('f'))
	assert.Equal(t, false, streamChecker.Query('g'))
	assert.Equal(t, false, streamChecker.Query('h'))
	assert.Equal(t, false, streamChecker.Query('i'))
	assert.Equal(t, false, streamChecker.Query('j'))
	assert.Equal(t, false, streamChecker.Query('k'))
	assert.Equal(t, true, streamChecker.Query('l'))

}
